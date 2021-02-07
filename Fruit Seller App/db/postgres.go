package db

import (
	"database/sql"
	"log"

	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
)

// NewPostgres ...
func NewPostgres(appCtx *appcontext.AppContext) (*sql.DB, error) {
	DB, err := sql.Open("postgres", appCtx.PostgresURI)
	if err != nil {
		log.Println("Cannot connect to postgres database")
		return nil, err
	}

	return DB, err
}
