package db

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository ...
type ProductRepository interface{}

// CartRepository ...
type CartRepository interface{}

// AppRepository ...
type AppRepository interface {
	PostgresHealthCheck() error
	MongoHealthCheck() error
}

// Repository having databases
type Repository struct {
	Postgres *PostgresRepo
	Mongo    *MongoRepo
}

// NewRepository Constructor to initialize repository
func NewRepository(postgres *sql.DB, mongo *mongo.Client) *Repository {
	return &Repository{
		Postgres: &PostgresRepo{
			db: postgres,
		},
		Mongo: &MongoRepo{
			db: mongo,
		},
	}
}
