package storages

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

type NotificationStates interface {
	UpdateNotifications(ctx context.Context) error
	GetNotifications(context.Context) ([]domains.NotifStatus, error)
}

type Notifications interface {
	AddNotification()
	AddNotifications()
}

type Storage struct {
	NotificationStates
	Notifications
}

func MustNew(a *adapters.Adapters) *Storage {
	return &Storage{
		NotificationStates: a.Postgres,
	}
}
