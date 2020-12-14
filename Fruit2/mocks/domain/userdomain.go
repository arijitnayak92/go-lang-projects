package mocks

import (
	"github.com/arijitnayak92/taskAfford/Fruit2/models"
	"github.com/stretchr/testify/mock"
)

// MockUserDomain ..struct
type MockUserDomain struct {
	mock.Mock
}

// AddUser ... mock adduser
func (mock *MockUserDomain) AddUser(user models.User) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}
