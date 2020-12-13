package apperrors

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ErrorModel ...
type ErrorModel struct {
	Message string
	Code    int
}

var (
	// ErrInternalServerError ...
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrPostgresConnection ...
	ErrPostgresConnection = errors.New("Unable to connect to the Postgres database")
	// ErrMongoConnection ...
	ErrMongoConnection = errors.New("Unable to connect to the Mongo database")
	// ErrDataNotFound ...
	ErrDataNotFound = errors.New("Failed to find requested information")
	// ErrActionFailed ...
	ErrActionFailed = errors.New("Failed to complete requested action")
	// ErrBadRequest ....
	ErrBadRequest = errors.New("Failed to parse body")
	// ErrDatabaseRecord ...
	ErrDatabaseRecord = errors.New("Information is distorted")
	// ErrResourceAlreadyExists ...
	ErrResourceAlreadyExists = errors.New("Resource already exists")
	// ErrEmailValidation ..
	ErrEmailValidation = errors.New("Wrong email format")
	// ErrFieldValidation ...
	ErrFieldValidation = errors.New("Field values have errors")
	// ErrPasswordVerification ...
	ErrPasswordVerification = errors.New("Password Verification Failed")
)

func assertError(err error) *ErrorModel {
	if errors.Is(err, ErrResourceAlreadyExists) {
		return &ErrorModel{
			Message: ErrResourceAlreadyExists.Error(),
			Code:    409,
		}
	}
	if errors.Is(err, ErrPasswordVerification) {
		return &ErrorModel{
			Message: ErrPasswordVerification.Error(),
			Code:    401,
		}
	}

	if errors.Is(err, ErrEmailValidation) {
		return &ErrorModel{
			Message: ErrEmailValidation.Error(),
			Code:    400,
		}
	}
	if errors.Is(err, ErrFieldValidation) {
		return &ErrorModel{
			Message: ErrFieldValidation.Error(),
			Code:    400,
		}
	}
	if errors.Is(err, ErrInternalServerError) {
		return &ErrorModel{
			Message: ErrInternalServerError.Error(),
			Code:    500,
		}
	}
	if errors.Is(err, ErrDatabaseRecord) {
		return &ErrorModel{
			Message: ErrDatabaseRecord.Error(),
			Code:    417,
		}
	}
	if errors.Is(err, ErrPostgresConnection) {
		return &ErrorModel{
			Message: ErrPostgresConnection.Error(),
			Code:    500,
		}
	}
	if errors.Is(err, ErrMongoConnection) {
		return &ErrorModel{
			Message: ErrMongoConnection.Error(),
			Code:    500,
		}
	}
	if errors.Is(err, ErrDataNotFound) {
		return &ErrorModel{
			Message: ErrDataNotFound.Error(),
			Code:    404,
		}
	}
	if errors.Is(err, ErrBadRequest) {
		return &ErrorModel{
			Message: ErrBadRequest.Error(),
			Code:    400,
		}
	}
	if errors.Is(err, ErrActionFailed) {
		return &ErrorModel{
			Message: ErrActionFailed.Error(),
			Code:    501,
		}
	}
	return &ErrorModel{
		Message: "Unidentified Error",
		Code:    500,
	}
}

// ErrorResponse : Pass error message to the server.
func ErrorResponse(e error, c *gin.Context) {
	err := assertError(e)
	c.JSON(err.Code, gin.H{"error": err.Message})
	return
}
