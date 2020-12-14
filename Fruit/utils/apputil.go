package utils

import "github.com/dgrijalva/jwt-go"

// AppUtil ...
type AppUtil interface {
	HashPassword(password string) (string, error)
	CompareHashedPasswords(password, hashedPassword string) bool
	IsEmail(email string) (bool, error)
	PasswordValidation(password string) (bool, error)
	CreateToken(email string) (*TokenDetails, error)
	TokenValid(t string) error
	VerifyToken(tokenString string) (*jwt.Token, error)
}

// Util ...
type Util struct {
}

// NewUtil ...
func NewUtil() *Util {
	return &Util{}
}
