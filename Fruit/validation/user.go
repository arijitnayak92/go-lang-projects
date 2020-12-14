package validation

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/Fruit/apperrors"
	"github.com/gin-gonic/gin"
)

// SignUpRequest ...
type SignUpRequest struct {
	Email           string `form:"email" json:"email" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required"`
	FirstName       string `form:"firstname" json:"firstName" binding:"required"`
	LastName        string `form:"lastname" json:"lastName" binding:"required"`
}

// SignInRequest ...
type SignInRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// SignInValidation ...
func (v *Validation) SignInValidation(c *gin.Context) (*SignInRequest, error) {
	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Required fields are missing",
		})
		return nil, err
	}

	return &req, nil
}

// SignUpValidation ...
func (v *Validation) SignUpValidation(c *gin.Context) (*SignUpRequest, error) {
	var req SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Required fields are missing",
		})
		return nil, err
	}
	if req.Password == "" || req.ConfirmPassword == "" || req.Email == "" || req.FirstName == "" || req.LastName == "" {
		c.JSON(406, gin.H{
			"message": apperrors.ErrRequiredField,
		})
		return nil, apperrors.ErrRequiredField
	}
	if req.Password != req.ConfirmPassword {

		c.JSON(http.StatusBadRequest, gin.H{
			"password":        "mismatched",
			"confirmPassword": "mismatched",
		})
		return nil, apperrors.ErrConfirmPasswordMismatched
	}
	isEmail, err := v.util.IsEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return nil, err
	}
	if !isEmail {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email address",
		})
		return nil, apperrors.ErrInvalidEmail
	}

	isPassword, err := v.util.PasswordValidation(req.Password)
	if !isPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return nil, apperrors.ErrInvalidPassword
	}
	return &req, nil
}
