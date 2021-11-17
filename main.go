package main

import (
	"userservice/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	app.DBConnect()
	app.Routers(r)
}
