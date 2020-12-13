package db

import (
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// AppPostgres ...
type AppPostgres interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Ping() error
}

// AppMongo ...
type AppMongo interface {
	Ping(ctx context.Context, rp *readpref.ReadPref) error
}
