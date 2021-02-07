package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/apperrors"
)

// SignUpUser ...
func (h *Handler) SignUpUser(c *gin.Context) {
	req, errs := h.validation.SignUpValidation(c)
	if errs != nil {
		return
	}

	hashedPassword, errp := h.util.HashPassword(req.Password)
	if errp != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errp.Error(),
		})
	}

	_, err := h.domain.UserSignup(req.Email, hashedPassword, req.FirstName, req.LastName, "User", "")

	if err != nil {
		if err == apperrors.ErrUserAlreadyPresent {
			c.JSON(409, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully SignedUp",
	})
}

// Login ...
func (h *Handler) Login(c *gin.Context) {
	req, errs := h.validation.SignInValidation(c)
	if errs != nil {
		return
	}

	token, err := h.domain.Login(req.Email, req.Password)

	if err != nil {
		if err == apperrors.ErrUserNotFound {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err == apperrors.ErrWrongPassword {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully LoggedIn",
		"token":   token,
	})
}
