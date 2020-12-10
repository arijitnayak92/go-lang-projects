package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/domain"
	"github.com/arijitnayak92/taskAfford/Fruit/mock"
	"github.com/gin-gonic/gin"
)

func executeRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestNewHandlers(t *testing.T) {
	var appdomain *domain.Domain
	appCtx := appcontext.NewAppContext("pg uri", "mongo uri")
	newHandler := NewHandler(appCtx, appdomain)

	if newHandler.domain == nil {
		t.Errorf("Error in NewHandlers() constructor")
	}

}

func TestHandlerHealthHandler(t *testing.T) {
	t.Log("When server and both the DB connected ")
	{
		cases := map[string]struct {
			want   string
			status int
		}{
			"when database is connected": {
				want:   `{"mongoIsAlive":true,"postgresIsAlive":true,"serverIsAlive":true}`,
				status: http.StatusOK,
			},
		}

		appCtx := appcontext.NewAppContext("pg uri", "mongo uri")
		pg := mock.NewPostgres()
		mongo := mock.NewMongo()
		db := mock.NewDB(pg, mongo)
		d := domain.NewDomain(appCtx, db)
		handler := NewHandler(appCtx, d)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		handler.HealthHandler(c)

		for k, v := range cases {
			var got gin.H
			err := json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}

			t.Run(k, func(t *testing.T) {
				if status := w.Code; status != v.status {
					t.Errorf("handler returned wrong status code: got %v want %v",
						status, http.StatusOK)
				}

				if w.Body.String() != v.want {
					t.Errorf("handler returned unexpected body: got %v want %v",
						w.Body.String(), v.want)
				}
			})
		}
	}

}
