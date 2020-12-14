package domain

import (
	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/db"
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
)

// AppDomain ...
type AppDomain interface {
	GetPostgresHealth() bool
	GetMongoHealth() bool
	UserSignup(email string, password string, firstname string, lastname string, role string, cartid string) (bool, error)
	GetUser(email string) (*model.User, error)
	Login(email string, password string) (string, error)
}

// Domain ...
type Domain struct {
	appCtx     *appcontext.AppContext
	appPgDB    db.AppPostgres
	appMongoDB db.AppMongo
	util       utils.AppUtil
}

// NewDomain ...
func NewDomain(appCtx *appcontext.AppContext, appPgDB db.AppPostgres, appMongoDB db.AppMongo, util utils.AppUtil) *Domain {
	return &Domain{appCtx: appCtx, appPgDB: appPgDB, appMongoDB: appMongoDB, util: util}
}
