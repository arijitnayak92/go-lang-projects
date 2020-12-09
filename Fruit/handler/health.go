package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
HealthHandler returns alive status
*/
func (h *Handler) HealthHandler(c *gin.Context) {
	postgresDBStatus := h.domain.GetPostgresHealth()
	if postgresDBStatus {
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
