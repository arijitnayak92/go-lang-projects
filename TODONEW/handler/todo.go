package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/arijitnayak92/taskAfford/TODONEW/db"
	"github.com/arijitnayak92/taskAfford/TODONEW/schema"
	"github.com/arijitnayak92/taskAfford/TODONEW/service"
	"github.com/arijitnayak92/taskAfford/TODONEW/utils"
	"github.com/gorilla/mux"
)

type TodoHandler struct {
	postgres *db.Postgres
}

func (handler *TodoHandler) SaveTodo(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "Can't able to read user data !",
			StatusCode: 500,
		}
		utils.ResponseError(w, apiError)
		return
	}

	var todo schema.Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		apiError := &utils.APIError{
			Message:    "Wrong Input Data !",
			StatusCode: 400,
		}
		utils.ResponseError(w, apiError)
		return
	}
	if todo.Title == "" || todo.Note == "" {
		apiError := &utils.APIError{
			Message:    "Enter all the details !",
			StatusCode: 406,
		}
		utils.ResponseError(w, apiError)
		return
	}

	id, errI := service.Insert(ctx, &todo)
	if errI != nil {
		apiError := &utils.APIError{
			Message:    errI.Message,
			StatusCode: errI.StatusCode,
		}
		utils.ResponseError(w, apiError)
		return
	}
	utils.ResponseOK(w, id)
}

func (handler *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)
	params := mux.Vars(r)
	todoID := params["todo_id"]
	intID, errC := strconv.Atoi(todoID)
	if errC != nil {
		apiError := &utils.APIError{
			Message:    "Todo Id should be a number !",
			StatusCode: 422,
		}
		utils.ResponseError(w, apiError)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "Can't able to read user data !",
			StatusCode: 500,
		}
		utils.ResponseError(w, apiError)
		return
	}

	var todo schema.Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		apiError := &utils.APIError{
			Message:    "Wrong Input Data !",
			StatusCode: 400,
		}
		utils.ResponseError(w, apiError)
		return
	}

	if todo.Title == "" || todo.Note == "" {
		apiError := &utils.APIError{
			Message:    "Enter all the details !",
			StatusCode: 406,
		}
		utils.ResponseError(w, apiError)
		return
	}

	errs := service.Update(ctx, intID, &todo)
	if errs != nil {
		apiError := &utils.APIError{
			Message:    errs.Message,
			StatusCode: errs.StatusCode,
		}
		utils.ResponseError(w, apiError)
		return
	}
	utils.ResponseOK(w, nil)
}

func (handler *TodoHandler) MarkAsDone(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)
	params := mux.Vars(r)
	todoID := params["todo_id"]
	intID, errC := strconv.Atoi(todoID)
	if errC != nil {
		apiError := &utils.APIError{
			Message:    "Todo Id should be a number !",
			StatusCode: 422,
		}
		utils.ResponseError(w, apiError)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "Can't able to read user data !",
			StatusCode: 500,
		}
		utils.ResponseError(w, apiError)
		return
	}

	var req struct {
		Status bool `json:"status"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		apiError := &utils.APIError{
			Message:    "Wrong Input Data !",
			StatusCode: 400,
		}
		utils.ResponseError(w, apiError)
		return
	}

	errs := service.MarkAsDone(ctx, intID, req.Status)
	if errs != nil {
		apiError := &utils.APIError{
			Message:    errs.Message,
			StatusCode: errs.StatusCode,
		}
		utils.ResponseError(w, apiError)
		return
	}

	utils.ResponseOK(w, nil)
}

func (handler *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)

	params := mux.Vars(r)
	todoID := params["todo_id"]
	intID, errC := strconv.Atoi(todoID)
	if errC != nil {
		apiError := &utils.APIError{
			Message:    "Todo Id should be a number !",
			StatusCode: 422,
		}
		utils.ResponseError(w, apiError)
		return
	}

	if err := service.Delete(ctx, intID); err != nil {
		apiError := &utils.APIError{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
		utils.ResponseError(w, apiError)
		return
	}
	utils.ResponseOK(w, nil)
}

func (handler *TodoHandler) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)

	todoList, err := service.GetAll(ctx)
	if err != nil {
		apiError := &utils.APIError{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
		utils.ResponseError(w, apiError)
		return
	}

	utils.ResponseOK(w, todoList)
}

func (handler *TodoHandler) GetOneTodo(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.postgres)
	params := mux.Vars(r)
	todoID := params["todo_id"]
	intID, errC := strconv.Atoi(todoID)
	if errC != nil {
		apiError := &utils.APIError{
			Message:    "Todo Id should be a number !",
			StatusCode: 422,
		}
		utils.ResponseError(w, apiError)
		return
	}
	todoOne, err := service.GetOne(ctx, intID)
	if err != nil {
		apiError := &utils.APIError{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
		utils.ResponseError(w, apiError)
		return
	}

	utils.ResponseOK(w, todoOne)
}
