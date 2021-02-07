package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/domain"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/mock"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/validation"
)

func executeRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestNewHandlers(t *testing.T) {
	var appdomain *domain.Domain
	var u utils.AppUtil
	var v validation.AppValidation
	appCtx := appcontext.NewAppContext("pg uri", "mongo uri")
	newHandler := NewHandler(appCtx, appdomain, v, u)

	if newHandler.domain == nil {
		t.Errorf("Error in NewHandlers() constructor")
	}

}

func TestHandler_GetAppHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var appCtx *appcontext.AppContext
	var u utils.AppUtil
	var vl validation.AppValidation
	cases := map[string]struct {
		want      string
		status    int
		pingMongo bool
		pingPg    bool
	}{
		"When Both Database Connection passes": {
			want:      `{"mongoIsAlive":true,"postgresIsAlive":true,"serverIsAlive":true}`,
			status:    http.StatusOK,
			pingMongo: true,
			pingPg:    true,
		},
		"When Postgres Database Connection fails Mongo Passes": {
			want:      `{"mongoIsAlive":true,"postgresIsAlive":false,"serverIsAlive":true}`,
			status:    http.StatusInternalServerError,
			pingMongo: true,
			pingPg:    false,
		},
		"When Mongo Database Connection Fails!..Postgres Passes": {
			want:      `{"mongoIsAlive":false,"postgresIsAlive":true,"serverIsAlive":true}`,
			status:    http.StatusInternalServerError,
			pingMongo: false,
			pingPg:    true,
		},
		"When Both Database Connection Fails!": {
			want:      `{"mongoIsAlive":false,"postgresIsAlive":false,"serverIsAlive":true}`,
			status:    http.StatusInternalServerError,
			pingMongo: false,
			pingPg:    false,
		},
	}

	for k, v := range cases {
		mockDomain := new(mock.Domain)
		t.Run(k, func(t *testing.T) {

			mockDomain.On("GetPostgresHealth").Return(v.pingPg)
			mockDomain.On("GetMongoHealth").Return(v.pingMongo)

			testHandler := NewHandler(appCtx, mockDomain, vl, u)

			router := gin.Default()
			router.GET("/health", testHandler.HealthHandler)

			w := executeRequest(router, "GET", "/health", nil)

			mockDomain.AssertExpectations(t)

			assert.Equal(t, v.status, w.Code)
			assert.Equal(t, v.want, w.Body.String())
		})
	}

}
