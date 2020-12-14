package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/arijitnayak92/taskAfford/Fruit2/apperrors"

	"github.com/arijitnayak92/taskAfford/Fruit2/models"
)

const queryInsertUser = "INSERT INTO users (email,cartid,password,firstname,lastname,role,createdat) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING email;"

// UserRepository ...
type UserRepository interface {
	CreateUser(m models.User) (bool, error)
}

// CreateUser function to insert user into the database
func (repo *PostgresRepo) CreateUser(user models.User) (bool, error) {
	var (
		email string
	)
	//userExistingError:='pq: duplicate key value violates unique constraint "users_pkey"'
	err := repo.db.QueryRow(queryInsertUser, user.Email, user.CartID, user.Password, user.FirstName, user.LastName, user.Role, user.CreatedAt).Scan(&email)
	if err != nil {
		fmt.Println(err)
		log.Println("db/user.go - CreateUser: ", err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			return false, apperrors.ErrResourceAlreadyExists
		}

		return false, err
	}

	return true, nil

}
