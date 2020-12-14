package validation

import (
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/gin-gonic/gin"
)

// AppValidation ...
type AppValidation interface {
	SignUpValidation(c *gin.Context) (*SignUpRequest, error)
	SignInValidation(c *gin.Context) (*SignInRequest, error)
}

// Validation ...
type Validation struct {
	util utils.AppUtil
}

// NewValidation ...
func NewValidation(u utils.AppUtil) *Validation {
	return &Validation{util: u}
}
