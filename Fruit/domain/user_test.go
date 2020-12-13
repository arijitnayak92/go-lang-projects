package domain

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/mock"
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
)

func TestGetUser(t *testing.T) {
	testuser := model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}
	query := "SELECT * FROM users WHERE email =$1"

	// var user *model.User
	// var rowss *sql.Row
	// rowss.Scan(&testuser.Email, &testuser.Password, &testuser.FirstName, &testuser.LastName, &testuser.Role, &testuser.CartID)
	var appCtx *appcontext.AppContext
	var util utils.AppUtil
	mockRepo := new(mock.AppPostgresDB)
	mockRepoMongo := new(mock.AppMongoDB)
	//mockDomain := new(mock.Domain)
	//errP := errors.New("Not found")
	//query := "SELECT * FROM users WHERE email =$1"
	mockRepo.On("QueryRow", query, testuser.Email).Return(testuser)

	//mockDomain.On("GetUser", testuser.Email).Return(nil, errP)
	testDomain := NewDomain(appCtx, mockRepo, mockRepoMongo, util)

	_, err := testDomain.GetUser(testuser.Email)
	mockRepo.AssertExpectations(t)
	if err != nil {
		t.Errorf("Error in get user")
	}

}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println("An error occured while creating new mock")
	}

	return db, mock

}
