package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	var notifs []domains.Notif
	if err == nil {
		var ids []domains.NotifState
		ids = append(ids, domains.NotifState{NotifID: id})
		notifs, err = h.service.NotifService.GetNotifs(c, ids)
		if err != nil {
			log.WithError(err).Error("Error getting notifs")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	} else {
		notifs, err = h.service.NotifService.GetAllNotifs(c)
		if err != nil {
			log.WithError(err).Error("Error getting notifs")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	if len(notifs) == 0 {
		log.WithError(err).Warn("notifs are empty")
		c.JSON(http.StatusNotFound, gin.H{"error": "notifs not found"})
		return
	}

	c.JSON(http.StatusOK, notifs)
}

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

	id, err := h.service.NotifService.AddNotification(c, newNotif)
	if err != nil {
		log.WithError(err).Error("Cannot add notification")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatusJSON(http.StatusCreated, gin.H{"id": id})
}

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
