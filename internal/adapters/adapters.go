package adapters

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters/httpserver"
	"github.com/HealthObservability/NotifSystemService/internal/adapters/looper"
	"github.com/HealthObservability/NotifSystemService/internal/adapters/postgres"
	"github.com/HealthObservability/NotifSystemService/internal/config"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

type Notifications interface {
	GetStatesByStatus(ctx context.Context, status string) ([]domains.NotifState, error)
	GetNotifs(ctx context.Context, ids []int64) ([]domains.Notif, error)
	InsertNotif(ctx context.Context, notif domains.Notif) (int64, error)
	InsertState(ctx context.Context, state domains.NotifState) (int64, error)
	InsertScheduled(ctx context.Context, state domains.NotifState) error
	UpdateStatuses(ctx context.Context, status string, ids []int64) error
	DeleteNotif(ctx context.Context, id int64) error
}

type Database interface {
	Shutdown(ctx context.Context) error
	Notifications
}

type Adapters struct {
	Database Database
	Looper   *looper.Looper
	HTTPServ *httpserver.HTTPServAdapter
}

func NewAdapters(cfg config.Config) *Adapters {
	return &Adapters{
		Database: postgres.MustNew(cfg.Postgres),
		HTTPServ: httpserver.New(cfg.HTTPServer),
		Looper:   looper.NewLooper(),
	}
}
