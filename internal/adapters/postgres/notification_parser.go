package postgres

import "github.com/HealthObservability/NotifSystemService/pkg/logger"

func (p *Postgres) GetNotifications() {
	log := logger.Logger.WithField("op", "Postgres.GetNotifications")
	log.Debug("Getting all notifications MOCK")
}
