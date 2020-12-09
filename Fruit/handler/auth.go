package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/affmed/apperrors"
	"net/http"
)

func (h *Handler) SignInHandler(c *gin.Context) {
	req, err := h.validation.SignInValidation(c)
	if err != nil {
		return
	}

	err = h.domain.SignInUser(req.Email, req.Password)
	if err != nil {
		if err == apperrors.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})

			return
		}

		if err == apperrors.ErrPasswordMismatched {
			c.JSON(http.StatusUnauthorized, gin.H{
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
		"message": "authentication successful",
	})
}

func (h *Handler) ChangePasswordHandler(c *gin.Context) {
	req, err := h.validation.ChangePasswordValidation(c)
	if err != nil {
		return
	}

	user, err := h.domain.GetUserByEmail(req.Email)
	if err != nil {
		if err == apperrors.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !h.util.CompareHashedPasswords(req.OldPassword, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "authentication failed",
		})
		return
	}

	hashedPassword, err := h.util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	err = h.domain.ChangePassword(req.Email, hashedPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}
