package app

import (
	"log"

	"net/http"

	"github.com/arijitnayak92/taskAfford/TODO/config"
	"github.com/arijitnayak92/taskAfford/TODO/domain"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var router = mux.NewRouter().StrictSlash(true)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config.Load()
	domain.InitDB()
}

func StartApp() {
	Routes()
	http.ListenAndServe(":8080", router)
}
