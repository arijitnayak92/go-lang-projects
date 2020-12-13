package db

import (
	"log"
	"strings"
	"time"

	"gitlab.com/affordmed/fruit-seller-b-backend/apperrors"

	"gitlab.com/affordmed/fruit-seller-b-backend/models"
)

const queryInsertUser = "INSERT INTO users (email,cartid,password,firstname,lastname,role,createdat) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING email;"

// UserRepository ...
type UserRepository interface {
	CreateUser(m models.User) (string, error)
}

// CreateUser function to insert user into the database
func (repo *PostgresRepo) CreateUser(user models.User) (string, error) {
	var (
		email     string
		cartID    string
		password  string
		firstName string
		lastName  string
		role      string
		createdAt time.Time
	)
	//userExistingError:='pq: duplicate key value violates unique constraint "users_pkey"'
	err := repo.db.QueryRow(queryInsertUser, user.Email, user.CartID, user.Password, user.FirstName, user.LastName, user.Role, user.CreatedAt).Scan(&email, &cartID, &password, &firstName, &lastName, &role, &createdAt)
	if err != nil {
		log.Println("db/user.go - CreateUser: ", err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			return "", apperrors.ErrResourceAlreadyExists
		}

		return "", err
	}

	return email, nil

}
