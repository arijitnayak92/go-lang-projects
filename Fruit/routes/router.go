package routes

import (
	"github.com/arijitnayak92/taskAfford/Fruit/handler"
	"github.com/gin-gonic/gin"
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
	r.Router.GET("/", r.Handler.HealthHandler)
	r.Router.POST("/signup", r.Handler.SignUpUser)
	r.Router.POST("/login", r.Handler.Login)
	routesInfo := r.Router.Routes()
	return r.Router, routesInfo
}
