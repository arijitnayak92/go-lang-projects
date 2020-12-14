package models

import (
	"time"
)

// User : Details of the User.
type User struct {
	FirstName       string    `form:"firstname" json:"firstName,omitempty" binding:"required,max=16,min=2"`
	LastName        string    `form:"lastname" json:"lastName,omitempty" binding:"required,max=16,min=3"`
	Email           string    `form:"email" json:"email,omitempty" binding:"required,email"`
	Password        string    `form:"password" json:"password" binding:"required,max=128,min=6"`
	ConfirmPassword string    `form:"confirmPassword" json:"confirmPassword" binding:"required,max=128,min=6"`
	CartID          int       `form:"cartid" json:"cartId,omitempty"`
	Role            string    `form:"role" json:"role,omitempty"`
	CreatedAt       time.Time `form:"createdAt" json:"createdAt,omitempty"`
}

//func NewUser(fname string,lname string,)
