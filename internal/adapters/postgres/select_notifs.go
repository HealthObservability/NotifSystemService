package postgres

import (
	"context"
	"database/sql"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
)

func (p *Postgres) GetNotifs(ctx context.Context, id uint64) ([]any, error) {
	log := logger.Logger.WithField("op", "GetNotifs")

	q := `
			SELECT *
			FROM notifs
			WHERE id = $1
	`

	rows, err := p.db.QueryContext(ctx, q, id)
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

	var notifs []any
	for rows.Next() {
		var notif any
		if err := rows.Scan(&notif); err != nil {
			return nil, err
		}
	}

	return notifs, nil
}
