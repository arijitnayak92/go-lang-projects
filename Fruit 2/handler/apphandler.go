package handler

import (
	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/domain"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/arijitnayak92/taskAfford/Fruit/validation"
	"github.com/gin-gonic/gin"
)

// AppHandler ...
type AppHandler interface {
	HealthHandler(c *gin.Context)
	SignUpUser(c *gin.Context)
	Login(c *gin.Context)
}

// Handler ...
type Handler struct {
	appContext *appcontext.AppContext
	domain     domain.AppDomain
	validation validation.AppValidation
	util       utils.AppUtil
}

// NewHandler returns new instance of Handler
func NewHandler(appContext *appcontext.AppContext, domain domain.AppDomain, v validation.AppValidation, util utils.AppUtil) *Handler {
	return &Handler{appContext: appContext, domain: domain, validation: v, util: util}
}
