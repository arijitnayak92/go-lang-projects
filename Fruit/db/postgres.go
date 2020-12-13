package db

import (
	"database/sql"
	"log"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(appCtx *appcontext.AppContext) (*sql.DB, error) {
	DB, err := sql.Open("postgres", appCtx.PostgresURI)
	if err != nil {
		log.Println("Cannot connect to postgres database")
		return nil, err
	}

	return DB, err
}
