package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// https://stackoverflow.com/questions/37990507/golang-mgo-how-can-i-store-date-not-isodate-in-mongodb

func CreateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	user := new(User)
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	id := primitive.NewObjectIDFromTimestamp(time.Time{})
	fmt.Println(id)
	if err != nil {
		log.Fatal(err)
	}
	user.ID = id
	dbcollection := DATABASE.Collection("users")
	result, err := dbcollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	json.NewEncoder(c.Writer).Encode(result)
}

func UpdateUser(c *gin.Context) {
	fmt.Println("Hello there")
	c.Header("Content-Type", "application/json")
	id, er := primitive.ObjectIDFromHex(c.Query("id"))
	fmt.Println(id)
	fmt.Println(er)
	if er != nil {
		log.Fatal(er)
	}
	var user User

	filter := bson.M{"_id": id}
	fmt.Println(filter)

	_ = json.NewDecoder(c.Request.Body).Decode(&user)

	update := bson.D{
		{
			"$set", bson.D{
				{"name", user.Name},
				{"dob", user.DOB},
				{"address", user.Address},
				{"description", user.Description},
			},
		},
	}

	err := DATABASE.Collection("users").FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	user.ID = id
	json.NewEncoder(c.Writer).Encode(user)
}

func GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	fmt.Println("get request heree")
	var users []User
	cursor, err := DATABASE.Collection("users").Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cursor)

	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(c.Writer).Encode(users)
}

func DeleteUser(c *gin.Context) {
	fmt.Println(" delete endpoint")
	id, er := primitive.ObjectIDFromHex(c.Query("id"))
	if er != nil {
		log.Fatal(er)
	}
	bson := bson.M{"_id": id}
	fmt.Println(bson)
	results, err := DATABASE.Collection("users").DeleteOne(context.TODO(), bson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" deleted one user , ", results.DeletedCount)
	json.NewEncoder(c.Writer).Encode(results)
}
