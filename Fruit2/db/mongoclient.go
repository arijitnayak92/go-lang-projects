package db

import (
	"context"
	"log"

	"github.com/arijitnayak92/taskAfford/Fruit2/apperrors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var client *mongo.Client

// ConnectToMongo func to connect to mongodb
func ConnectToMongo(connStr string) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(connStr)

	clientInstance, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println("Error in connection", err)
		return nil, err
	}
	//client = clientInstance
	return clientInstance, nil

}

// MongoHealthCheck to ping database and check for errors
func (repo *Repository) MongoHealthCheck() error {
	err := repo.Mongo.db.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Could not connect to Mongo database", err)
		return apperrors.ErrMongoConnection
	}
	return nil
}

// MongoRepo struct having mongo client
type MongoRepo struct {
	db *mongo.Client
}

// NewMongoRepo function to create instance
// func NewMongoRepo(database *mongo.Client) *MongoRepo {
// 	return &MongoRepo{
// 		db: database,
// 	}
// }

//TODO: change test cases.
