package domain

import (
	"errors"
	"reflect"
	"testing"

	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/db"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/mock"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
)

func TestNewDomain(t *testing.T) {
	cases := map[string]struct {
		want       *Domain
		appContext *appcontext.AppContext
		pg         db.AppPostgres
		mongo      db.AppMongo
		util       utils.AppUtil
	}{
		"when every parameter is passed as nil": {
			want:       NewDomain(nil, nil, nil, nil),
			appContext: nil,
			pg:         nil,
			mongo:      nil,
			util:       nil,
		},
		"when every parameter is passed": {
			want: NewDomain(
				appcontext.NewAppContext("postgres uri", "mongo uri"),
				mock.NewPostgres(nil),
				mock.NewMongo(),
				utils.NewUtil(),
			),
			appContext: appcontext.NewAppContext("postgres uri", "mongo uri"),
			pg:         mock.NewPostgres(nil),
			mongo:      mock.NewMongo(),
			util:       utils.NewUtil(),
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			h := NewDomain(v.appContext, v.pg, v.mongo, v.util)
			if !reflect.DeepEqual(v.want, h) {
				t.Errorf("domain mismatched\nwant: %v\ngot:%v\n", v.want, h)
			}
		})
	}

}

var appDomainMock appsDomainMock
var databaseHealthCheck func() error

type appsDomainMock struct{}

func TestGetPostgresHealth(t *testing.T) {
	cases := map[string]struct {
		want    bool
		pingErr error
	}{
		"when postgres database connection is successful": {
			want:    true,
			pingErr: nil,
		},
		"when postgres database connection is unsuccessful": {
			want:    false,
			pingErr: errors.New("Unable to connect to the Postgres database"),
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
			pg := mock.NewPostgres(v.pingErr)
			var util utils.AppUtil
			mockRepoMongo := new(mock.Mongo)

			testDomain := NewDomain(appCtx, pg, mockRepoMongo, util)

			got := testDomain.GetPostgresHealth()

			if got != v.want {
				t.Errorf("postgres health return mismatched\nwant: %v\ngot: %v", v.want, got)
			}
		})
	}

}

func TestGetMongoHealth(t *testing.T) {
	cases := map[string]struct {
		want    bool
		pingErr error
	}{
		"when mongo database connection is successful": {
			want:    true,
			pingErr: nil,
		},
		"when mongo database connection is unsuccessful": {
			want:    false,
			pingErr: errors.New("Unable to connect to the mongo database"),
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
			pg := mock.NewPostgres(nil)
			var util utils.AppUtil
			mockRepoMongo := new(mock.Mongo)
			mockRepoMongo.On("Ping").Return(v.pingErr)
			testDomain := NewDomain(appCtx, pg, mockRepoMongo, util)

			got := testDomain.GetMongoHealth()

			if got != v.want {
				t.Errorf("mongo health return mismatched\nwant: %v\ngot: %v", v.want, got)
			}
		})
	}
}
