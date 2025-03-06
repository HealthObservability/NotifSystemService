package postgres

import (
	"context"
	"database/sql"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"github.com/lib/pq"
)

func (p *Postgres) GetNotifs(ctx context.Context, ids []int64) ([]domains.Notif, error) {
	log := logger.GetLogger().WithField("op", "GetNotifs")

	q := `
			SELECT id, creation_timestamp, id_to_send, text, repeat_interval, last_notif, since_time
			FROM notif_info
	`

	if len(ids) > 0 {
		q += `WHERE id = ANY($1)`
	}

	var rows *sql.Rows
	var err error
	if len(ids) > 0 {
		rows, err = p.db.QueryContext(ctx, q, pq.Array(ids))
	} else {
		rows, err = p.db.QueryContext(ctx, q)
	}
	if err != nil {
		return nil, err
	} else if rows != nil && rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func() {
		if closeErr := rows.Close(); err != nil {
			log.WithError(closeErr).Error("failed to close rows")
		}
	}()

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
