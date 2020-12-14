package db

import (
	"database/sql"
	"log"

	// Postgres driver import
	"github.com/arijitnayak92/taskAfford/Fruit2/apperrors"
	_ "github.com/lib/pq"
)

// ConnectToPostgres func to connect to mongodb
func ConnectToPostgres(connStr string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("postgres-client ", err)
		return nil, err

	}
	//postgres = db
	return db, nil

}

// PostgresHealthCheck to ping database and check for errors
func (repo *Repository) PostgresHealthCheck() error {
	if err := repo.Postgres.db.Ping(); err != nil {
		log.Println("Could not connect to Postgres database", err)
		return apperrors.ErrPostgresConnection
	}
	return nil
}

// PostgresRepo ...
type PostgresRepo struct {
	db *sql.DB
}

// NewPostgresRepo contrructor
// func NewPostgresRepo(database *sql.DB) *PostgresRepo {
// 	return &PostgresRepo{
// 		db: database,
// 	}
// }

//TODO: Change test cases.
