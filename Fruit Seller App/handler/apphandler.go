package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/domain"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/validation"
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
