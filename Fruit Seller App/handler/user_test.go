package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/apperrors"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/mock"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/model"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/validation"
)

func TestSignUpHandler(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	u := utils.NewUtil()
	validationMock := mock.NewValidation()
	testuser := &model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}

	testuserv := &validation.SignUpRequest{
		Email:           "abc@gmail.com",
		Password:        "Arijitnayak@92",
		ConfirmPassword: "Arijitnayak@92",
		FirstName:       "Arijit",
		LastName:        "Nayak",
	}

	mockUserDomain := new(mock.Domain)
	testUser := `{"firstname":"Arijit","lastname":"Nayak","email":"abc@gmail.com","password":"Arijitnayak@92","confirmPassword":"Arijitnayak@92"}`
	validationMock.On("SignUpValidation").Return(testuserv, nil)
	utilss := utils.NewUtil()
	hashp, _ := utilss.HashPassword(testuserv.Password)
	testuser.Password = hashp
	mockUserDomain.On("UserSignup").Return(true, nil)
	testUserHandler := NewHandler(appCtx, mockUserDomain, validationMock, u)

	router := gin.Default()
	router.POST("/signup", testUserHandler.SignUpUser)
	w := executeRequest(router, "POST", "/signup", bytes.NewBufferString(testUser))

	mockUserDomain.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	message := "Successfully SignedUp"
	body := gin.H{
		"message": message,
	}
	value, _ := response["message"]
	assert.Equal(t, body["message"], value)

}

func TestSignUpHandlerErrorUserAlreadyExist(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	u := utils.NewUtil()
	validationMock := mock.NewValidation()
	testuser := &model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}

	testuserv := &validation.SignUpRequest{
		Email:           "abc@gmail.com",
		Password:        "Arijitnayak@92",
		ConfirmPassword: "Arijitnayak@92",
		FirstName:       "Arijit",
		LastName:        "Nayak",
	}

	mockUserDomain := new(mock.Domain)
	testUser := `{"firstname":"Arijit","lastname":"Nayak","email":"abc@gmail.com","password":"Arijitnayak@92","confirmPassword":"Arijitnayak@92"}`
	validationMock.On("SignUpValidation").Return(testuserv, nil)
	utilss := utils.NewUtil()
	hashp, _ := utilss.HashPassword(testuserv.Password)
	testuser.Password = hashp
	mockUserDomain.On("UserSignup").Return(false, apperrors.ErrUserAlreadyPresent)
	testUserHandler := NewHandler(appCtx, mockUserDomain, validationMock, u)

	router := gin.Default()
	router.POST("/signup", testUserHandler.SignUpUser)
	w := executeRequest(router, "POST", "/signup", bytes.NewBufferString(testUser))

	mockUserDomain.AssertExpectations(t)
	assert.Equal(t, 409, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	message := "User already present"
	body := gin.H{
		"message": message,
	}
	value, _ := response["message"]
	assert.Equal(t, body["message"], value)

}

func TestHandlerSignInHandler(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	u := utils.NewUtil()
	validationMock := mock.NewValidation()
	testuserv := &validation.SignInRequest{
		Email:    "abc@gmail.com",
		Password: "Arijitnayak@92",
	}
	testUser := `{"email":"abc@gmail.com","password":"Arijitnayak@92"}`
	utilss := utils.NewUtil()

	token, _ := utilss.CreateToken(testuserv.Email)
	cases := map[string]struct {
		want   string
		status int
	}{
		"when user is unique and signup is successful": {
			want:   "Successfully LoggedIn",
			status: http.StatusOK,
		},
		"when user not found": {
			want:   "user not found",
			status: 404,
		},
		"when user entered wrong password": {
			want:   "Wrong Password",
			status: 401,
		},
	}

	for k, v := range cases {
		mockUserDomain := new(mock.Domain)
		validationMock.On("SignInValidation").Return(testuserv, nil)
		testUserHandler := NewHandler(appCtx, mockUserDomain, validationMock, u)
		var hashedP string
		if v.status == 401 {
			hashedP, _ = utilss.HashPassword("wrong password")
		}
		hashedP, _ = utilss.HashPassword(testuserv.Password)
		utilss.CompareHashedPasswords(testuserv.Password, hashedP)

		t.Run(k, func(t *testing.T) {
			if v.status == 200 {
				mockUserDomain.On("Login").Return(token.RefreshToken, nil)
			} else if v.status == 404 {
				mockUserDomain.On("Login").Return("", apperrors.ErrUserNotFound)
			} else {
				mockUserDomain.On("Login").Return("", apperrors.ErrWrongPassword)
			}
			router := gin.Default()
			router.POST("/login", testUserHandler.Login)
			w := executeRequest(router, "POST", "/login", bytes.NewBufferString(testUser))

			mockUserDomain.AssertExpectations(t)
			var response map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &response)
			assert.Nil(t, err)
			assert.Equal(t, v.status, w.Code)
			assert.Equal(t, v.want, response["message"])
		})
	}

}
