package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/affordmed/fruit-seller-b-backend/apperrors"
	mocks "gitlab.com/affordmed/fruit-seller-b-backend/mocks/repository"
	"gitlab.com/affordmed/fruit-seller-b-backend/models"

	"gitlab.com/affordmed/fruit-seller-b-backend/db"
)

func TestNewUser(t *testing.T) {
	var userRepo *db.PostgresRepo
	newUser := NewUser(userRepo)

	if newUser.userRepo == nil {
		t.Errorf("error in NewUser() constructor")
	}

}

func TestUser_AddUser(t *testing.T) {
	t.Run("When user added successfully!!", func(t *testing.T) {
		mockUserRepo := new(mocks.MockUserRepository)
		testUser := models.User{
			FirstName:       "Test123",
			LastName:        "Test12345",
			Email:           "Test@123.com",
			Password:        "test12345",
			ConfirmPassword: "test12345",
		}
		mockUserRepo.On("CreateUser").Return(testUser.Email, nil)
		testUserDomain := NewUser(mockUserRepo)
		email, err := testUserDomain.AddUser(testUser)

		assert.NoError(t, err)
		assert.Equal(t, "Test@123.com", email)

	})
	t.Run("When user addition Failed!!", func(t *testing.T) {
		mockUserRepo := new(mocks.MockUserRepository)
		testUser := models.User{
			FirstName:       "Test123",
			LastName:        "Test12345",
			Email:           "Test@123.com",
			Password:        "test12345",
			ConfirmPassword: "test12345",
		}
		mockUserRepo.On("CreateUser").Return("", apperrors.ErrActionFailed)
		testUserDomain := NewUser(mockUserRepo)
		email, err := testUserDomain.AddUser(testUser)

		if assert.Error(t, err) {
			assert.Equal(t, apperrors.ErrActionFailed, err)
		}
		assert.Equal(t, "", email)

	})
	t.Run("When user already exists!!", func(t *testing.T) {
		mockUserRepo := new(mocks.MockUserRepository)
		testUser := models.User{
			FirstName:       "Test123",
			LastName:        "Test12345",
			Email:           "Test@123.com",
			Password:        "test12345",
			ConfirmPassword: "test12345",
		}
		mockUserRepo.On("CreateUser").Return("", apperrors.ErrResourceAlreadyExists)
		testUserDomain := NewUser(mockUserRepo)
		email, err := testUserDomain.AddUser(testUser)

		if assert.Error(t, err) {
			assert.Equal(t, apperrors.ErrResourceAlreadyExists, err)
		}
		assert.Equal(t, "", email)

	})

}
