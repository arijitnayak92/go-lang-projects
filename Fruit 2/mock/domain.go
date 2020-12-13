package mock

import (
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/stretchr/testify/mock"
)

// MockDomain struct containing testify mock
type MockDomain struct {
	mock.Mock
}

func (mock *MockDomain) CheckDatabaseHealth() (error, error) {
	args := mock.Called()

	return args.Error(0), args.Error(1)
}

func (mock *MockDomain) UserSignup(email string, password string, firstname string, lastname string, role string) (bool, error) {
	args := mock.Called()

	return false, args.Error(0)
}

func (mock *MockDomain) GetUser(email string) (*model.User, error) {
	args := mock.Called()

	return nil, args.Error(0)
}
