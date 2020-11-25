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

// "SetUpRouting  ..."
func SetUpRouting(postgres *db.Postgres) *mux.Router {
	todoHandler := &todoHandler{
		postgres: postgres,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", response)
	router.HandleFunc("/getAll", todoHandler.getAllTodo).Methods("GET")
	router.HandleFunc("/getOneTodo/{todo_id}", todoHandler.getOneTodo).Methods("GET")
	router.HandleFunc("/addTodo", todoHandler.saveTodo).Methods("POST")
	router.HandleFunc("/deleteTodo/{todo_id}", todoHandler.deleteTodo).Methods("DELETE")
	router.HandleFunc("/updateTodo/{todo_id}", todoHandler.updateTodo).Methods("PUT")
	router.HandleFunc("/mark-as-done/{todo_id}", todoHandler.markAsDone).Methods("POST")
	return router
}
