package db

import (
	"context"

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
