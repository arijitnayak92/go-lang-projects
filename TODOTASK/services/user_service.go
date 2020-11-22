package services

import (
	"net/http"

	"github.com/arijitnayak92/taskAfford/RESTTODO/domain"
	"github.com/arijitnayak92/taskAfford/RESTTODO/utils"
)

var (
	UserServiceMux userServiceIntrface
)

type userServiceIntrface interface {
	Login(u *domain.User) (map[string]string, *utils.APIError)
	CreateUser(user *domain.User) (interface{}, *utils.APIError)
	RefreshTokens(req *http.Request) (map[string]string, *utils.APIError)
	LogoutUser(req *http.Request) (int, *utils.APIError)
}

func init() {
	UserServiceMux = &usersServices{}
}

type usersServices struct{}

func (u *usersServices) Login(user *domain.User) (map[string]string, *utils.APIError) {
	return domain.UserMethodMux.Login(user)
}

func (u *usersServices) CreateUser(user *domain.User) (interface{}, *utils.APIError) {
	return domain.UserMethodMux.CreateUser(user)
}

func (u *usersServices) RefreshTokens(req *http.Request) (map[string]string, *utils.APIError) {
	return domain.UserMethodMux.RefreshToken(req)
}

func (u *usersServices) LogoutUser(req *http.Request) (int, *utils.APIError) {
	return domain.UserMethodMux.LogoutUser(req)
}
