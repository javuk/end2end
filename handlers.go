package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET all employees
func getEmployees(c *gin.Context) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

// GET one specific emplayee
func getEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errror": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// POST add new employee
func addNewEmployee(c *gin.Context) {
	newEmployee := Employee{
		FirstName:        c.PostForm("first_name"),
		LastName:         c.PostForm("last_name"),
		Gender:           c.PostForm("gender"),
		BirthYear:        parseInt(c.PostForm("birth_year")),
		StartDate:        c.PostForm("start_date"),
		ContractType:     c.PostForm("contract_type"),
		ContractDuration: parseInt(c.PostForm("contract_duration")),
		Department:       c.PostForm("department"),
		AnnualLeaveDays:  parseInt(c.PostForm("annual_leave_days")),
		FreeDays:         parseInt(c.PostForm("free_days")),
		PaidLeaveDays:    parseInt(c.PostForm("paid_leave_days")),
	}

	file, err := c.FormFile("image_path")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed: " + err.Error()})
		return
	}

	// image path
	imagePath := filepath.Join("static", "images", file.Filename)

	// Save the file to the server
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image: " + err.Error()})
		return
	}

	// Save image path
	newEmployee.ImagePath = "/" + imagePath

	// Insert into database
	query := `INSERT INTO employees (first_name, last_name, gender, birth_year, start_date, contract_type, contract_duration, department, annual_leave_days, free_days, paid_leave_days, image_path)
	          VALUES (:first_name, :last_name, :gender, :birth_year, :start_date, :contract_type, :contract_duration, :department, :annual_leave_days, :free_days, :paid_leave_days, :image_path)`

	_, err = db.NamedExec(query, newEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insert failed: " + err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message":    "Employee added successfully",
		"employee":   newEmployee,
		"image_path": newEmployee.ImagePath,
	})
}

// PUT an employee
func updateEmployee(c *gin.Context) {
	// Get the employee ID from URL parameters
	id := c.Param("id")

	// Parse form fields
	updatedEmployee := Employee{
		FirstName:        c.PostForm("first_name"),
		LastName:         c.PostForm("last_name"),
		Gender:           c.PostForm("gender"),
		BirthYear:        parseInt(c.PostForm("birth_year")),
		StartDate:        c.PostForm("start_date"),
		ContractType:     c.PostForm("contract_type"),
		ContractDuration: parseInt(c.PostForm("contract_duration")),
		Department:       c.PostForm("department"),
		AnnualLeaveDays:  parseInt(c.PostForm("annual_leave_days")),
		FreeDays:         parseInt(c.PostForm("free_days")),
		PaidLeaveDays:    parseInt(c.PostForm("paid_leave_days")),
	}

	// Handle image upload if present
	file, err := c.FormFile("image_path")
	if err == nil { // Only handle file if it exists
		// Define the save path for the image
		imagePath := filepath.Join("static", "images", file.Filename)

		// Save the file to the server
		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image: " + err.Error()})
			return
		}

		// Save image path to the struct
		updatedEmployee.ImagePath = "/" + imagePath // Store relative path
	} else {
		// Dohvati trenutni image_path iz baze ako korisnik nije upload-ao novu sliku
		var currentImagePath string
		err = db.Get(&currentImagePath, "SELECT image_path FROM employees WHERE id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve current image path: " + err.Error()})
			return
		}
		updatedEmployee.ImagePath = currentImagePath
	}

	// Convert ID from string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Update the employee ID
	updatedEmployee.ID = intID

	// Define the update query
	query := `UPDATE employees SET 
		first_name = :first_name, 
		last_name = :last_name, 
		gender = :gender, 
		birth_year = :birth_year, 
		start_date = :start_date, 
		contract_type = :contract_type, 
		contract_duration = :contract_duration, 
		department = :department, 
		annual_leave_days = :annual_leave_days, 
		free_days = :free_days, 
		paid_leave_days = :paid_leave_days, 
		image_path = :image_path
		WHERE id = :id`

	// Execute the update query
	_, err = db.NamedExec(query, updatedEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database update failed: " + err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message":    "Employee updated successfully",
		"employee":   updatedEmployee,
		"image_path": updatedEmployee.ImagePath,
	})
}

// DELETE an employee
func deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM employees WHERE id = $1`

	res, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

// Helper function to convert string to int
func parseInt(value string) int {
	var result int
	fmt.Sscanf(value, "%d", &result)
	return result
}
