package validation

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
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
