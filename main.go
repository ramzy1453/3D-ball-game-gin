package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramzy1453/3D-ball-game-gin/config"
	"github.com/ramzy1453/3D-ball-game-gin/routes"
)

func main() {

	// Initialize Gin router

	router := gin.Default()

	// Define routes

	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": "Micro Ball API made with Go Gin",
		})

	})

	config.ConnectDB()

	routes.PlayerRoute(router)

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}

}
