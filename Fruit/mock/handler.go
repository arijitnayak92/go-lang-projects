package mock

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/affmed/appcontext"
	"gitlab.com/affordmed/affmed/domain"
	"gitlab.com/affordmed/affmed/util"
	"gitlab.com/affordmed/affmed/validation"
	"net/http"
)

type Handler struct {
	appContext *appcontext.AppContext
	domain     domain.AppDomain
	validation validation.AppValidation
	util       util.AppUtil
}

func NewHandler(appContext *appcontext.AppContext, domain domain.AppDomain, validation validation.AppValidation, util util.AppUtil) *Handler {
	return &Handler{appContext: appContext, domain: domain, validation: validation, util: util}
}

func (h *Handler) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "mocked handlers called",
	})
}

func (h *Handler) SignInHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "mocked handlers called",
	})
}

func (h *Handler) ChangePasswordHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "mocked handlers called",
	})
}
