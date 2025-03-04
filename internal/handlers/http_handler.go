package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *Handler) AddNotification(c *gin.Context) {
	log := logger.GetLogger().WithField("op", "Handler.AddNotification")
	log.Info("Received add notification req")

	var newNotif domains.Notif
	if err := c.ShouldBindJSON(&newNotif); err != nil {
		log.WithError(err).Error("Cannot bind json body")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := validator.New()
	if err := v.Struct(newNotif); err != nil {
		log.WithError(err).Error("Cannot bind json body")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Notif.AddNotification(c, newNotif)
	if err != nil {
		log.WithError(err).Error("Cannot add notification")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatusJSON(http.StatusCreated, gin.H{"id": id})
}
