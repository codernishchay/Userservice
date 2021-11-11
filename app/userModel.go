package app

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id",omitempty`
	Name        string             `json:"name" bson:"name",omitempty`
	DOB         primitive.DateTime `json:"dob" bson:"dob",omitempty`
	Address     string             `json:"address" bson:"address",omitempty`
	Description string             `json:"description" bson:"description"`
}
