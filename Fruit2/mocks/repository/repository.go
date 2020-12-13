package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockRepository ...
type MockRepository struct {
	mock.Mock
}

// PostgresHealthCheck mocked postgres health check method
func (mock *MockRepository) PostgresHealthCheck() error {
	args := mock.Called()

	return args.Error(0)
}

// MongoHealthCheck mocked mongo health check method
func (mock *MockRepository) MongoHealthCheck() error {
	args := mock.Called()

	return args.Error(0)
}
