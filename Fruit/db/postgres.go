package db

import (
	"database/sql"
	"gitlab.com/affordmed/affmed/appcontext"
	"log"
)

type PostgresClient interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Ping() error
}

type Postgres struct {
}

func NewPostgres(appCtx *appcontext.AppContext) (*sql.DB, error) {
	DB, err := sql.Open("postgres", appCtx.PostgresURI)
	if err != nil {
		log.Println("Cannot connect to postgres database")
		return nil, err
	}

	return DB, err
}
