package domain

import (
	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/db"
)

type AppDomain interface {
	GetPostgresHealth() bool
	GetMongoHealth() bool
}

type Domain struct {
	appCtx *appcontext.AppContext
	appDB  db.AppDB
}

func NewDomain(appCtx *appcontext.AppContext, appDB db.AppDB) *Domain {
	return &Domain{appCtx: appCtx, appDB: appDB}
}
