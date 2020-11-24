package handler

import (
	// ...
	"encoding/json"
	"io/ioutil"
	"net/http"
	// ...
	"github.com/arijitnayak92/taskAfford/TODONEW/schema"
	"github.com/taskAfford/TODONEW/service"
)

type todoHandler struct {
	postgres *db.Postgres
	samples  *db.Sample
}

//...
func (handler *todoHandler) saveTodo(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todo schema.Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := service.Insert(ctx, &todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, id)
}

func (handler *todoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := service.Delete(ctx, req.ID); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *todoHandler) getAllTodo(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)

	todoList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todoList)
}
