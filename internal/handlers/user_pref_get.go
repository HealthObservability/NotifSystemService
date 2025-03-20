package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUserPrefs(c *gin.Context) {
	log := logger.GetLogger().WithField("op", "Handler.GetUserPrefs")
	log.Info("received get user prefs request")

	var filter domains.UserPrefs
	if err := c.ShouldBindQuery(&filter); err != nil {
		log.WithError(err).Error("error binding user prefs")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	log.WithField("filter", filter).Info("query filters")

	prefs, err := h.service.UserPreferences.GetPreferences(c, filter)
	if err != nil {
		log.WithError(err).Error("error getting user prefs")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user prefs"})
		return
	}

	if len(prefs) == 0 {
		log.Warn("user prefs are empty")
		c.JSON(http.StatusNotFound, gin.H{"error": "user prefs not found"})
		return
	}

	c.JSON(http.StatusOK, prefs)
}
