package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/apperrors"
	"github.com/arijitnayak92/taskAfford/Fruit/mock"
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/arijitnayak92/taskAfford/Fruit/validation"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestLoginHandler(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	u := utils.NewUtil()
	validationMock := mock.NewValidation()
	testuserv := &validation.SignInRequest{
		Email:    "abc@gmail.com",
		Password: "Arijitnayak@92",
	}

	mockUserDomain := new(mock.Domain)
	testUser := `{"email":"abc@gmail.com","password":"Arijitnayak@92"}`
	validationMock.On("SignInValidation").Return(testuserv, nil)
	utilss := utils.NewUtil()
	hashedP, _ := utilss.HashPassword(testuserv.Password)
	utilss.CompareHashedPasswords(testuserv.Password, hashedP)
	token, _ := utilss.CreateToken(testuserv.Email)
	mockUserDomain.On("Login").Return(token.RefreshToken, nil)
	testUserHandler := NewHandler(appCtx, mockUserDomain, validationMock, u)

	router := gin.Default()
	router.POST("/login", testUserHandler.Login)
	w := executeRequest(router, "POST", "/login", bytes.NewBufferString(testUser))

	mockUserDomain.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	message := "Successfully LoggedIn"
	body := gin.H{
		"message": message,
		"token":   token,
	}
	value, _ := response["message"]
	assert.Equal(t, body["message"], value)
	assert.Equal(t, body["token"], token)
}

func TestLoginHandlerUserNotFound(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	u := utils.NewUtil()
	validationMock := mock.NewValidation()
	testuserv := &validation.SignInRequest{
		Email:    "abc@gmail.com",
		Password: "Arijitnayak@92",
	}

	mockUserDomain := new(mock.Domain)
	testUser := `{"email":"abc@gmail.com","password":"Arijitnayak@92"}`
	validationMock.On("SignInValidation").Return(testuserv, nil)
	utilss := utils.NewUtil()
	hashedP, _ := utilss.HashPassword(testuserv.Password)
	utilss.CompareHashedPasswords(testuserv.Password, hashedP)
	utilss.CreateToken(testuserv.Email)

	mockUserDomain.On("Login").Return("", apperrors.ErrUserNotFound)
	testUserHandler := NewHandler(appCtx, mockUserDomain, validationMock, u)

	router := gin.Default()
	router.POST("/login", testUserHandler.Login)
	w := executeRequest(router, "POST", "/login", bytes.NewBufferString(testUser))

	mockUserDomain.AssertExpectations(t)
	assert.Equal(t, 404, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	message := "user not found"
	body := gin.H{
		"message": message,
	}
	value, _ := response["message"]
	assert.Equal(t, body["message"], value)
}

func TestLoginHandlerWrongPassword(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	u := utils.NewUtil()
	validationMock := mock.NewValidation()
	testuserv := &validation.SignInRequest{
		Email:    "abc@gmail.com",
		Password: "Arijitnayak@92",
	}

	mockUserDomain := new(mock.Domain)
	testUser := `{"email":"abc@gmail.com","password":"Arijitnayak@92"}`
	validationMock.On("SignInValidation").Return(testuserv, nil)
	utilss := utils.NewUtil()
	hashedP, _ := utilss.HashPassword("wrong password")
	utilss.CompareHashedPasswords(testuserv.Password, hashedP)
	utilss.CreateToken(testuserv.Email)

	mockUserDomain.On("Login").Return("", apperrors.ErrWrongPassword)
	testUserHandler := NewHandler(appCtx, mockUserDomain, validationMock, u)

	router := gin.Default()
	router.POST("/login", testUserHandler.Login)
	w := executeRequest(router, "POST", "/login", bytes.NewBufferString(testUser))

	mockUserDomain.AssertExpectations(t)
	assert.Equal(t, 401, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	message := "Wrong Password"
	body := gin.H{
		"message": message,
	}
	value, _ := response["message"]
	assert.Equal(t, body["message"], value)
}
