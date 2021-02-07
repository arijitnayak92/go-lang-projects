package domain

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/apperrors"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/mock"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/model"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
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
	cases := map[string]struct {
		getErr error
	}{
		"when getuser return a user": {
			getErr: nil,
		},
		"when getuser does not return a user": {
			getErr: apperrors.ErrUserNotFound,
		},
	}
	query := "SELECT * FROM userdb WHERE email =$1"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
			if v.getErr == nil {
				prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
				prep.WithArgs(testuser.Email).WillReturnRows(sqlmock.NewRows([]string{"email", "password", "firstname", "lastname", "role", "cartid"}).AddRow(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID))
			}
			prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
			prep.WithArgs(testuser.Email).WillReturnError(apperrors.ErrUserNotFound)

			_, err := testDomain.GetUser(testuser.Email)
			assert.Equal(t, v.getErr, err)
		})
	}

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
	goterror := errors.New("User already present")
	cases := map[string]struct {
		getErr error
	}{
		"when sign-up done successfully": {
			getErr: nil,
		},
		"when sign-up arise error of user already present": {
			getErr: goterror,
		},
	}
	query2 := "SELECT * FROM userdb WHERE email =$1"
	query := "INSERT INTO userdb (email, password, firstname, lastname,role,cartid) VALUES ($1, $2, $3, $4,$5,$6)"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	var util utils.AppUtil
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			testDomain := NewDomain(appCtx, repopg, mockRepoMongo, util)
			if v.getErr == nil {
				prep2 := mocks.ExpectQuery(regexp.QuoteMeta(query2))
				prep2.WithArgs(testuser.Email).WillReturnError(apperrors.ErrUserAlreadyPresent)
				prep := mocks.ExpectExec(regexp.QuoteMeta(query))

				prep.WithArgs(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID).WillReturnResult(sqlmock.NewResult(0, 0))
			}
			prep2 := mocks.ExpectQuery(regexp.QuoteMeta(query2))
			prep2.WithArgs(testuser.Email).WithArgs(testuser.Email).WillReturnRows(sqlmock.NewRows([]string{"email", "password", "firstname", "lastname", "role", "cartid"}).AddRow(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID))
			prep := mocks.ExpectExec(regexp.QuoteMeta(query))

			prep.WithArgs(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID).WillReturnResult(sqlmock.NewErrorResult(goterror))

			_, err := testDomain.UserSignup(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID)
			assert.Equal(t, v.getErr, err)
		})
	}

}

func TestUserLogin(t *testing.T) {
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

	cases := map[string]struct {
		getErr error
		status int
	}{
		"when user signed in successfully": {
			getErr: nil,
			status: 200,
		},
		"when user not found": {
			getErr: apperrors.ErrUserNotFound,
			status: 404,
		},
		"when password is wrong": {
			getErr: apperrors.ErrWrongPassword,
			status: 401,
		},
	}
	query := "SELECT * FROM userdb WHERE email =$1"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	for k, v := range cases {
		testDomain := NewDomain(appCtx, repopg, mockRepoMongo, utilp)
		prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
		t.Run(k, func(t *testing.T) {
			if v.getErr == nil {
				prep.WithArgs(testuser.Email).WillReturnRows(sqlmock.NewRows([]string{"email", "password", "firstname", "lastname", "role", "cartid"}).AddRow(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID))
				_, err := testDomain.Login(testuser.Email, "Arijitnayak@92")
				assert.Equal(t, v.getErr, err)
			} else if v.getErr == apperrors.ErrUserNotFound {
				prep.WithArgs(testuser.Email).WillReturnError(apperrors.ErrUserNotFound)
				_, err := testDomain.Login(testuser.Email, "Arijitnayak@92")
				assert.Equal(t, v.getErr, err)
			} else {
				prep.WithArgs(testuser.Email).WillReturnRows(sqlmock.NewRows([]string{"email", "password", "firstname", "lastname", "role", "cartid"}).AddRow(testuser.Email, testuser.Password, testuser.FirstName, testuser.LastName, testuser.Role, testuser.CartID))
				_, err := testDomain.Login(testuser.Email, "Arijitnayak92")
				assert.Equal(t, v.getErr, err)
			}

		})
	}

}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println("An error occured while creating new mock")
	}

	return db, mock

}
