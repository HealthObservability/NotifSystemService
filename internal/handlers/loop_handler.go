package handlers

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
)

func (h *Handler) handleLoop(ctx context.Context) {
	log := logger.Logger.WithField("op", "handler.handleLoop")
	log.Info("Loop iteration")

	notifications, err := h.service.GetNotifications(ctx)
	if err != nil {
		log.WithError(err).Error("Error getting notifications")
		return
	}

	if err := h.service.SendNotifications(ctx, notifications); err != nil {
		log.WithError(err).Error("Error sending notifications")
		return
	}
}
