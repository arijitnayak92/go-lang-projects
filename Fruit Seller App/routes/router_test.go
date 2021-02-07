package routes

import (
	"fmt"
	"testing"

	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/domain"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/handler"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/validation"
)

func TestNewRouter(t *testing.T) {
	var appdomain *domain.Domain
	var v validation.AppValidation
	var u utils.AppUtil
	appCtx := appcontext.NewAppContext("pg uri", "mongo uri")
	newHandler := handler.NewHandler(appCtx, appdomain, v, u)
	fmt.Println(newHandler)

	t.Run("Testing for constructor of Router...", func(t *testing.T) {
		if got := NewRouter(newHandler); got.Router == nil {
			t.Errorf("Failed to execuate NewRouter() method. Want gin router instances.Got=%v", got)
		}
	})
}

func TestRoutes(t *testing.T) {
	var appdomain *domain.Domain
	var v validation.AppValidation
	var u utils.AppUtil
	appCtx := appcontext.NewAppContext("pg uri", "mongo uri")
	newHandler := handler.NewHandler(appCtx, appdomain, v, u)
	router := NewRouter(newHandler)

	t.Run("Testing Url Mapping Function...", func(t *testing.T) {
		_, rd := router.Routes()
		got := len(rd)
		want := 3
		if got != want {
			t.Errorf("Routes() Method Failed ! Want %v ,Got = %v", want, got)
		}

	})
}
