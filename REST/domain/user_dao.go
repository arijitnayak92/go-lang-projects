package domain

import (
	"github.com/arijitnayak92/taskAfford/REST/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Arijit", LastName: "Nayak", Email: "arijitnayak92@gmail.com"},
	}
	UserMethods userInterface
)

func init() {
	UserMethods = &usersStruct{}
}

type userInterface interface {
	GetUser(userId int64) (*User, *utils.APIError)
}

type usersStruct struct{}

func (c *usersStruct) GetUser(userId int64) (*User, *utils.APIError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.APIError{
		Message:    "User Not Found !",
		StatusCode: 404,
	}
}
