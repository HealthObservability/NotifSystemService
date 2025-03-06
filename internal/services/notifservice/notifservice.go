package notifservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

const (
	NotSentStatus = "not_sent"
	SentStatus    = "sent"
	PendingStatus = "pending"
)

type NotifService interface {
	GetNotifs(ctx context.Context, n []domains.NotifState) ([]domains.Notif, error)
	GetNotifStates(ctx context.Context) ([]domains.NotifState, error)
	ChangeStatesStatus(ctx context.Context, n []domains.NotifState, newStatus string) error
	AddNotification(ctx context.Context, n domains.Notif) (int64, error)
	DeleteNotif(ctx context.Context, notifID int64) error
	//UpdateNotification(ctx context.Context, n domains.Notif) error
	//DeleteNotification(ctx context.Context, id uint64) error
}

type Service struct {
	storage *storages.Storage
}

func NewNotifService(s *storages.Storage) *Service {
	return &Service{s}
}
