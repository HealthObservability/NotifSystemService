package services

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/services/notifservice"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

type Notif interface {
	GetNotifStates(context.Context) ([]domains.NotifState, error)
	ChangeNotifState(context.Context, uint64, string) error
	GetFullNotifs(context.Context, []domains.NotifState) ([]domains.Notification, error)
	UpdateNotifications(ctx context.Context, id, msg string) error
}

type Sender interface {
	SendNotifications(ctx context.Context, notifs []domains.Notification) error
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
