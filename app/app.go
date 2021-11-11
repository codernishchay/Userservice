package app

import (
	"github.com/gin-gonic/gin"
)

func App(c *gin.Engine) {

	DBConnect()
	Routers(c)
}
