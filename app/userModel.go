package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	DOB         string             `json:"dob" bson:"dob"`
	Address     string             `json:"address" bson:"address"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   string
}
