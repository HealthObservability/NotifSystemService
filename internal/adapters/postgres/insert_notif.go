package postgres

import (
	"context"
	"fmt"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (p *Postgres) InsertNotif(ctx context.Context, n domains.Notif) (int64, error) {
	q := `
		INSERT INTO notif_info
		(creation_timestamp, id_to_send, text, repeat_interval, last_scheduled_time, since_time)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	fmt.Println("n.Since:", n.Since)
	fmt.Println("n.LastScheduled:", n.LastScheduled)

	rows, err := p.pool.Query(
		ctx,
		q,
		n.CreatedAt,
		n.ToID,
		n.Message,
		n.RepeatInterval,
		n.LastScheduled,
		n.Since,
	)
	if err != nil {
		return 0, err
	} else if rows != nil && rows.Err() != nil {
		return 0, rows.Err()
	}
	defer rows.Close()

	var lastID int64
	for rows != nil && rows.Next() {
		if err := rows.Scan(&lastID); err != nil {
			return 0, err
		}
	}

	return lastID, nil
}

func (p *Postgres) InsertState(ctx context.Context, n domains.NotifState) (int64, error) {
	q := `
		INSERT INTO notif_states 
		    (notif_id, send_due, creation_timestamp, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	rows, err := p.pool.Query(
		ctx,
		q,
		n.NotifID,
		n.SendDue,
		n.CreateDate,
		n.Status,
	)
	if err != nil {
		return 0, err
	} else if rows != nil && rows.Err() != nil {
		return 0, rows.Err()
	}
	defer rows.Close()

	var lastID int64
	for rows != nil && rows.Next() {
		if err := rows.Scan(&lastID); err != nil {
			return 0, err
		}
	}

	return lastID, nil
}
