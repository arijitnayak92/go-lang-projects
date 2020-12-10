package routes

import (
	"github.com/arijitnayak92/taskAfford/Fruit/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router  *gin.Engine
	Handler handler.AppHandler
}

type AppRouter interface {
	Routes() *gin.Engine
}

//...
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

//"Routes ..."
func (r *Router) Routes() (*gin.Engine, gin.RoutesInfo) {
	r.Router.GET("/", r.Handler.HealthHandler)
	routesInfo := r.Router.Routes()
	return r.Router, routesInfo
}
