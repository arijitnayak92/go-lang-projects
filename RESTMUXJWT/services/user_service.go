package services

import (
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/domain"
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/utils"
	"github.com/gin-gonic/gin"
)

var (
	UserService userServiceIntrface
)

type userServiceIntrface interface {
	Login(u *domain.User) (map[string]string, *utils.APIError)
	CreateUser(user *domain.User) (*domain.User, *utils.APIError)
	RefreshToken(c *gin.Context) (map[string]string, *utils.APIError)
	Logout(c *gin.Context) (int, *utils.APIError)
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

func (u *usersServices) RefreshToken(c *gin.Context) (map[string]string, *utils.APIError) {
	return domain.UserMethods.Refresh(c)
}

func (u *usersServices) Logout(c *gin.Context) (int, *utils.APIError) {
	return domain.UserMethods.Logout(c)
}
