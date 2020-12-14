package apperrors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// BindingError ..
type BindingError struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewValidatorError ...
func NewValidatorError(err error) BindingError {
	res := BindingError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}

	}
	return res
}
