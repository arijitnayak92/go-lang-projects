package handler

import (
	"net/http"

	"github.com/arijitnayak92/taskAfford/TODONEW/db"
)

// - func SetUpRouting() {
func SetUpRouting(postgres *db.Postgres) *http.ServeMux {
	todoHandler := &todoHandler{
		postgres: postgres,
		samples:  &db.Sample{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/samples", todoHandler.GetSamples)
	mux.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.getAllTodo(w, r)
		case http.MethodPost:
			todoHandler.saveTodo(w, r)
		case http.MethodDelete:
			todoHandler.deleteTodo(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})

	return mux
}
