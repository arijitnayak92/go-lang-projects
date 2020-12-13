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

// SignUpValidation ...
func (mock *Validation) SignUpValidation(c *gin.Context) (*validation.SignUpRequest, error) {
	// args := mock.Called()
	// return nil, args.Error(0)

	return &validation.SignUpRequest{
		Email:           "abc@gmail.com",
		Password:        "Arijitnayak@92",
		ConfirmPassword: "Arijitnayak@92",
		FirstName:       "Arijit",
		LastName:        "Nayak",
	}, nil

}
