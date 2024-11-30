package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ramzy1453/3D-ball-game-gin/handlers"
)

func PlayerRoute(router *gin.Engine) {
	router.POST("/player", controllers.CreatePlayer())
	router.GET("/player", controllers.GetPlayers())
	router.GET("/leaderboard", controllers.GetLeaderboard())
	router.PUT("/player/score/:id", controllers.UpdateScore())
	router.DELETE("/leaderboard", controllers.ResetLeaderboard())
}
