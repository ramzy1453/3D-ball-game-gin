package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"github.com/ramzy1453/3D-ball-game-gin/config"
	"github.com/ramzy1453/3D-ball-game-gin/models"
	"github.com/ramzy1453/3D-ball-game-gin/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var playerCollection *mongo.Collection = config.GetCollection(config.DB, "players")
var validate = validator.New()

func CreatePlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var player models.Player
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&player); err != nil {
			c.JSON(http.StatusBadRequest, responses.PlayerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&player); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PlayerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newUser := models.Player{
			Id:     primitive.NewObjectID(),
			Name:   player.Name,
			Scores: []float32{},
		}
		userExistent := playerCollection.FindOne(ctx, bson.M{"name": player.Name})
		if userExistent.Err() == nil {
			c.JSON(http.StatusOK, responses.PlayerResponse{
				Status:  http.StatusBadRequest,
				Message: "player-already-exists",
			})
			return
		}

		result, err := playerCollection.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.PlayerResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"player": result}})
	}

}

func GetPlayers() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.Player
		defer cancel()

		cursor, err := playerCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var singleUser models.Player
			if err = cursor.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.PlayerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}

}

func UpdateScore() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		playerId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		var body responses.ScoreBody
		err := c.BindJSON(&body)

		if err != nil {
			c.JSON(http.StatusBadRequest, responses.PlayerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		filter := bson.M{"_id": playerId}
		update := bson.M{
			"$push": bson.M{
				"scores": body.Score,
			},
		}

		result, err := playerCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PlayerResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
			})
			return
		}

		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, responses.PlayerResponse{Status: http.StatusNotFound, Message: "not found"})
			return
		}

		c.JSON(http.StatusOK, responses.PlayerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"player": result}})

	}

}

func GetLeaderboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pipeline := mongo.Pipeline{
			{
				{"$project", bson.M{
					"name":      1,
					"bestScore": bson.M{"$min": "$scores"},
					"scores":    1,
				}},
			},
			{
				{"$match", bson.M{
					"scores": bson.M{"$ne": bson.A{}},
				}},
			},
			{
				{"$sort", bson.D{
					{"bestScore", 1},
				}},
			},
		}

		cursor, err := playerCollection.Aggregate(ctx, pipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PlayerResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
			})
		}

		defer cursor.Close(ctx)

		var results []responses.Leaderboard
		if err = cursor.All(ctx, &results); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PlayerResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
			})
			return
		}

		c.JSON(http.StatusOK, responses.PlayerResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{
				"results": results,
			},
		})

	}

}

func ResetLeaderboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		results, err := playerCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PlayerResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
			})
		}

		c.JSON(http.StatusOK, responses.PlayerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"results": results}})

	}

}
