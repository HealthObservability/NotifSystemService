package adapters

import (
	"github.com/HealthObservability/NotifSystemService/internal/adapters/httpserver"
	"github.com/HealthObservability/NotifSystemService/internal/adapters/looper"
	"github.com/HealthObservability/NotifSystemService/internal/adapters/postgres"
	"github.com/HealthObservability/NotifSystemService/internal/config"
	"github.com/HealthObservability/NotifSystemService/internal/ports"
)

type Adapters struct {
	Database ports.Database
	Looper   ports.Looper
	HTTPServ ports.HTTPServer
}

func NewAdapters(cfg config.Config) *Adapters {
	return &Adapters{
		Database: postgres.MustNew(cfg.Postgres),
		HTTPServ: httpserver.New(cfg.HTTPServer),
		Looper:   looper.NewLooper(),
	}
}
