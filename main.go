// main.go
package main

import (
	"github.com/gin-gonic/gin"
)


func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("before")
		c.Next()
	}
}


func main() {
	// Create a new Gin router
	router := gin.Default()

	router.Use(Logger())

	// Define a route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Nm!",
		})
	})

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
