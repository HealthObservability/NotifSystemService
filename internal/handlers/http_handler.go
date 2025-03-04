package handlers

import (
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddNotification(c *gin.Context) {
	log := logger.GetLogger().WithField("op", "Handler.AddNotification")
	log.Info("Received add notification req")

	log.Info(c.Err())
}
