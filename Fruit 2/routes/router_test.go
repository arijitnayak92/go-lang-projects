package routes

import (
	"fmt"
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/domain"
	"github.com/arijitnayak92/taskAfford/Fruit/handler"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/arijitnayak92/taskAfford/Fruit/validation"
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
		want := 2
		if got != want {
			t.Errorf("Routes() Method Failed ! Want %v ,Got = %v", want, got)
		}

	})
}
