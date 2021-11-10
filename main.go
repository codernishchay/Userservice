package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println(r)
	DBConnect()
	fmt.Println(" User service ")
}
