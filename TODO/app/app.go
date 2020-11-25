package app

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var router = mux.NewRouter().StrictSlash(true)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

func StartApp() {

	Routes()
	http.ListenAndServe(":8080", router)
}
