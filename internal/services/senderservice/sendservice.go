package senderservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"golang.org/x/sync/errgroup"
	"time"
)

type SenderService interface {
	SendNotifications(ctx context.Context, notifs []domains.Notif) error
	SendNotification(ctx context.Context, notif domains.Notif) error
}

type Service struct {
	adapters *adapters.Adapters
}

func New(a *adapters.Adapters) *Service {
	return &Service{adapters: a}
}

func (s Service) SendNotification(ctx context.Context, notif domains.Notif) error {
	//TODO implement me
	return nil
}

func (s Service) SendNotifications(ctx context.Context, notifs []domains.Notif) error {
	// log := logger.GetLogger().WithField("op", "Service.SendNotifications")

	// TODO implement me

	g, gCtx := errgroup.WithContext(ctx)

	// FIXME SEND each notif to sender service
	for _, notif := range notifs {
		notifCopy := notif
		g.Go(func() error {
			err := s.SendNotification(gCtx, notifCopy)
			if err != nil {
				return err
			}

			// FIXME it is a mock
			time.Sleep(time.Second)
			// update status
			// FIXME
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
