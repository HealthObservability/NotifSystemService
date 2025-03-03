package postgres

import (
	"context"
	"database/sql"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"time"
)

func (p *Postgres) GetNotifStates(ctx context.Context) ([]domains.NotifState, error) {
	log := logger.Logger.WithField("op", "Postgres.getNotifStates")

	q := `
			SELECT * 
			FROM notif_states
			WHERE send_due <= $1 AND status = 'not_sent'
	`

	rows, err := p.db.QueryContext(ctx, q, time.Now())
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

	var notifs []domains.NotifState
	for rows.Next() {
		var notif domains.NotifState
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
