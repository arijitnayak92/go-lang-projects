package handler

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/TODONEW/db"
	"github.com/gorilla/mux"
)

func response(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Backend Responding !")
}

//...
func SetUpRouting(postgres *db.Postgres) *mux.Router {
	todoHandler := &todoHandler{
		postgres: postgres,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", response)
	router.HandleFunc("/todos", todoHandler.getAllTodo).Methods("GET")
	router.HandleFunc("/todos/{todo_id}", todoHandler.getOneTodo).Methods("GET")
	router.HandleFunc("/todo", todoHandler.saveTodo).Methods("POST")
	router.HandleFunc("/todos/{todo_id}", todoHandler.deleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{todo_id}", todoHandler.updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{todo_id}", todoHandler.markAsDone).Methods("PATCH")
	return router
}
