package main

import (
	//...
	"fmt"
	"log"
	"net/http"

	"github.com/arijitnayak92/taskAfford/TODONEW/db"
	"github.com/arijitnayak92/taskAfford/TODONEW/handler"
	"github.com/arijitnayak92/taskAfford/TODONEW/utils"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	//...
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func main() {
	var postgres *db.Postgres
	var err *utils.APIError

	postgres, err = db.ConnectPostgres()

	if err != nil {
		apiError := &utils.APIError{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
		fmt.Println(apiError)
	} else if postgres == nil {
		panic("postgres is nil")
	}

	mux := handler.SetUpRouting(postgres)
	routes := cors.New(cors.Options{AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"}}).Handler(mux)

	fmt.Println("Connecting : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
