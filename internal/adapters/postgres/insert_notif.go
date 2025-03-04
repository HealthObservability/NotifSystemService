package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (p *Postgres) InsertNotification(ctx context.Context, n domains.Notif) (int64, error) {
	q := `
		INSERT INTO notif_info
		(creation_timestamp, id_to_send, text, repeat_interval, last_notif, since_time)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	res, err := p.db.ExecContext(
		ctx,
		q,
		n.CreatedAt,
		n.ToID,
		n.Message,
		n.RepeatInterval,
		n.LastSent,
		n.Since,
	)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (p *Postgres) InsertNotifState(ctx context.Context, n domains.NotifState) (int64, error) {
	q := `
		INSERT INTO notif_states 
		    (notif_id, send_due, creation_timestamp, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	res, err := p.db.ExecContext(
		ctx,
		q,
		n.NotifID,
		n.SendDue,
		n.CreateDate,
		n.Status,
	)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}
