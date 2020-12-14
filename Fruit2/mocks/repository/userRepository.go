package mocks

import (
	"github.com/arijitnayak92/taskAfford/Fruit2/models"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository ...
type MockUserRepository struct {
	mock.Mock
}

// CreateUser : mocked create user for unit testing
func (mock *MockUserRepository) CreateUser(m models.User) (bool, error) {

	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)

}
