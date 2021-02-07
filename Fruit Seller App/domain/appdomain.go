package domain

import (
	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/db"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/model"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
)

// AppDomain ...
type AppDomain interface {
	GetPostgresHealth() bool
	GetMongoHealth() bool
	UserSignup(email string, password string, firstname string, lastname string, role string, cartid string) (bool, error)
	GetUser(email string) (*model.User, error)
	Login(email string, password string) (string, error)
	CreateProduct(name string, price int, imageID string, description string) (int, error)
}

// Domain ...
type Domain struct {
	appCtx *appcontext.AppContext
	pg     db.AppPostgres
	mongo  db.AppMongo
	util   utils.AppUtil
}

// NewDomain ...
func NewDomain(appCtx *appcontext.AppContext, pg db.AppPostgres, mongo db.AppMongo, util utils.AppUtil) *Domain {
	return &Domain{appCtx: appCtx, pg: pg, mongo: mongo, util: util}
}
