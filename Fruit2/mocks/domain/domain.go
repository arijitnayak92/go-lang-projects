package mocks

import "github.com/stretchr/testify/mock"

// MockDomain struct containing testify mock
type MockDomain struct {
	mock.Mock
}

// DatabaseHealthCheck method of mockDomain for unit testing
func (mock *MockDomain) DatabaseHealthCheck() (error, error) {
	args := mock.Called()

	return args.Error(0), args.Error(1)
}
