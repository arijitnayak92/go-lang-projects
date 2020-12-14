package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit2/apperrors"
	"github.com/arijitnayak92/taskAfford/Fruit2/domain"
	mocks "github.com/arijitnayak92/taskAfford/Fruit2/mocks/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	var appUserDomain *domain.User
	newUser := NewUser(appUserDomain)

	if newUser.domain == nil {
		t.Errorf("error in NewUser() constructor")
	}
}

func TestUser_AddUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("1: When user added successfully!!", func(t *testing.T) {
		mockUserDomain := new(mocks.MockUserDomain)
		testUser := `{"firstname":"Test123","lastname":"Test12345","email":"Test@123.com","password":"test12345","confirmPassword":"test12345"}`
		mockUserDomain.On("AddUser").Return(true, nil)
		testUserHandler := NewUser(mockUserDomain)

		router := gin.Default()
		router.POST("/signup", testUserHandler.AddUser)
		w := executeRequest(router, "POST", "/signup", bytes.NewBufferString(testUser))

		mockUserDomain.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		message := "Successfully SignedUp!!"
		body := gin.H{
			"message": message,
		}
		value, _ := response["message"]
		fmt.Println(body["message"])
		fmt.Println(value)
		assert.Equal(t, body["message"], value)

	})
	t.Run("2: When user Already Exists!!", func(t *testing.T) {
		mockUserDomain := new(mocks.MockUserDomain)
		testUser := `{"firstname":"Test123","lastname":"Test12345","email":"Test@123.com","password":"test12345","confirmPassword":"test12345"}`
		mockUserDomain.On("AddUser").Return(false, apperrors.ErrResourceAlreadyExists)
		testUserHandler := NewUser(mockUserDomain)

		router := gin.Default()
		router.POST("/signup", testUserHandler.AddUser)
		w := executeRequest(router, "POST", "/signup", bytes.NewBufferString(testUser))

		mockUserDomain.AssertExpectations(t)
		assert.Equal(t, 409, w.Code)

		var response map[string]string
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		body := gin.H{
			"error": apperrors.ErrResourceAlreadyExists.Error(),
		}
		value, _ := response["error"]
		assert.Equal(t, body["error"], value)

	})
	t.Run("3: When user validations fails!!", func(t *testing.T) {
		mockUserDomain := new(mocks.MockUserDomain)
		testUser := `{"firstname":"","lastname":"T","email":"Test@123","password":"test12345","confirmPassword":"test12345"}`
		testUserHandler := NewUser(mockUserDomain)

		router := gin.Default()
		router.POST("/signup", testUserHandler.AddUser)
		w := executeRequest(router, "POST", "/signup", bytes.NewBufferString(testUser))

		mockUserDomain.AssertExpectations(t)
		assert.Equal(t, 400, w.Code)
	})
	t.Run("4: When user password verification fails!!", func(t *testing.T) {
		mockUserDomain := new(mocks.MockUserDomain)
		testUser := `{"firstname":"Test123","lastname":"Test12345","email":"Test@123.com","password":"test12345","confirmPassword":"test12"}`
		testUserHandler := NewUser(mockUserDomain)

		router := gin.Default()
		router.POST("/signup", testUserHandler.AddUser)
		w := executeRequest(router, "POST", "/signup", bytes.NewBufferString(testUser))

		mockUserDomain.AssertExpectations(t)
		assert.Equal(t, 401, w.Code)

		var response map[string]string
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		body := gin.H{
			"error": apperrors.ErrPasswordVerification.Error(),
		}
		value, _ := response["error"]
		assert.Equal(t, body["error"], value)
	})
}
