package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

func (p *Postgres) GetNotifs(ctx context.Context, ids []int64) ([]domains.Notif, error) {
	// log := logger.GetLogger().WithField("op", "GetNotifs")

	q := `
			SELECT id, creation_timestamp, id_to_send, text, repeat_interval, last_scheduled_time, since_time
			FROM notif_info
	`

	if len(ids) > 0 {
		q += `WHERE id = ANY($1)`
	}

	var rows pgx.Rows
	var err error
	if len(ids) > 0 {
		rows, err = p.pool.Query(ctx, q, pq.Array(ids))
	} else {
		rows, err = p.pool.Query(ctx, q)
	}
	if err != nil {
		return nil, err
	} else if rows != nil && rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var notifs []domains.Notif
	for rows != nil && rows.Next() {
		var notif domains.Notif
		err = rows.Scan(
			&notif.ID,
			&notif.CreatedAt,
			&notif.ToID,
			&notif.Message,
			&notif.RepeatInterval,
			&notif.LastScheduled,
			&notif.Since,
		)
		if err != nil {
			return nil, err
		}

		notifs = append(notifs, notif)
	}

	return notifs, nil
}
