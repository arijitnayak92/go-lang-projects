package domain

import (
	"errors"
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/apperrors"
	"github.com/arijitnayak92/taskAfford/Fruit/db"
	"github.com/arijitnayak92/taskAfford/Fruit/mock"
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewDomain(t *testing.T) {
	var appCtx *appcontext.AppContext
	var appRepo db.AppDB
	var util utils.AppUtil
	newDomain := NewDomain(appCtx, appRepo, util)
	var want db.AppDB
	want = nil
	if newDomain.appDB != want {
		t.Errorf("error in NewDomain() constructor")
	}

}

var appDomainMock appsDomainMock
var databaseHealthCheck func() error

type appsDomainMock struct{}

func TestDatabaseHealthCheck(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil
	mockRepo := new(mock.MockRepository)
	mockRepo.On("PingPostgres").Return(nil)
	mockRepo.On("CheckMongoAlive").Return(nil)
	testDomain := NewDomain(appCtx, mockRepo, util)

	postgresErr, mongoErr := testDomain.CheckDatabaseHealth()
	mockRepo.AssertExpectations(t)

	assert.Nil(t, postgresErr)
	assert.Nil(t, mongoErr)

}

func TestDatabaseHealthErrorPostgres(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil
	errPostgresConnection := errors.New("Unable to connect to the Postgres database")

	mockRepo := new(mock.MockRepository)
	mockRepo.On("PingPostgres").Return(errPostgresConnection)
	mockRepo.On("CheckMongoAlive").Return(nil)
	testDomain := NewDomain(appCtx, mockRepo, util)

	postgresErr, mongoErr := testDomain.CheckDatabaseHealth()

	if postgresErr == nil {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}
	if mongoErr != nil {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

func TestDatabaseHealthErrorMongo(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil
	errMongoConnection := errors.New("Unable to connect to the Mongo database")
	mockRepo := new(mock.MockRepository)
	mockRepo.On("PingPostgres").Return(nil)
	mockRepo.On("CheckMongoAlive").Return(errMongoConnection)
	testDomain := NewDomain(appCtx, mockRepo, util)

	postgresErr, mongoErr := testDomain.CheckDatabaseHealth()

	if postgresErr != nil {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}
	if mongoErr == nil {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}
}

func TestGetUserSignUpError(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil

	testuser := &model.User{
		Email:     "a@gmail.com",
		Password:  "abc",
		FirstName: "abc",
		LastName:  "abc",
		Role:      "user",
	}

	mockRepo := new(mock.MockRepository)

	mockRepo.On("GetUser").Return(nil, apperrors.ErrUserNotFound)
	testDomain := NewDomain(appCtx, mockRepo, util)

	status, _ := testDomain.UserSignup(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role)

	if status == true {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}
}

func TestGetUserSignUp(t *testing.T) {
	var appCtx *appcontext.AppContext
	var util utils.AppUtil

	testuser := &model.User{
		Email:     "a@gmail.com",
		Password:  "abc",
		FirstName: "abc",
		LastName:  "abc",
		Role:      "user",
	}
	var user *model.User
	mockRepo := new(mock.MockRepository)

	mockRepo.On("GetUser").Return(user, nil)
	testDomain := NewDomain(appCtx, mockRepo, util)

	status, _ := testDomain.UserSignup(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role)

	if status != false {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}
}
