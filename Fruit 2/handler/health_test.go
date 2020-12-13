package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/domain"
	"github.com/arijitnayak92/taskAfford/Fruit/mock"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/arijitnayak92/taskAfford/Fruit/validation"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func executeRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
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
	var v validation.AppValidation
	t.Run("1: When Both Database Connection passes!", func(t *testing.T) {
		mockDomain := new(mock.MockDomain)
		mockDomain.On("CheckDatabaseHealth").Return(nil, nil)

		testHandler := NewHandler(appCtx, mockDomain, v, u)

		// appRouter := routes.NewRouter()
		// appRouter.SetupRoutes(testHandler)
		router := gin.Default()
		router.GET("/", testHandler.HealthHandler)

		w := executeRequest(router, "GET", "/")
		// testHandler.HealthHandler()
		mockDomain.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]bool
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		body := gin.H{
			"postgresIsAlive": true,
			"mongoIsAlive":    true,
			"serverIsAlive":   true,
		}
		value1, _ := response["postgresIsAlive"]
		assert.Equal(t, body["postgresIsAlive"], value1)

		value2, _ := response["serverIsAlive"]
		assert.Equal(t, body["serverIsAlive"], value2)

		value3, _ := response["mongoIsAlive"]
		assert.Equal(t, body["mongoIsAlive"], value3)

	})

	t.Run("2: When Postgres Database Connection fails Mongo Passes", func(t *testing.T) {
		mockDomain := new(mock.MockDomain)
		errPostgresConnection := errors.New("Unable to connect to the Postgres database")
		mockDomain.On("CheckDatabaseHealth").Return(errPostgresConnection, nil)

		testHandler := NewHandler(appCtx, mockDomain, v, u)

		router := gin.Default()
		router.GET("/", testHandler.HealthHandler)

		w := executeRequest(router, "GET", "/")

		mockDomain.AssertExpectations(t)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]bool
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		body := gin.H{
			"postgresIsAlive": false,
			"mongoIsAlive":    true,
			"serverIsAlive":   true,
		}
		value1, _ := response["postgresIsAlive"]
		assert.Equal(t, body["postgresIsAlive"], value1)

		value2, _ := response["serverIsAlive"]
		assert.Equal(t, body["serverIsAlive"], value2)

		value3, _ := response["mongoIsAlive"]
		assert.Equal(t, body["mongoIsAlive"], value3)

	})

	t.Run("3: When Mongo Database Connection Fails!..Postgres Passes", func(t *testing.T) {
		mockDomain := new(mock.MockDomain)
		errMongoConnection := errors.New("Unable to connect to the Mongo database")
		mockDomain.On("CheckDatabaseHealth").Return(nil, errMongoConnection)

		testHandler := NewHandler(appCtx, mockDomain, v, u)

		// appRouter := routes.NewRouter()
		// appRouter.SetupRoutes(testHandler)
		router := gin.Default()
		router.GET("/", testHandler.HealthHandler)

		w := executeRequest(router, "GET", "/")
		// testHandler.HealthHandler()
		mockDomain.AssertExpectations(t)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]bool
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		body := gin.H{
			"postgresIsAlive": true,
			"mongoIsAlive":    false,
			"serverIsAlive":   true,
		}
		value1, _ := response["postgresIsAlive"]
		assert.Equal(t, body["postgresIsAlive"], value1)

		value2, _ := response["serverIsAlive"]
		assert.Equal(t, body["serverIsAlive"], value2)

		value3, _ := response["mongoIsAlive"]
		assert.Equal(t, body["mongoIsAlive"], value3)

	})

	t.Run("4: When Both Database Connection Fails!", func(t *testing.T) {
		mockDomain := new(mock.MockDomain)
		errMongoConnection := errors.New("Unable to connect to the Mongo database")
		errPostgresConnection := errors.New("Unable to connect to the Postgres database")
		mockDomain.On("CheckDatabaseHealth").Return(errPostgresConnection, errMongoConnection)

		testHandler := NewHandler(appCtx, mockDomain, v, u)

		// appRouter := routes.NewRouter()
		// appRouter.SetupRoutes(testHandler)
		router := gin.Default()
		router.GET("/", testHandler.HealthHandler)

		w := executeRequest(router, "GET", "/")
		// testHandler.HealthHandler()
		mockDomain.AssertExpectations(t)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]bool
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		body := gin.H{
			"postgresIsAlive": false,
			"mongoIsAlive":    false,
			"serverIsAlive":   true,
		}
		value1, _ := response["postgresIsAlive"]
		assert.Equal(t, body["postgresIsAlive"], value1)

		value2, _ := response["serverIsAlive"]
		assert.Equal(t, body["serverIsAlive"], value2)

		value3, _ := response["mongoIsAlive"]
		assert.Equal(t, body["mongoIsAlive"], value3)

	})

}
