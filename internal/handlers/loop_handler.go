package handlers

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
)

func (h *Handler) handleLoop(ctx context.Context) {
	log := logger.Logger.WithField("op", "handler.handleLoop")
	log.Info("Loop iteration")

	notifStates, err := h.service.GetNotifStates(ctx)
	if err != nil {
		log.WithError(err).Error("Error getting notifStates")
		return
	}

	fullNotifications, err := h.service.GetFullNotifs(ctx, notifStates)
	if err != nil {
		log.WithError(err).Error("Error getting fullNotifications")
		return
	}

	if err := h.service.SendNotifications(ctx, fullNotifications); err != nil {
		log.WithError(err).Error("Error sending notifStates")
		return
	}
}
