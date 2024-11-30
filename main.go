package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramzy1453/3D-ball-game-gin/config"
)



func main() {

	// Initialize Gin router

	router := gin.Default()

	// Define routes

	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": "Hello, Gin!",

		})

	})

	config.ConnectDB()
	router.Run(":8080")

}
