package mock

import (
	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/db"
	"github.com/stretchr/testify/mock"
)

type Domain struct {
	appCtx *appcontext.AppContext
	appDB  db.AppDB
}

func NewDomain(appCtx *appcontext.AppContext, appDB db.AppDB) *Domain {
	return &Domain{appCtx: appCtx, appDB: appDB}
}

func (d *Domain) GetPostgresHealth() bool {
	return true
}

type MockDomain struct {
	mock.Mock
}

func (mock *MockDomain) GetMongoHealth() (error, error) {
	args := mock.Called()

	return args.Error(0), args.Error(1)
}
