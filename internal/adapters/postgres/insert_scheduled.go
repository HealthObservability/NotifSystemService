package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (p *Postgres) InsertScheduled(ctx context.Context, n domains.NotifState) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}

	q := `
		UPDATE notif_info
		SET last_scheduled_time = $1
		WHERE id = $2
	`

	_, err = tx.Exec(ctx, q, n.SendDue, n.NotifID)
	if err != nil {
		return tx.Rollback(ctx)
	}

	q = `
		INSERT INTO notif_states (notif_id, send_due, creation_timestamp, status)
		VALUES ($1, $2, $3, $4)
	`

	_, err = tx.Exec(ctx, q, n.NotifID, n.SendDue, n.CreateDate, n.Status)
	if err != nil {
		return tx.Rollback(ctx)
	}

	return tx.Commit(ctx)
}
