package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"time"
)

func (p *Postgres) GetNotifications(context.Context) ([]domains.NotifStatus, error) {
	log := logger.Logger.WithField("op", "Postgres.GetNotifStates")
	log.Debug("Getting all notifications MOCK")

	q := fmt.Sprintf(
		"SELECT * FROM notif_states WHERE send_due <= %s AND status = 'not_sent'",
		time.Now().Format(time.RFC3339),
	)

	rows, err := p.db.Query(q)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.WithError(err).Error("Error closing rows")
		}
	}(rows)

	var notifs []domains.NotifStatus
	for rows.Next() {
		var notif domains.NotifStatus
		err := rows.Scan(
			&notif.ID,
			&notif.NotificationID,
			&notif.SendDue,
			&notif.CreateDate,
			&notif.Status,
		)
		if err != nil {
			return nil, err
		}
		notifs = append(notifs, notif)
	}

	return notifs, nil
}
