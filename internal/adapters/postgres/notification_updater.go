package postgres

import (
	"context"
	"github.com/lib/pq"
)

func (p *Postgres) SetNewStatuses(ctx context.Context, status string, ids []uint64) error {
	q := `
		UPDATE notif_states
		SET status = $1
		WHERE id = ANY($2)
	`

	_, err := p.db.ExecContext(ctx, q, status, pq.Array(ids))
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateNotifications(context.Context) error {
	// TODO implement me
	panic("implement me")
}
