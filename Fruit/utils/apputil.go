package utils

import "github.com/dgrijalva/jwt-go"

type AppUtil interface {
	HashPassword(password string) (string, error)
	CompareHashedPasswords(password, hashedPassword string) bool
	IsEmail(email string) (bool, error)
	PasswordValidation(password string) (bool, error)
	CreateToken(email string) (*TokenDetails, error)
	TokenValid(t string) error
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type Util struct {
}

func NewUtil() *Util {
	return &Util{}
}
