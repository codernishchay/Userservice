package app

import (
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.GET("/get", GetUser)
	router.POST("/create", CreateUser)
	router.PUT("/update", UpdateUser)
	router.DELETE("/delete", DeleteUser)
	router.Run()
}
