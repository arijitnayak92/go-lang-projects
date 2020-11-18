package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func StartApp() {
	Routes()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
