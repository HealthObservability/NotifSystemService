package storages

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

type NotificationStates interface {
	UpdateNotifications(ctx context.Context) error
	SetNewStatuses(ctx context.Context, status string, ids []uint64) error
	GetNotifStates(context.Context) ([]domains.NotifState, error)
}

type Notifications interface {
	GetNotifs(ctx context.Context, ids []uint64) ([]domains.Notif, error)
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
