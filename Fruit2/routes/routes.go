package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/fruit-seller-b-backend/handler"
)

// Router : struct contains router of type *gin.Engine
type Router struct {
	Router *gin.Engine
	//version string
}

// NewRouter : Constructor to return *Router struct
func NewRouter() *Router {
	r := gin.Default()
	cosrConfig := cors.DefaultConfig()
	cosrConfig.AllowAllOrigins = true
	cosrConfig.AllowCredentials = true
	cosrConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(cosrConfig))
	return &Router{
		Router: r, //TODO: add version
	}
}

// AppRouter : Interface of routes package
type AppRouter interface {
	SetupRoutes(appHandler handler.AppHandler, userHandler handler.UserHandler) gin.RoutesInfo
}

// SetupRoutes :To set up all the http routes of the app
func (r *Router) SetupRoutes(appHandler handler.AppHandler, userHandler handler.UserHandler) gin.RoutesInfo {

	r.Router.GET("/health", appHandler.GetAppHealth)
	r.Router.POST("/signup", userHandler.AddUser)
	routesInfo := r.Router.Routes()

	return routesInfo
}

// TODO: take v1 as property variable.
