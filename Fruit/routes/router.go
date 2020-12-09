package routes

import (
	"github.com/arijitnayak92/taskAfford/Fruit/handler"
	"github.com/gin-gonic/gin"
)

/*
Route Structure of new routes
*/
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

/*
Routes Array of all available routes
*/
type Routes []Route

// NewRoutes returns all the routes
func NewRoutes(h handler.AppHandler) Routes {
	var routes = Routes{
		Route{
			"Health",
			"GET",
			"/health",
			h.HealthHandler,
		},
	}

	return routes
}

/*
AttachRoutes Attaches routes to the provided server
*/
func AttachRoutes(server *gin.Engine, routes Routes) {
	for _, route := range routes {
		server.
			Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}
