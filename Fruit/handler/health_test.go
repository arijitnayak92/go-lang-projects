package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/affmed/appcontext"
	"gitlab.com/affordmed/affmed/domain"
	"gitlab.com/affordmed/affmed/mock"
	"gitlab.com/affordmed/affmed/util"
	"gitlab.com/affordmed/affmed/validation"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHealthHandler(t *testing.T) {
	cases := map[string]struct {
		want   string
		status int
	}{
		"when database is connected": {
			want:   `{"alive":true,"db_status":"connected"}`,
			status: http.StatusOK,
		},
	}

	appCtx := appcontext.NewAppContext("pg uri")
	pg := mock.NewPostgres()
	u := util.NewUtil()
	d := domain.NewDomain(appCtx, pg, u)
	v := validation.NewValidation(u)
	handler := NewHandler(appCtx, d, v, u)
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
