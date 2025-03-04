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
	AddNotification(ctx context.Context, n domains.Notif) (int64, error)
	//UpdateNotification(ctx context.Context, n domains.Notif) error
	//DeleteNotification(ctx context.Context, id uint64) error
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
