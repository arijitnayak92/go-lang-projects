package validation

import "github.com/gin-gonic/gin"

type CreateCategoryRequest struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"name" json:"description"`
}

func (v *Validation) CreateCategoryValidation(c *gin.Context) (*CreateCategoryRequest, error) {
	var req CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v.RespondWithValidationErrors(c, v.ListOfErrors(&req, err), nil)
		return nil, err
	}

	return &req, nil
}
