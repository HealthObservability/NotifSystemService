package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteNotification(c *gin.Context) {
	log := logger.GetLogger().WithField("op", "Handler.DeleteNotification")
	log.Info("received delete notif request")

	notifID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.WithError(err).Error("error parsing id param")
		c.JSON(http.StatusBadRequest, gin.H{"error": " is required"})
		return
	}

	notifs, err := h.service.NotifService.GetNotifs(c, []domains.NotifState{{NotifID: notifID}})
	if err != nil {
		log.WithError(err).Error("error getting notif from db")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if len(notifs) == 0 {
		log.WithError(err).Warn("notifs are empty")
		c.JSON(http.StatusNotFound, gin.H{"error": "notifs not found"})
		return
	}

	err = h.service.NotifService.DeleteNotif(c, notifID)
	if err != nil {
		log.WithError(err).Error("error deleting notif")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete notif"})
		return
	}

	c.Status(http.StatusNoContent)
}
