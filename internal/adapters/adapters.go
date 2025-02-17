package adapters

import (
	"github.com/HealthObservability/NotifSystemService/internal/adapters/looper"
	"github.com/HealthObservability/NotifSystemService/internal/adapters/postgres"
	"github.com/HealthObservability/NotifSystemService/internal/config"
)

type Adapters struct {
	Postgres *postgres.Postgres
	Looper   *looper.Looper
}

func NewAdapters(cfg config.Config) *Adapters {
	return &Adapters{
		Postgres: postgres.MustNew(cfg.Postgres),
		Looper:   looper.NewLooper(),
	}
}
