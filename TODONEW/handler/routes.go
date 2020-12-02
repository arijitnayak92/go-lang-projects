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
	todoHandler := &TodoHandler{
		postgres: postgres,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", response)
	router.HandleFunc("/todos", todoHandler.GetAllTodo).Methods("GET")
	router.HandleFunc("/todos/{todo_id}", todoHandler.GetOneTodo).Methods("GET")
	router.HandleFunc("/todo", todoHandler.SaveTodo).Methods("POST")
	router.HandleFunc("/todos/{todo_id}", todoHandler.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{todo_id}", todoHandler.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{todo_id}", todoHandler.MarkAsDone).Methods("PATCH")
	return router
}
