package notifservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/ports"
)

const (
	NotSentStatus = "not_sent"
	SentStatus    = "sent"
	PendingStatus = "pending"
)

type NotifService interface {
	GetNotifs(ctx context.Context, n []domains.NotifState) ([]domains.Notif, error)
	GetAllNotifs(ctx context.Context) ([]domains.Notif, error)
	GetNotifStates(ctx context.Context) ([]domains.NotifState, error)
	UpdateSendStatus(ctx context.Context, n []domains.NotifState, newStatus string) error
	AddNotification(ctx context.Context, n domains.Notif) (int64, error)
	ScheduleNotifications(ctx context.Context, n []domains.Notif) error
	DeleteNotif(ctx context.Context, notifID int64) error
	//UpdateNotification(ctx context.Context, n domains.Notif) error
	//DeleteNotification(ctx context.Context, id uint64) error
}

type service struct {
	Db ports.Database
}

func NewNotifService(db ports.Database) NotifService {
	return &service{db}
}
