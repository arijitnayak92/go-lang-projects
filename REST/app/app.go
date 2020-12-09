package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
}

func StartApp() {
	Routes()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
