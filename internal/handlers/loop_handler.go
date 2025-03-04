package handlers

import (
	"context"

	"github.com/HealthObservability/NotifSystemService/pkg/logger"
)

func (h *Handler) handleLoop(ctx context.Context) {
	log := logger.GetLogger().WithField("op", "handler.handleLoop")
	log.Info("Loop iteration")

	notifStates, err := h.service.GetNotifStates(ctx)
	if err != nil {
		log.WithError(err).Error("Error getting notifStates")
		return
	}
	log.WithField("states", notifStates).Debug("Got notifStates")

	// if err := h.service.ChangeStatesStatus(ctx, notifStates, notifservice.PendingStatus); err != nil {
	// 	log.WithError(err).Error("Error changing states")
	// 	return
	// }

	fullNotifs, err := h.service.GetNotifs(ctx, notifStates)
	if err != nil {
		log.WithError(err).Error("Error getting notifs")
	}

	log.WithField("full_notifs", fullNotifs).Debug("Got full notifs")
}
