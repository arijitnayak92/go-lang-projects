package apperrors

import (
	"errors"
)

var ErrPasswordMismatched = errors.New("authentication failed")

var ErrPasswordUpdateFailed = errors.New("password update failed")

var ErrUserNotFound = errors.New("user not found")
