package services

import (
	"net/http"

	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/domain"
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/utils"
)

var (
	UserService userServiceIntrface
)

type userServiceIntrface interface {
	Login(u *domain.User) (map[string]string, *utils.APIError)
	CreateUser(user *domain.User) (*domain.User, *utils.APIError)
	RefreshToken(req *http.Request) (map[string]string, *utils.APIError)
	Logout(req *http.Request) (int, *utils.APIError)
}

func init() {
	UserService = &usersServices{}
}

type usersServices struct{}

func (u *usersServices) Login(user *domain.User) (map[string]string, *utils.APIError) {
	return domain.UserMethods.Login(user)
}

func (u *usersServices) CreateUser(user *domain.User) (*domain.User, *utils.APIError) {
	return domain.UserMethods.CreateUser(user)
}

func (u *usersServices) RefreshToken(req *http.Request) (map[string]string, *utils.APIError) {
	return domain.UserMethods.Refresh(req)
}

func (u *usersServices) Logout(req *http.Request) (int, *utils.APIError) {
	return domain.UserMethods.Logout(req)
}
