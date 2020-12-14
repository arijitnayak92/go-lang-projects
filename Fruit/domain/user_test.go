package domain

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/apperrors"
	"github.com/arijitnayak92/taskAfford/Fruit/mock"
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/stretchr/testify/assert"
)

type Scan struct {
}

type AppScan interface {
	Scan() error
}

func TestGetUser(t *testing.T) {
	testuser := model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}
	query := "SELECT * FROM userdb WHERE email =$1"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)

	prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
	prep.WithArgs(testuser.Email).WillReturnRows(sqlmock.NewRows([]string{"email", "password", "firstname", "lastname", "role", "cartid"}).AddRow(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID))
	testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
	stat, err := testDomain.GetUser(testuser.Email)
	assert.NoError(t, err)
	assert.Equal(t, &testuser, stat)

}

func TestGetUserError(t *testing.T) {
	testuser := model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}
	query := "SELECT * FROM userdb WHERE email =$1"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
	prep.WithArgs(testuser.Email).WillReturnError(apperrors.ErrUserNotFound)
	testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
	_, err := testDomain.GetUser(testuser.Email)
	assert.Equal(t, apperrors.ErrUserNotFound, err)

}

func TestUserSignup(t *testing.T) {
	testuser := model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}
	query2 := "SELECT * FROM userdb WHERE email =$1"
	query := "INSERT INTO userdb (email, password, firstname, lastname,role,cartid) VALUES ($1, $2, $3, $4,$5,$6)"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	prep2 := mocks.ExpectQuery(regexp.QuoteMeta(query2))
	prep2.WithArgs(testuser.Email).WillReturnError(apperrors.ErrUserAlreadyPresent)
	prep := mocks.ExpectExec(regexp.QuoteMeta(query))

	prep.WithArgs(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID).WillReturnResult(sqlmock.NewResult(0, 0))
	testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
	_, err := testDomain.UserSignup(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID)
	assert.NoError(t, err)

}

func TestUserSignupError(t *testing.T) {
	testuser := model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}
	query2 := "SELECT * FROM userdb WHERE email =$1"
	query := "INSERT INTO userdb (email, password, firstname, lastname,role,cartid) VALUES ($1, $2, $3, $4,$5,$6)"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	prep2 := mocks.ExpectQuery(regexp.QuoteMeta(query2))
	prep2.WithArgs(testuser.Email).WithArgs(testuser.Email).WillReturnRows(sqlmock.NewRows([]string{"email", "password", "firstname", "lastname", "role", "cartid"}).AddRow(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID))
	prep := mocks.ExpectExec(regexp.QuoteMeta(query))
	goterror := errors.New("Failed to execute")
	prep.WithArgs(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID).WillReturnResult(sqlmock.NewErrorResult(goterror))
	testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
	stat, _ := testDomain.UserSignup(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID)
	assert.Equal(t, false, stat)

}

func TestLogin(t *testing.T) {
	testuser := model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}
	utilp := utils.NewUtil()
	hashedp, _ := utilp.HashPassword(testuser.Password)
	testuser.Password = hashedp
	query := "SELECT * FROM userdb WHERE email =$1"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
	prep.WithArgs(testuser.Email).WillReturnRows(sqlmock.NewRows([]string{"email", "password", "firstname", "lastname", "role", "cartid"}).AddRow(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID))

	testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
	_, err := testDomain.Login(testuser.Email, "Arijitnayak@92")
	assert.NoError(t, err)

}

func TestLoginErrorUserNotFound(t *testing.T) {
	testuser := model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}
	utilp := utils.NewUtil()
	hashedp, _ := utilp.HashPassword(testuser.Password)
	testuser.Password = hashedp
	query := "SELECT * FROM userdb WHERE email =$1"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
	prep.WithArgs(testuser.Email).WillReturnError(apperrors.ErrUserNotFound)

	testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
	_, err := testDomain.Login(testuser.Email, "Arijitnayak@92")
	assert.Equal(t, apperrors.ErrUserNotFound, err)
}

func TestLoginWrongPassword(t *testing.T) {
	testuser := model.User{
		Email:     "abc@gmail.com",
		Password:  "Arijitnayak@92",
		FirstName: "Arijit",
		LastName:  "Nayak",
		Role:      "User",
		CartID:    "",
	}
	utilp := utils.NewUtil()
	hashedp, _ := utilp.HashPassword(testuser.Password)
	testuser.Password = hashedp
	query := "SELECT * FROM userdb WHERE email =$1"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
	prep.WithArgs(testuser.Email).WillReturnRows(sqlmock.NewRows([]string{"email", "password", "firstname", "lastname", "role", "cartid"}).AddRow(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID))

	testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
	_, err := testDomain.Login(testuser.Email, "Arijitnayak92")
	assert.Equal(t, apperrors.ErrWrongPassword, err)

}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println("An error occured while creating new mock")
	}

	return db, mock

}
