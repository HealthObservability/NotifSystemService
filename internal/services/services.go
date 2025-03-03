package services

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/services/notifservice"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

type Notif interface {
	GetNotifs(ctx context.Context, n []domains.NotifState) ([]domains.Notif, error)
	GetNotifStates(ctx context.Context) ([]domains.NotifState, error)
	ChangeStatesStatus(ctx context.Context, n []domains.NotifState, newStatus string) error
	UpdateNotifications(ctx context.Context, id, msg string) error
}

type Sender interface {
	SendNotifications(ctx context.Context, notifs []domains.Notif) error
}

type Service struct {
	Notif
	Sender
}

func MustNew(s *storages.Storage) *Service {
	return &Service{
		Notif: notifservice.NewNotifService(s),
	}
}
