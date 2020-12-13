package apperrors

import (
	"errors"
)

var ErrWrongPassword = errors.New("Wrong Password")
var ErrConfirmPasswordMismatched = errors.New("Password and confirm password mismatched")
var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyPresent = errors.New("User already present")
var ErrInvalidEmail = errors.New("Invalid email address")
var ErrRequiredField = errors.New("Enter all required fields")

var ErrTokenExpired = errors.New("Token Expired")
var ErrWrongJwtMethod = errors.New("Unexpected signin method for token")
