package domain

import (
	"errors"
	"time"

	"github.com/arijitnayak92/taskAfford/Fruit2/apperrors"
	"github.com/arijitnayak92/taskAfford/Fruit2/db"
	"github.com/arijitnayak92/taskAfford/Fruit2/models"
	"github.com/arijitnayak92/taskAfford/Fruit2/utils"
)

// UserDomain Interface
type UserDomain interface {
	AddUser(user models.User) (bool, error)
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
func (u *User) AddUser(user models.User) (bool, error) {

	user.Password = utils.GetHash([]byte(user.Password))
	user.ConfirmPassword = ""
	user.Role = "User"
	user.CreatedAt = time.Now()
	_, err := u.userRepo.CreateUser(user)
	if err != nil {
		if errors.Is(err, apperrors.ErrResourceAlreadyExists) {
			return false, apperrors.ErrResourceAlreadyExists
		}
		return false, apperrors.ErrActionFailed
	}
	return true, nil
}
