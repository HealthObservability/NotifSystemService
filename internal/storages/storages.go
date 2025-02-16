package storages

import (
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
)

type NotificationStates interface {
	UpdateNotifications()
	GetNotifications()
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
