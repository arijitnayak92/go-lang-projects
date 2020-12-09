package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/affordmed/affmed/util"
	"net/http"
	"reflect"
)

type VErrors []map[string]string

type AppValidation interface {
	SignInValidation(c *gin.Context) (*SignInRequest, error)
	ChangePasswordValidation(c *gin.Context) (*ChangePasswordRequest, error)
	CreateCategoryValidation(c *gin.Context) (*CreateCategoryRequest, error)
	RespondWithValidationErrors(c *gin.Context, vErrors VErrors, err error)
	ListOfErrors(v interface{}, e error) VErrors
}

type Validation struct {
	util util.AppUtil
}

func NewValidation(u util.AppUtil) *Validation {
	return &Validation{util: u}
}

func (v *Validation) ListOfErrors(validationStruct interface{}, e error) VErrors {
	ve := e.(validator.ValidationErrors)
	InvalidFields := make([]map[string]string, 0)

	for _, e := range ve {
		errors := map[string]string{}
		field, _ := reflect.TypeOf(validationStruct).Elem().FieldByName(e.Field())
		jsonTag := field.Tag.Get("json")
		errors[jsonTag] = e.Tag()
		InvalidFields = append(InvalidFields, errors)
	}

	return InvalidFields
}

func (v *Validation) RespondWithValidationErrors(c *gin.Context, vErrors VErrors, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": vErrors,
	})

}
