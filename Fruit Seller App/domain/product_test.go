package domain

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/apperrors"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/mock"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/model"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
)

func TestCreateProduct(t *testing.T) {
	testproduct := model.NewProduct(1, "Orange", 100, "008Image.png", "This is a orange")
	utilp := utils.NewUtil()

	cases := map[string]struct {
		getErr error
		want   int
	}{
		"when product added successfully": {
			getErr: nil,
			want:   1,
		},
		"when product already exist": {
			getErr: apperrors.ErrDuplicateEnrty,
			want:   0,
		},
	}
	query := "INSERT INTO products (id,name,price,imageid,description,createdat,updatedat) VALUES(nextval('prod_id'),$2,$3,$4,$5,$6,$7) RETURNING id;"
	appCtx := appcontext.NewAppContext("postgres uri", "mongo uri")
	dbs, mocks := NewMock()
	repopg := dbs
	mockRepoMongo := new(mock.Mongo)
	for k, v := range cases {
		testDomain := NewDomain(appCtx, repopg, mockRepoMongo, utilp)
		prep := mocks.ExpectQuery(regexp.QuoteMeta(query))
		t.Run(k, func(t *testing.T) {
			if v.getErr == nil {
				prep.WithArgs(testproduct.ID, testproduct.Name, testproduct.Price, testproduct.ImageID, testproduct.Description, time.Now(), time.Now()).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testproduct.ID))
				_, err := testDomain.CreateProduct(testproduct.Name, testproduct.Price, testproduct.ImageID, testproduct.Description)
				assert.Equal(t, v.getErr, err)
			} else if v.getErr == apperrors.ErrDuplicateEnrty {
				prep.WithArgs(testproduct.ID, testproduct.Name, testproduct.Price, testproduct.ImageID, testproduct.Description, time.Now(), time.Now()).WillReturnError(apperrors.ErrDuplicateEnrty)
				_, err := testDomain.CreateProduct(testproduct.Name, testproduct.Price, testproduct.ImageID, testproduct.Description)
				assert.Equal(t, v.getErr, err)
			}

		})
	}

}
