package routes

import (
	"testing"

	"github.com/arijitnayak92/taskAfford/Fruit2/handler"
)

func TestSetupRoutes(t *testing.T) {
	var (
		apphandler  *handler.Handler
		userhandler *handler.User
	)
	appRouter := NewRouter()
	t.Run("When function returns router!!", func(t *testing.T) {
		r := appRouter.SetupRoutes(apphandler, userhandler)
		got := len(r)
		want := 2
		if got != want {
			t.Errorf("SetupRoutes() = %v", got)
		}
	})

}

func TestNewRouter(t *testing.T) {
	t.Run("When Router is returned", func(t *testing.T) {
		if got := NewRouter(); got == nil {
			t.Errorf("NewRouter() = %v, want *gin.Router", got)
		}
	})

}
