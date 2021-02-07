package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/handler"
)

// Router ...
type Router struct {
	Router  *gin.Engine
	Handler handler.AppHandler
}

// AppRouter ...
type AppRouter interface {
	Routes() *gin.Engine
}

// NewRouter ...
func NewRouter(h handler.AppHandler) Router {
	rt := gin.Default()
	cosrConfig := cors.DefaultConfig()
	cosrConfig.AllowAllOrigins = true
	cosrConfig.AllowCredentials = true
	cosrConfig.AddAllowMethods("OPTIONS")
	rt.Use(cors.New(cosrConfig))

	return Router{
		Router:  rt,
		Handler: h,
	}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// Routes ...
func (r *Router) Routes() (*gin.Engine, gin.RoutesInfo) {
	r.Router.GET("/health", r.Handler.HealthHandler)
	r.Router.POST("/signup", r.Handler.SignUpUser)
	r.Router.POST("/login", r.Handler.Login)
	routesInfo := r.Router.Routes()
	return r.Router, routesInfo
}
