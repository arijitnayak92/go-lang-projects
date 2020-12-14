package db

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/arijitnayak92/taskAfford/Fruit2/models"

	"github.com/DATA-DOG/go-sqlmock"
)

var u = models.User{
	Email:           "abc@gmail.com",
	Password:        "test123",
	ConfirmPassword: "test123",
	FirstName:       "test",
	LastName:        "test",
	Role:            "user",
	CreatedAt:       time.Now(),
	CartID:          1,
}

func TestPostgresRepo_CreateUser(t *testing.T) {
	t.Run("WHen returning nil error", func(t *testing.T) {
		db, mock := NewMock()

		repo := &PostgresRepo{db}
		query := `INSERT INTO users (email,cartid,password,firstname,lastname,role,createdat) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING email;`
		prep := mock.ExpectQuery(regexp.QuoteMeta(query))
		prep.WithArgs(u.Email, u.CartID, u.Password, u.FirstName, u.LastName, u.Role, u.CreatedAt).WillReturnRows(sqlmock.NewRows([]string{"email"}).AddRow(u.Email))

		res, err := repo.CreateUser(u)
		if err != nil {
			t.Errorf("Error in CreateUser: %v", err)
		}

		if res == false {
			t.Error("Error in email :", res)
		}
	})

	t.Run("WHen returning not nil error", func(t *testing.T) {
		db, mock := NewMock()
		errorDb := errors.New("DB error")
		repo := &PostgresRepo{db}
		query := `INSERT INTO users (email,cartid,password,firstname,lastname,role,createdat) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING email;`
		prep := mock.ExpectQuery(regexp.QuoteMeta(query))
		prep.WithArgs(u.Email, u.CartID, u.Password, u.FirstName, u.LastName, u.Role, u.CreatedAt).WillReturnError(errorDb)

		_, err := repo.CreateUser(u)
		if err == nil {
			t.Errorf("Error in CreateUser: %v", err)
		}

	})

}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println("An error occured while creating new mock")
	}

	return db, mock

}
