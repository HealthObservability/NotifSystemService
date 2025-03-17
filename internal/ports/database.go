package ports

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

type Database interface {
	GracefulShutdown
	Notifications
	NotificationStates
}

type Notifications interface {
	GetNotifs(ctx context.Context, ids []int64) ([]domains.Notif, error)
	InsertNotif(ctx context.Context, notif domains.Notif) (int64, error)
	InsertState(ctx context.Context, state domains.NotifState) (int64, error)
	InsertScheduled(ctx context.Context, state domains.NotifState) error
	DeleteNotif(ctx context.Context, id int64) error
}

type NotificationStates interface {
	UpdateStatuses(ctx context.Context, status string, ids []int64) error
	GetStatesByStatus(ctx context.Context, status string) ([]domains.NotifState, error)
}
