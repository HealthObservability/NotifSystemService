package postgres

import (
	"context"
	"database/sql"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/lib/pq"
)

func (p *Postgres) GetNotifs(ctx context.Context, ids []uint64) ([]domains.Notif, error) {
	log := logger.Logger.WithField("op", "GetNotifs")

	q := `
			SELECT id, creation_timestamp, id_to_send, text, repeat_interval, last_notif, since_time
			FROM notif_info
			WHERE id = ANY($1)
	`

	rows, err := p.db.QueryContext(ctx, q, pq.Array(ids))
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.WithError(err).Error("failed to close rows")
		}
	}(rows)

	var notifs []domains.Notif
	for rows.Next() {
		var notif domains.Notif
		err = rows.Scan(
			&notif.ID,
			&notif.CreatedAt,
			&notif.ToID,
			&notif.Message,
			&notif.RepeatInterval,
			&notif.LastSent,
			&notif.Since,
		)
		if err != nil {
			return nil, err
		}

		notifs = append(notifs, notif)
	}

	return notifs, nil
}
