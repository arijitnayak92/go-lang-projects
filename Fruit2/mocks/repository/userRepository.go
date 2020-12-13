package mocks

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/affordmed/fruit-seller-b-backend/models"
)

// MockUserRepository ...
type MockUserRepository struct {
	mock.Mock
}

// CreateUser : mocked create user for unit testing
func (mock *MockUserRepository) CreateUser(m models.User) (string, error) {

	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)

}
