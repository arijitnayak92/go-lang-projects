package mock

import (
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/stretchr/testify/mock"
)

// MockDomain struct containing testify mock
type Domain struct {
	mock.Mock
}

func (mock *Domain) GetPostgresHealth() bool {
	args := mock.Called()

	return args.Bool(0)
}

func (mock *Domain) GetMongoHealth() bool {
	args := mock.Called()

	return args.Bool(0)
}

func (mock *Domain) UserSignup(email string, password string, firstname string, lastname string, role string) (bool, error) {
	args := mock.Called(email, password, firstname, lastname, role)
	return args.Bool(0), args.Error(1)
}

func (mock *Domain) GetUser(email string) (*model.User, error) {
	args := mock.Called(email)
	result := args.Get(0)
	return result.(*model.User), args.Error(1)
}

func (mock *Domain) Login(email string, password string) (string, error) {
	args := mock.Called(email, password)

	return args.String(0), args.Error(1)
}
