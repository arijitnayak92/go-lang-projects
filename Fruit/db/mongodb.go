package db

import (
	"context"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	DB *mongo.Client
}

func NewMongo(appCtx *appcontext.AppContext) (*mongo.Client, *utils.APIError) {
	clientOptions := options.Client().ApplyURI(appCtx.MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, &utils.APIError{
			Message:    "Unable to connect MongoDB !",
			StatusCode: 400,
		}
	}
	return client, nil
}

// To check connection status of mongo db ...
func (repo *DB) CheckMongoAlive() *utils.APIError {
	err := repo.Mongo.DB.Ping(context.TODO(), nil)

	if err != nil {
		return &utils.APIError{
			Message:    "MongoDB not responding !",
			StatusCode: 400,
		}
	}
	return nil
}
