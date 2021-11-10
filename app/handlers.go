package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
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
	// find and update
}

func GetUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
