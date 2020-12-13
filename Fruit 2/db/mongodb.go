package db

import (
	"context"
	"fmt"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	DB *mongo.Client
}

func NewMongo(appCtx *appcontext.AppContext) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(appCtx.MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}
	return client, nil
}

// To check connection status of mongo db ...
func (repo *DB) CheckMongoAlive() error {
	fmt.Println("Ping Mongo Called")
	err := repo.Mongo.DB.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}
	return nil
}
