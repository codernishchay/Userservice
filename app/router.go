package app

import (
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.GET("/get", GetUser)
	// fmt.Println("/get  := get all users")
	router.POST("/create", CreateUser)
	// fmt.Println("/post := create a user")

	router.PUT("/update", UpdateUser)
	// fmt.Println("/update  := update a user")

	router.DELETE("/delete", DeleteUser)
	// fmt.Println("/delete  := delete a user")
	router.Run()
}
