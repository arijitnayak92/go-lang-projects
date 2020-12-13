package db

import (
	"database/sql"

	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	godotenv.Load(".env")
}

// AppDB ...
type AppDB interface {
	PingPostgres() error
	CheckMongoAlive() error
	UserSignup(email string, password string, firstname string, lastname string, role string) (bool, error)
	GetUser(email string) (*model.User, error)
	LoginUser(email string, password string) (string, error)
}

type DB struct {
	Postgres *Postgres
	Mongo    *Mongo
}

func NewDB(postgres *sql.DB, mongo *mongo.Client) *DB {
	return &DB{
		Postgres: &Postgres{
			DB: postgres,
		},
		Mongo: &Mongo{
			DB: mongo,
		},
	}
}
