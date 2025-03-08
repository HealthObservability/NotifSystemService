package handlers

import (
	"context"
	"fmt"
	"github.com/HealthObservability/NotifSystemService/internal/services/notifservice"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"golang.org/x/sync/errgroup"
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

	if err := h.service.UpdateSendStatus(ctx, notifStates, notifservice.PendingStatus); err != nil {
		log.WithError(err).Error("Error changing states")
		return
	}

	fullNotifs, err := h.service.GetNotifs(ctx, notifStates)
	if err != nil {
		log.WithError(err).Error("Error getting notifs")
	}
	log.WithField("full_notifs", fullNotifs).Debug("Got full notifs")

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := h.service.NotifService.ScheduleNotifications(ctx, fullNotifs); err != nil {
			return fmt.Errorf("failed to schedule notifications: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		if err := h.service.SenderService.SendNotifications(ctx, fullNotifs); err != nil {
			return fmt.Errorf("failed to send notifications: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.WithError(err).Error("Error in err group during loop")
		return
	}
}
