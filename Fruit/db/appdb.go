package db

import (
	"database/sql"
	"os"

	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	godotenv.Load(".env")
}

// AppDB ...
type AppDB interface {
	PingPostgres() *utils.APIError
	CheckMongoAlive() *utils.APIError
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

func GetConnectionUri(contype string) string {
	uri := ""
	if contype == "mongo" {
		uri = os.Getenv("MONGODB_URI")
	}
	if contype == "postgres" {
		uri = os.Getenv("POSTGRES_URI")
	}
	return uri
}
