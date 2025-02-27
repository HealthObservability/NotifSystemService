package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
)

func (p *Postgres) GetNotifications(context.Context) ([]domains.NotifStatus, error) {
	log := logger.Logger.WithField("op", "Postgres.GetNotifStates")
	log.Debug("Getting all notifications MOCK")

	return []domains.NotifStatus{}, nil
}
