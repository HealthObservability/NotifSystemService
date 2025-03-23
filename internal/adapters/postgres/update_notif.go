package postgres

import (
	"context"
	"github.com/lib/pq"
)

func (p *Postgres) UpdateStatuses(ctx context.Context, status string, ids []int64) error {
	q := `
		UPDATE notif_states
		SET status = $1
		WHERE id = ANY($2)
	`

	_, err := p.pool.Exec(ctx, q, status, pq.Array(ids))
	if err != nil {
		return err
	}

	return nil
}
