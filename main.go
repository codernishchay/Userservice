package main

import (
	"fmt"
	"userservice/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	app.App(r)
	fmt.Println(r)
	fmt.Println(" User service ")
}
