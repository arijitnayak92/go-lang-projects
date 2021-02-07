package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GetPostgresHealth ...
func (d *Domain) GetPostgresHealth() bool {
	err := d.pg.Ping()

	if err != nil {
		return false
	}
	return true
}

// GetMongoHealth ...
func (d *Domain) GetMongoHealth() bool {
	err := d.mongo.Ping(context.Background(), readpref.Primary())

	if err != nil {
		return false
	}
	return true
}
