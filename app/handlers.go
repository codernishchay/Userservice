package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *gin.Context) {
	user := new(User)
	err := json.NewDecoder(c.Request.Body).Decode(user)
	if err != nil {
		return
	}
	result, _ := Database.Collection("users").InsertOne(context.TODO(), user)

	fmt.Println(result)

}

func UpdateUser(c *gin.Context) {
	// userid , update by it
	// find and update
	// objectid := c.Request.Body
	id, _ := primitive.ObjectIDFromHex("objectidfromapp")
	result, err := Database.Collection("user").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"user", "username"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated ", result.ModifiedCount)
}

func GetUser(c *gin.Context) {
	cursor, err := Database.Collection("users").Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}

func DeleteUser(c *gin.Context) {
	id := c.Request.Body
	bson := bson.M{"id": id}
	results, err := Database.Collection("users").DeleteOne(context.TODO(), bson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" deleted one user , ", results.DeletedCount)
}
