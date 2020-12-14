package domain

import (
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit2/apperrors"
	mocks "github.com/arijitnayak92/taskAfford/Fruit2/mocks/repository"

	"github.com/arijitnayak92/taskAfford/Fruit2/db"
	"github.com/stretchr/testify/assert"
)

var appDomainMock appsDomainMock
var databaseHealthCheck func() error

type appsDomainMock struct{}

func TestNewDomain(t *testing.T) {
	var appRepo *db.Repository
	newDomain := NewDomain(appRepo)

	if newDomain.appRepository == nil {
		t.Errorf("error in NewDomain() constructor")
	}

}

func TestDatabaseHealthCheck(t *testing.T) {

	mockRepo := new(mocks.MockRepository)
	mockRepo.On("PostgresHealthCheck").Return(nil)
	mockRepo.On("MongoHealthCheck").Return(nil)
	testDomain := NewDomain(mockRepo)

	postgresErr, mongoErr := testDomain.DatabaseHealthCheck()
	mockRepo.AssertExpectations(t)

	assert.Nil(t, postgresErr)
	assert.Nil(t, mongoErr)

}

func TestDatabaseHealthErrorPostgres(t *testing.T) {

	mockRepo := new(mocks.MockRepository)
	mockRepo.On("PostgresHealthCheck").Return(apperrors.ErrPostgresConnection)
	mockRepo.On("MongoHealthCheck").Return(nil)
	testDomain := NewDomain(mockRepo)

	postgresErr, mongoErr := testDomain.DatabaseHealthCheck()

	if postgresErr == nil {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}
	if mongoErr != nil {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}

}

func TestDatabaseHealthErrorMongo(t *testing.T) {

	mockRepo := new(mocks.MockRepository)
	mockRepo.On("PostgresHealthCheck").Return(nil)
	mockRepo.On("MongoHealthCheck").Return(apperrors.ErrMongoConnection)
	testDomain := NewDomain(mockRepo)

	postgresErr, mongoErr := testDomain.DatabaseHealthCheck()

	if postgresErr != nil {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}
	if mongoErr == nil {
		t.Errorf("DatabaseHealthCheck not throwing error")
	}
}
