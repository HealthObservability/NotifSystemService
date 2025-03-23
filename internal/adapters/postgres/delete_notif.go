package postgres

import "context"

func (p *Postgres) DeleteNotif(ctx context.Context, notifID int64) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}

	q := `DELETE FROM notif_states WHERE notif_id = $1`

	_, err = tx.Exec(ctx, q, notifID)
	if err != nil {
		return tx.Rollback(ctx)
	}

	q = `DELETE FROM notif_info WHERE id = $1`
	_, err = tx.Exec(ctx, q, notifID)
	if err != nil {
		return tx.Rollback(ctx)
	}

	return tx.Commit(ctx)
}
