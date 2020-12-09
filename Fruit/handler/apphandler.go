package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/affmed/appcontext"
	"gitlab.com/affordmed/affmed/domain"
	"gitlab.com/affordmed/affmed/util"
	"gitlab.com/affordmed/affmed/validation"
)

// AppHandler ...
type AppHandler interface {
	HealthHandler(c *gin.Context)
	SignInHandler(c *gin.Context)
	ChangePasswordHandler(c *gin.Context)
	CreateCategoryHandler(c *gin.Context)
}

// Handler ...
type Handler struct {
	appContext *appcontext.AppContext
	domain     domain.AppDomain
	validation validation.AppValidation
	util       util.AppUtil
}

// NewHandler returns new instance of Handler
func NewHandler(appContext *appcontext.AppContext, domain domain.AppDomain, v validation.AppValidation, u util.AppUtil) *Handler {
	return &Handler{appContext: appContext, domain: domain, validation: v, util: u}
}
