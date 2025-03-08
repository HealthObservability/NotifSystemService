package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (p *Postgres) ScheduleNotification(ctx context.Context, n domains.NotifState) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := `
		UPDATE notif_info
		SET last_scheduled_time = $1
		WHERE id = $2
	`

	_, err = tx.ExecContext(ctx, q, n.SendDue, n.NotifID)
	if err != nil {
		return tx.Rollback()
	}

	q = `
		INSERT INTO notif_states (notif_id, send_due, creation_timestamp, status)
		VALUES ($1, $2, $3, $4)
	`

	_, err = tx.ExecContext(ctx, q, n.NotifID, n.SendDue, n.CreateDate, n.Status)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
