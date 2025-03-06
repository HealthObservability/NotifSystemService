package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetNotifs(c *gin.Context) {
	log := logger.GetLogger().WithField("op", "GetNotifs")
	log.Info("Received get notifs req")

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil && c.Param("id") != "" {
		log.WithError(err).Error("Invalid id param")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		log.WithError(err).Warn("Invalid id param")
	}

	var ids []domains.NotifState
	if err == nil {
		ids = append(ids, domains.NotifState{NotifID: id})
	}

	notifs, err := h.service.NotifService.GetNotifs(c, ids)
	if err != nil {
		log.WithError(err).Error("Error getting notifs")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	
	if len(notifs) == 0 {
		log.WithError(err).Warn("notifs are empty")
		c.JSON(http.StatusNotFound, gin.H{"error": "notifs not found"})
		return
	}

	c.JSON(http.StatusOK, notifs)
}
