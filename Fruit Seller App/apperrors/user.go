package apperrors

import (
	"errors"
)

// ErrDuplicateEnrty ...
var ErrDuplicateEnrty = errors.New("Duplicate Entry Can Not Be Processed")

// ErrWrongPassword ...
var ErrWrongPassword = errors.New("Wrong Password")

// ErrConfirmPasswordMismatched ...
var ErrConfirmPasswordMismatched = errors.New("Password and confirm password mismatched")

// ErrUserNotFound ...
var ErrUserNotFound = errors.New("user not found")

// ErrUserAlreadyPresent ...
var ErrUserAlreadyPresent = errors.New("User already present")

// ErrInvalidEmail ...
var ErrInvalidEmail = errors.New("Invalid email address")

// ErrRequiredField ...
var ErrRequiredField = errors.New("Enter all required fields")

// ErrInvalidPassword ...
var ErrInvalidPassword = errors.New("Password Validation Failed")

// ErrTokenExpired ...
var ErrTokenExpired = errors.New("Token Expired")

// ErrWrongJwtMethod ...
var ErrWrongJwtMethod = errors.New("Unexpected signin method for token")
