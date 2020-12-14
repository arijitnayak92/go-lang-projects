package mock

import (
	"github.com/arijitnayak92/taskAfford/Fruit/validation"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// Validation ...
type Validation struct {
	mock.Mock
}

// NewValidation ...
func NewValidation() *Validation {
	return &Validation{}
}

// SignUpValidation ...
func (mock *Validation) SignUpValidation(_ *gin.Context) (*validation.SignUpRequest, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*validation.SignUpRequest), args.Error(1)
}

// SignInValidation ...
func (mock *Validation) SignInValidation(_ *gin.Context) (*validation.SignInRequest, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*validation.SignInRequest), args.Error(1)
}
