package domain

import (
	"gitlab.com/affordmed/affmed/appcontext"
	"gitlab.com/affordmed/affmed/db"
	"gitlab.com/affordmed/affmed/model"
	"gitlab.com/affordmed/affmed/util"
)

type AppDomain interface {
	GetPostgresHealth() bool
	SignInUser(email, password string) error
	ChangePassword(email, password string) error
	GetUserByEmail(email string) (*model.User, error)
	CreateCategory(name, description string) (*model.Category, error)
}

type Domain struct {
	appCtx *appcontext.AppContext
	pg     db.PostgresClient
	util   util.AppUtil
}

func NewDomain(appCtx *appcontext.AppContext, pg db.PostgresClient, util util.AppUtil) *Domain {
	return &Domain{appCtx: appCtx, pg: pg, util: util}
}
