package app

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/REST/controllers"
)

func StartApp() {
	fmt.Print("Here")
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
