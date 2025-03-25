package handlers

import (
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	prefs, err := h.service.UserPref.GetPreferences(c, filter)
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

	err = h.service.UserPref.InsertPreference(c, userPref)
	if err != nil {
		log.WithError(err).Error("error inserting user pref")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert user pref"})
		return
	}

	c.Status(http.StatusNoContent)
}

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

	prefs, err := h.service.UserPref.GetPreferences(c, domains.UserPrefs{ID: &id})
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

	err = h.service.UserPref.DeletePreference(c, id)
	if err != nil {
		log.WithError(err).Error("error deleting user pref")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user pref"})
		return
	}

	c.Status(http.StatusNoContent)
}
