package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arijitnayak92/taskAfford/REST/services"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		res.WriteHeader(403)
		res.Write([]byte("User Id Should be a number "))
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("User Not Found !"))
		return
	}
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
