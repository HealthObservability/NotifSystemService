package services

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/services/checker"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

type Notif interface {
	GetNotifications(context.Context) ([]any, error)
	UpdateNotifications(ctx context.Context, id, msg string) error
}

type Sender interface {
	SendNotifications(ctx context.Context, notifs []any) error
}

type Service struct {
	Notif
	Sender
}

func MustNew(s *storages.Storage) *Service {
	return &Service{
		Notif: checker.NewNotifService(s),
	}
}
