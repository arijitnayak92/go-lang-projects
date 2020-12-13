package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// GetHash function to hash the password using bcrypt
func GetHash(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
