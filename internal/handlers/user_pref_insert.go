package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) InsertUserPref(c *gin.Context) {
	log := logger.GetLogger().WithField("op", "Handler.InsertUserPref")
	log.Info("received insert user pref request")

	var userPref domains.UserPrefs
	err := c.BindJSON(&userPref)
	if err != nil {
		log.WithError(err).Error("error binding json")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err = h.service.UserPreferences.InsertPreference(c, userPref)
	if err != nil {
		log.WithError(err).Error("error inserting user pref")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert user pref"})
		return
	}

	c.Status(http.StatusNoContent)
}
