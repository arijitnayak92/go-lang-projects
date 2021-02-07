package db

import (
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
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

// Mongo ...
type Mongo struct {
	DB *mongo.Client
}

// Postgres ...
type Postgres struct {
	DB *sql.DB
}
