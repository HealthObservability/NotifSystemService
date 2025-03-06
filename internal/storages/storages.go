package storages

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

type NotificationStates interface {
	SetNewStatuses(ctx context.Context, status string, ids []int64) error
	GetNotifStates(ctx context.Context, status string) ([]domains.NotifState, error)
}

type Notifications interface {
	GetNotifs(ctx context.Context, ids []int64) ([]domains.Notif, error)
	InsertNotification(ctx context.Context, n domains.Notif) (int64, error)
	InsertNotifState(ctx context.Context, n domains.NotifState) (int64, error)
	DeleteNotif(ctx context.Context, notifID int64) error
}

type Storage struct {
	NotificationStates
	Notifications
}

func MustNew(a *adapters.Adapters) *Storage {
	return &Storage{
		NotificationStates: a.Postgres,
		Notifications:      a.Postgres,
	}
}
