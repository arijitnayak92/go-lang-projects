package mock

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo ...
type Mongo struct {
	mock.Mock
}

// Postgres ...
type Postgres struct {
	err error
}

// NewPostgres ...
func NewPostgres(err error) *Postgres {
	return &Postgres{
		err: err,
	}
}

// NewMongo ...
func NewMongo() *Mongo {
	return &Mongo{}
}

// Exec ...
func (p *Postgres) Exec(_ string, _ ...interface{}) (sql.Result, error) {
	return nil, nil
}

// QueryRow ...
func (p *Postgres) QueryRow(_ string, _ ...interface{}) *sql.Row {
	return nil
}

// Ping ...
func (p *Postgres) Ping() error {
	if p.err != nil {
		if p.err.Error() == "Unable to connect to the Postgres database" {
			return p.err
		}
	}

	return nil
}

// Ping ...
func (mock *Mongo) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	args := mock.Called()

	return args.Error(0)
}
