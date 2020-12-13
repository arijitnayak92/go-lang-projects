package domain

import (
	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/db"
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
)

type AppDomain interface {
	CheckDatabaseHealth() (error, error)
	UserSignup(email string, password string, firstname string, lastname string, role string) (bool, error)
	GetUser(email string) (*model.User, error)
	Login(email string, password string) (string, error)
}

type Domain struct {
	appCtx *appcontext.AppContext
	appDB  db.AppDB
	util   utils.AppUtil
}

func NewDomain(appCtx *appcontext.AppContext, appDB db.AppDB, util utils.AppUtil) *Domain {
	return &Domain{appCtx: appCtx, appDB: appDB, util: util}
}
