package domain

import (
	"database/sql"
	"log"

	"github.com/arijitnayak92/taskAfford/TODO/config"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	db, err = sql.Open("postgres", config.Db().ConnString())
	if err != nil {
		log.Fatalf("failed to load the database: %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("ping to database failed :%s", err)
	}
	return db
}
