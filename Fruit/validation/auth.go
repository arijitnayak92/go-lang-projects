package validation

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/affmed/apperrors"
	"net/http"
)

type SignInRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required"`
	OldPassword     string `form:"oldPassword" json:"oldPassword" binding:"required"`
	Email           string `json:"email"`
}

func (v *Validation) SignInValidation(c *gin.Context) (*SignInRequest, error) {
	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v.RespondWithValidationErrors(c, v.ListOfErrors(&req, err), nil)
		return nil, err
	}

	return &req, nil
}

func (v *Validation) ChangePasswordValidation(c *gin.Context) (*ChangePasswordRequest, error) {
	var req ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v.RespondWithValidationErrors(c, v.ListOfErrors(&req, err), nil)
		return nil, err
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"password":        "mismatched",
			"confirmPassword": "mismatched",
		})

		return nil, apperrors.ErrEmailMismatched
	}

	if req.OldPassword == req.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"password":    "matched",
			"oldPassword": "matched",
		})

		return nil, apperrors.ErrEmailMismatched
	}

	req.Email, _ = c.Params.Get("email")

	isEmail, err := v.util.IsEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if !isEmail {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email address",
		})
		return nil, apperrors.ErrInvalidEmail
	}

	return &req, nil
}
