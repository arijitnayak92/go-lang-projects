package domain

import (
	"errors"
	"time"

	"gitlab.com/affordmed/fruit-seller-b-backend/apperrors"
	"gitlab.com/affordmed/fruit-seller-b-backend/db"
	"gitlab.com/affordmed/fruit-seller-b-backend/models"
	"gitlab.com/affordmed/fruit-seller-b-backend/utils"
)

// UserDomain Interface
type UserDomain interface {
	AddUser(user models.User) (string, error)
}

// User structs
type User struct {
	userRepo db.UserRepository
}

// NewUser Domain Constructor
func NewUser(db db.UserRepository) *User {
	return &User{userRepo: db}
}

// AddUser ..domain method to add user.
func (u *User) AddUser(user models.User) (string, error) {

	user.Password = utils.GetHash([]byte(user.Password))
	user.ConfirmPassword = ""
	user.Role = "User"
	user.CreatedAt = time.Now()
	//user.UpdatedAt = time.Now()
	email, err := u.userRepo.CreateUser(user)
	if err != nil {
		if errors.Is(err, apperrors.ErrResourceAlreadyExists) {
			return "", apperrors.ErrResourceAlreadyExists
		}
		return "", apperrors.ErrActionFailed
	}
	return email, nil
}
