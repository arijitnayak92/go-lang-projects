package handler

import (
	"net/http"

	"github.com/arijitnayak92/taskAfford/Fruit2/apperrors"
	"github.com/arijitnayak92/taskAfford/Fruit2/domain"
	"github.com/arijitnayak92/taskAfford/Fruit2/models"
	"github.com/gin-gonic/gin"
)

// UserHandler : UserHandler Interface.
type UserHandler interface {
	AddUser(c *gin.Context)
}

// User : User Handler struct.
type User struct {
	domain domain.UserDomain
}

// NewUser : Constructor for User struct.
func NewUser(domain domain.UserDomain) *User {
	return &User{domain: domain}
}

// AddUser : Handler for adding user.
func (u *User) AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		bindErr := apperrors.NewValidatorError(err)
		c.JSON(400, bindErr)
		return
	}

	if user.Password != user.ConfirmPassword {
		apperrors.ErrorResponse(apperrors.ErrPasswordVerification, c)
		return
	}

	_, err := u.domain.AddUser(user)
	if err != nil {
		apperrors.ErrorResponse(err, c)
		return
	}
	message := "Successfully SignedUp!!"
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
	return

}
