package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/fruit-seller-b-backend/apperrors"
	"gitlab.com/affordmed/fruit-seller-b-backend/domain"
	"gitlab.com/affordmed/fruit-seller-b-backend/models"
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

	//log.Println(user)

	email, err := u.domain.AddUser(user)
	if err != nil {
		apperrors.ErrorResponse(err, c)
		return
	}
	message := email + " added successfully!!"
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
	return

}
