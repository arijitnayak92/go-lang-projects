package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/domain"
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/services"
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/utils"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var u *domain.User
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		apiError := &utils.APIError{
			Message:    "Invalid token provided !",
			StatusCode: 422,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	fmt.Println(u)
	if u.Username == "" || u.Password == "" {
		apiError := &utils.APIError{
			Message:    "Enter all the details !",
			StatusCode: 406,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, apiError := services.UserServiceMux.CreateUser(u)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	res.WriteHeader(200)
	res.Write(jsonValue)
}

func Login(res http.ResponseWriter, req *http.Request) {
	var u *domain.User
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		apiError := &utils.APIError{
			Message:    "Invalid user data !",
			StatusCode: 422,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	if u.Username == "" || u.Password == "" {
		apiError := &utils.APIError{
			Message:    "Enter all the details !",
			StatusCode: 406,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, apiError := services.UserServiceMux.Login(u)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	res.WriteHeader(200)
	res.Write(jsonValue)
}

func RefreshToken(res http.ResponseWriter, req *http.Request) {
	token, apiError := services.UserServiceMux.RefreshTokens(req)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(token)
	res.WriteHeader(200)
	res.Write(jsonValue)
}

func Logout(res http.ResponseWriter, req *http.Request) {
	_, apiError := services.UserServiceMux.LogoutUser(req)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	res.WriteHeader(200)
	res.Write([]byte("Successfully Loggedout !"))
}
