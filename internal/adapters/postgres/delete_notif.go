package postgres

import "context"

func (p *Postgres) DeleteNotif(ctx context.Context, notifID int64) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	q := `DELETE FROM notif_states WHERE notif_id = $1`

	_, err = tx.ExecContext(ctx, q, notifID)
	if err != nil {
		return tx.Rollback()
	}

	q = `DELETE FROM notif_info WHERE id = $1`
	_, err = tx.Exec(q, notifID)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
