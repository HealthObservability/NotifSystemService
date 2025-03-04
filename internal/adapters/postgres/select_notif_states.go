package postgres

import (
	"context"
	"time"

	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
)

func (p *Postgres) GetNotifStates(ctx context.Context, status string) ([]domains.NotifState, error) {
	log := logger.GetLogger().WithField("op", "Postgres.getNotifStates")

	q := `
			SELECT * 
			FROM notif_states
			WHERE send_due <= $1 AND status = $2
	`

	rows, err := p.db.QueryContext(ctx, q, time.Now(), status)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.WithError(err).Error("Error closing rows")
		}
	}()

	var notifs []domains.NotifState
	for rows.Next() {
		var notif domains.NotifState
		err := rows.Scan(
			&notif.ID,
			&notif.NotifID,
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
