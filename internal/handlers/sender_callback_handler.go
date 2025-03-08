package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HandleHTTPCallback(c *gin.Context) {
	log := logger.GetLogger().WithField("op", "Handler.HandleHTTPCallback")

	var cb domains.SenderCallback
	if err := c.ShouldBindJSON(&cb); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CallbackService.ProcessCallback(c, cb); err != nil {
		log.WithError(err).Error("Cannot process callback")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
