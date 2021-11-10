package main

import (
	"fmt"
	"userservice/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	app.App()
	app.DBConnect()
	fmt.Println(r)
	fmt.Println(" User service ")
}
