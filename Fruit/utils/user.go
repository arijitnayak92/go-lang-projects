package utils

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// EmailRegex ...
var EmailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

//IsEmail ...
func (u *Util) IsEmail(email string) (bool, error) {
	return regexp.MatchString(EmailRegex, email)
}

// PasswordValidation ...
func (u *Util) PasswordValidation(password string) (bool, error) {
	if len(password) < 8 {
		lenError := errors.New("Password Length is less than 8")
		return false, lenError
	}
	num := `[0-9]{1}`
	small := `[a-z]{1}`
	upper := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, password); !b || err != nil {
		lenError := errors.New("Password Should Have A Number")
		return false, lenError
	}
	if b, err := regexp.MatchString(small, password); !b || err != nil {
		lenError := errors.New("Password Should Have A Lower Case Letter")
		return false, lenError
	}
	if b, err := regexp.MatchString(upper, password); !b || err != nil {
		lenError := errors.New("Password Should Have A Upper Case Letter")
		return false, lenError
	}
	if b, err := regexp.MatchString(symbol, password); !b || err != nil {
		lenError := errors.New("Password Should Have A Special Character")
		return false, lenError
	}
	return true, nil
}

// HashPassword ...
func (u *Util) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CompareHashedPasswords ...
func (u *Util) CompareHashedPasswords(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}

	return true
}
