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
	mock.Mock
}

func (_m *Postgres) Exec(query string, params ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) sql.Result); ok {
		r0 = rf(query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func (_m *Postgres) QueryRow(query string, params ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *sql.Row); ok {
		r0 = rf(query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0

}

func (mock *Postgres) Ping() error {
	args := mock.Called()

	return args.Error(0)
}

// Ping ...
func (mock *Mongo) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	args := mock.Called()

	return args.Error(0)
}
