package mocks

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/affordmed/fruit-seller-b-backend/models"
)

// MockUserDomain ..struct
type MockUserDomain struct {
	mock.Mock
}

// AddUser ... mock adduser
func (mock *MockUserDomain) AddUser(user models.User) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
