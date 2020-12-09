package handler

import (
	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/domain"
	"github.com/gin-gonic/gin"
)

// AppHandler ...
type AppHandler interface {
	HealthHandler(c *gin.Context)
}

// Handler ...
type Handler struct {
	appContext *appcontext.AppContext
	domain     domain.AppDomain
}

// NewHandler returns new instance of Handler
func NewHandler(appContext *appcontext.AppContext, domain domain.AppDomain) *Handler {
	return &Handler{appContext: appContext, domain: domain}
}
