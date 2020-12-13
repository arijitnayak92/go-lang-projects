package mock

// import (
// 	"net/http"
//
// 	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
// 	"github.com/arijitnayak92/taskAfford/Fruit/domain"
// 	"github.com/gin-gonic/gin"
// )
//
// type Handler struct {
// 	appContext *appcontext.AppContext
// 	domain     domain.AppDomain
// }
//
// func NewHandler(appContext *appcontext.AppContext, domain domain.AppDomain) *Handler {
// 	return &Handler{appContext: appContext, domain: domain}
// }
//
// func (h *Handler) HealthHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "mocked handlers called",
// 	})
// }
