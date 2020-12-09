package mock

import (
	"database/sql"
	"gitlab.com/affordmed/affmed/appcontext"
	"gitlab.com/affordmed/affmed/model"
	"gitlab.com/affordmed/affmed/util"
)

type Domain struct {
	appCtx *appcontext.AppContext
	pg     *sql.DB
	util   util.AppUtil
}

func NewDomain(appCtx *appcontext.AppContext, pg *sql.DB, util util.AppUtil) *Domain {
	return &Domain{appCtx: appCtx, pg: pg, util: util}
}

func (d *Domain) GetPostgresHealth() bool {
	return true
}

func (d *Domain) SignInUser(email, password string) error {
	return nil
}

func (d *Domain) ChangePassword(email, password string) error {
	return nil
}

func (d *Domain) GetUserByEmail(email string) (*model.User, error) {
	return model.NewUser("captain@avengers.com", "avengers assemble", "::1", false), nil
}
