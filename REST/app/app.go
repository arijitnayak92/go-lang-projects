package app

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/REST/controllers"
)

func response(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Backend Responding !")
}

func StartApp() {
	http.HandleFunc("/", response)
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Print("Server Connected !")
}
