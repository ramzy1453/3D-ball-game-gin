package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Scores []int              `json:"scores"`
}
