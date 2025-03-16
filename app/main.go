package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/javuk/end2end/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	pkg.InitDB()

	r := gin.Default()

	r.LoadHTMLFiles(
		filepath.Join("templates", "employees.html"),
		filepath.Join("templates", "employee_details.html"),
	)
	r.Static("/static", "./static")

	//Routes
	r.GET("/", serveEmployeesPage)
	r.GET("/employees/:id/details", serveEmployeeDetailsPage)

	r.GET("/employees", pkg.GetEmployees)
	r.POST("/employees", pkg.AddNewEmployee)
	r.GET("/employees/:id", pkg.GetEmployeeByID)
	r.DELETE("/employees/:id", pkg.DeleteEmployee)
	r.PUT("/employees/:id", pkg.UpdateEmployee)

	r.Use(func(c *gin.Context) {
		if len(c.Request.URL.Path) >= 8 && c.Request.URL.Path[:8] == "/static/" {
			fmt.Printf("Static file requested: %s\n", c.Request.URL.Path)
		}
		c.Next()
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080" // Default for local testing
	}
	r.Run(":" + port) // Listen on the specified port
}

func serveEmployeesPage(c *gin.Context) {
	c.HTML(http.StatusOK, "employees.html", nil)
}

func serveEmployeeDetailsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "employee_details.html", nil)
}
