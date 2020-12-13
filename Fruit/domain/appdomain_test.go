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
	var want db.Postgres
	want = nil
	if newDomain.appPgDB != want {
		t.Errorf("error in NewDomain() constructor")
	}

}

var appDomainMock appsDomainMock
var databaseHealthCheck func() error

type appsDomainMock struct{}

func TestGetPostgresHealth(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil
	mockRepo := new(mock.Postgres)
	mockRepoMongo := new(mock.Mongo)
	mockRepo.On("Ping").Return(nil)

	testDomain := NewDomain(appCtx, mockRepo, mockRepoMongo, util)

	postgresErr := testDomain.GetPostgresHealth()
	mockRepo.AssertExpectations(t)

	if postgresErr == false {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

func TestGetPostgresHealthError(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil
	mockRepo := new(mock.Postgres)
	mockRepoMongo := new(mock.Mongo)
	errPostgresConnection := errors.New("Unable to connect to the Postgres database")
	mockRepo.On("Ping").Return(errPostgresConnection)

	testDomain := NewDomain(appCtx, mockRepo, mockRepoMongo, util)

	postgresErr := testDomain.GetPostgresHealth()
	mockRepo.AssertExpectations(t)

	if postgresErr != false {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

func TestGetMongoHealth(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil
	mockRepo := new(mock.Postgres)
	mockRepoMongo := new(mock.Mongo)
	mockRepoMongo.On("Ping").Return(nil)

	testDomain := NewDomain(appCtx, mockRepo, mockRepoMongo, util)

	postgresErr := testDomain.GetMongoHealth()
	mockRepo.AssertExpectations(t)

	if postgresErr == false {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

func TestGetMongoHealthError(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil
	mockRepo := new(mock.Postgres)
	mockRepoMongo := new(mock.Mongo)
	errPostgresConnection := errors.New("Unable to connect to the Postgres database")
	mockRepoMongo.On("Ping").Return(errPostgresConnection)

	testDomain := NewDomain(appCtx, mockRepo, mockRepoMongo, util)

	postgresErr := testDomain.GetMongoHealth()
	mockRepo.AssertExpectations(t)

	if postgresErr != false {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

// func TestGetUser(t *testing.T) {
// 	testuser := &model.User{
// 		Email:     "abc@gmail.com",
// 		Password:  "Arijitnayak@92",
// 		FirstName: "Arijit",
// 		LastName:  "Nayak",
// 		Role:      "User",
// 		CartID:    "",
// 	}
// 	// var user *model.User
// 	var appCtx *appcontext.AppContext
// 	var util utils.AppUtil
// 	mockRepo := new(mock.AppPostgresDB)
// 	mockRepoMongo := new(mock.AppMongoDB)
// 	//mockDomain := new(mock.Domain)
// 	errP := errors.New("Not found")
// 	query := "SELECT * FROM users WHERE email =$1"
// 	mockRepo.On("QueryRow", query, testuser.Email).Return(nil, errP)
//
// 	//mockDomain.On("GetUser", testuser.Email).Return(nil, errP)
// 	testDomain := NewDomain(appCtx, mockRepo, mockRepoMongo, util)
//
// 	_, err := testDomain.GetUser(testuser.Email)
// 	mockRepo.AssertExpectations(t)
// 	if err != nil {
// 		t.Errorf("DatabaseHealthCheck not throwing error")
// 	}
//
// }
