package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteUserPref(c *gin.Context) {
	log := logger.GetLogger().WithField("op", "Handler.DeleteUserPref")
	log.Info("received delete user pref request")

	if _, ok := c.Params.Get("id"); !ok {
		log.Error("user pref id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.WithError(err).Error("failed to parse id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse id. it should be a number"})
		return
	}

	prefs, err := h.service.UserPreferences.GetPreferences(c, domains.UserPreferences{ID: id})
	if err != nil {
		log.WithError(err).Error("error reading user pref")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user pref"})
		return
	}

	if len(prefs) == 0 {
		log.Warn("user pref not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "user pref not found"})
		return
	}

	err = h.service.UserPreferences.DeletePreference(c, id)
	if err != nil {
		log.WithError(err).Error("error deleting user pref")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user pref"})
		return
	}

	c.Status(http.StatusNoContent)
}
