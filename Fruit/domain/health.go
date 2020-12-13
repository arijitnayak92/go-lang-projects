package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// CheckDatabaseHealth ...
func (d *Domain) GetPostgresHealth() bool {
	err := d.appPgDB.Ping()

	if err != nil {
		return false
	}
	return true
}

// CheckDatabaseHealth ...
func (d *Domain) GetMongoHealth() bool {
	err := d.appMongoDB.Ping(context.Background(), readpref.Primary())

	if err != nil {
		return false
	}
	return true
}
