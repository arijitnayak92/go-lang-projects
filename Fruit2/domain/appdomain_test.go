package domain

import (
	"testing"

	"gitlab.com/affordmed/fruit-seller-b-backend/apperrors"
	mocks "gitlab.com/affordmed/fruit-seller-b-backend/mocks/repository"

	"github.com/stretchr/testify/assert"
	"gitlab.com/affordmed/fruit-seller-b-backend/db"
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
