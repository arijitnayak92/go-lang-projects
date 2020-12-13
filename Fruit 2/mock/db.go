package mock

import (
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/stretchr/testify/mock"
)

// MockRepository ...
type MockRepository struct {
	mock.Mock
}

// PostgresHealthCheck mocked postgres health check method
func (mock *MockRepository) PingPostgres() error {
	args := mock.Called()

	return args.Error(0)
}

// MongoHealthCheck mocked mongo health check method
func (mock *MockRepository) CheckMongoAlive() error {
	args := mock.Called()

	return args.Error(0)
}

// MongoHealthCheck mocked mongo health check method
func (mock *MockRepository) GetUser(email string) (*model.User, error) {
	return nil, nil
}

// MongoHealthCheck mocked mongo health check method
func (mock *MockRepository) UserSignup(email string, password string, firstname string, lastname string, role string) (bool, error) {
	return false, nil
}
