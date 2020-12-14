package domain

import (
	"errors"
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/db"
	"github.com/arijitnayak92/taskAfford/Fruit/mock"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
)

func TestNewDomain(t *testing.T) {
	var appCtx *appcontext.AppContext
	var pg db.AppPostgres
	var mongo db.AppMongo
	var util utils.AppUtil
	newDomain := NewDomain(appCtx, pg, mongo, util)
	if newDomain.appMongoDB != nil {
		t.Errorf("error in NewDomain() constructor")
	}

}

var appDomainMock appsDomainMock
var databaseHealthCheck func() error

type appsDomainMock struct{}

func TestGetPostgresHealth(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	pg := mock.NewPostgres(nil)
	mockRepoMongo := new(mock.Mongo)

	testDomain := NewDomain(appCtx, pg, mockRepoMongo, util)

	postgresErr := testDomain.GetPostgresHealth()

	if postgresErr == false {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

func TestGetPostgresHealthError(t *testing.T) {
	errPostgresConnection := errors.New("Unable to connect to the Postgres database")
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	pg := mock.NewPostgres(errPostgresConnection)
	mockRepoMongo := new(mock.Mongo)

	testDomain := NewDomain(appCtx, pg, mockRepoMongo, util)

	postgresErr := testDomain.GetPostgresHealth()

	if postgresErr == true {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

func TestGetMongoHealth(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	mockRepo := mock.NewPostgres(nil)
	mockRepoMongo := new(mock.Mongo)
	mockRepoMongo.On("Ping").Return(nil)

	testDomain := NewDomain(appCtx, mockRepo, mockRepoMongo, util)

	postgresErr := testDomain.GetMongoHealth()
	mockRepoMongo.AssertExpectations(t)

	if postgresErr == false {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

func TestGetMongoHealthError(t *testing.T) {
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	mockRepo := mock.NewPostgres(nil)
	mockRepoMongo := new(mock.Mongo)
	errPostgresConnection := errors.New("Unable to connect to the Postgres database")
	mockRepoMongo.On("Ping").Return(errPostgresConnection)

	testDomain := NewDomain(appCtx, mockRepo, mockRepoMongo, util)

	postgresErr := testDomain.GetMongoHealth()
	mockRepoMongo.AssertExpectations(t)

	if postgresErr != false {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}
