package services

import (
	"github.com/arijitnayak92/taskAfford/REST/domain"
	"github.com/arijitnayak92/taskAfford/REST/utils"
)

var (
	UserService userServiceIntrface
)

type userServiceIntrface interface {
	GetUser(userId int64) (*domain.User, *utils.APIError)
}

func init() {
	UserService = &usersServices{}
}

type usersServices struct{}

func (u *usersServices) GetUser(userId int64) (*domain.User, *utils.APIError) {
	return domain.UserMethods.GetUser(userId)
}
