package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	initDB()

	r := gin.Default()

	r.LoadHTMLFiles(
		filepath.Join("templates", "employees.html"),
		filepath.Join("templates", "employee_details.html"),
	)
	r.Static("/static", "./static")

	//Routes
	r.GET("/", serveEmployeesPage)
	r.GET("/employees/:id/details", serveEmployeeDetailsPage)

	r.GET("/employees", getEmployees)
	r.POST("/employees", addNewEmployee)
	r.GET("/employees/:id", getEmployeeByID)
	r.DELETE("/employees/:id", deleteEmployee)
	r.PUT("/employees/:id", updateEmployee)

	r.Use(func(c *gin.Context) {
		if len(c.Request.URL.Path) >= 8 && c.Request.URL.Path[:8] == "/static/" {
			fmt.Printf("Static file requested: %s\n", c.Request.URL.Path)
		}
		c.Next()
	})

	r.Run(":8080")
}

func serveEmployeesPage(c *gin.Context) {
	c.HTML(http.StatusOK, "employees.html", nil)
}

func serveEmployeeDetailsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "employee_details.html", nil)
}
