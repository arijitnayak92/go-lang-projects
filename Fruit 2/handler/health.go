package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//...
func (h *Handler) HealthHandler(c *gin.Context) {
	postgresDBStatus, mongoDBStatus := h.domain.CheckDatabaseHealth()
	if mongoDBStatus != nil && postgresDBStatus != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"mongoIsAlive":    false,
			"postgresIsAlive": false,
			"serverIsAlive":   true,
		})
		return
	}
	if mongoDBStatus == nil && postgresDBStatus != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"mongoIsAlive":    true,
			"postgresIsAlive": false,
			"serverIsAlive":   true,
		})
		return
	}

	if mongoDBStatus != nil && postgresDBStatus == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"mongoIsAlive":    false,
			"postgresIsAlive": true,
			"serverIsAlive":   true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mongoIsAlive":    true,
		"postgresIsAlive": true,
		"serverIsAlive":   true,
	})
}
