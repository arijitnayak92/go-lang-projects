package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
HealthHandler returns alive status
*/
func (h *Handler) HealthHandler(c *gin.Context) {
	postgresDBStatus := h.domain.GetPostgresHealth()
	mongoDBStatus := h.domain.GetMongoHealth()
	if postgresDBStatus && mongoDBStatus {
		c.JSON(http.StatusOK, gin.H{
			"alive":     true,
			"db_status": "connected",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"alive":         true,
		"postgresAlive": postgresDBStatus,
	})
}
