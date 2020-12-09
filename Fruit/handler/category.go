package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/affmed/apperrors"
	"net/http"
)

func (h *Handler) CreateCategoryHandler(c *gin.Context) {
	req, err := h.validation.CreateCategoryValidation(c)
	if err != nil {
		return
	}

	category, err := h.domain.CreateCategory(req.Name, req.Description)
	if err != nil {
		if err == apperrors.ErrCreateCategoryFailed {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "category created successfully",
		"id":      category.ID,
	})
}
