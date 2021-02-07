package db

import (
	"context"

	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongo ...
func NewMongo(appCtx *appcontext.AppContext) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(appCtx.MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}
	return client, nil
}
