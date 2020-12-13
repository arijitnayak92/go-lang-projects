package apperrors

import (
	"errors"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_assertError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *ErrorModel
	}{
		{"Resource exists Error", args{ErrResourceAlreadyExists}, &ErrorModel{
			Message: ErrResourceAlreadyExists.Error(),
			Code:    409,
		}},
		{"Password Verification Error", args{ErrPasswordVerification}, &ErrorModel{
			Message: ErrPasswordVerification.Error(),
			Code:    401,
		}},
		{"Email Validation Failed", args{ErrEmailValidation}, &ErrorModel{
			Message: ErrEmailValidation.Error(),
			Code:    400,
		}},
		{"Other Field Validation Failed", args{ErrFieldValidation}, &ErrorModel{
			Message: ErrFieldValidation.Error(),
			Code:    400,
		}},
		{"Internal Server Error", args{ErrInternalServerError}, &ErrorModel{
			Message: ErrInternalServerError.Error(),
			Code:    500,
		}},
		{"Database Record Error", args{ErrDatabaseRecord}, &ErrorModel{
			Message: ErrDatabaseRecord.Error(),
			Code:    417,
		}},
		{"Postgres Connection Error", args{ErrPostgresConnection},
			&ErrorModel{
				Message: ErrPostgresConnection.Error(),
				Code:    500,
			}},
		{"Mongo Connection Error", args{ErrMongoConnection},
			&ErrorModel{
				Message: ErrMongoConnection.Error(),
				Code:    500,
			}},
		{"Data Not Found", args{ErrDataNotFound}, &ErrorModel{
			Message: ErrDataNotFound.Error(),
			Code:    404,
		}},
		{"Bad Request Error", args{ErrBadRequest}, &ErrorModel{
			Message: ErrBadRequest.Error(),
			Code:    400,
		}},
		{"Action Failed Error", args{ErrActionFailed}, &ErrorModel{
			Message: ErrActionFailed.Error(),
			Code:    501,
		}},
		{"Default Error", args{errors.New("Unidentified error")}, &ErrorModel{
			Message: "Unidentified Error",
			Code:    500,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assertError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assertError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorResponse(t *testing.T) {
	// w := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(w)
	type args struct {
		e error
		//c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Bad Request", args{ErrBadRequest}, 400},
		{"Action Failed", args{ErrActionFailed}, 501},
		{"Data Not Found", args{ErrDataNotFound}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			ErrorResponse(tt.args.e, c)
			if !reflect.DeepEqual(w.Code, tt.want) {
				t.Errorf("ErrorResponse() = %v, want %v", w.Code, tt.want)
			}
		})
	}
}
