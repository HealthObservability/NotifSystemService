package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"log"
)

func (p *Postgres) InsertNotification(ctx context.Context, n domains.Notif) (int64, error) {
	q := `
		INSERT INTO notif_info
		(creation_timestamp, id_to_send, text, repeat_interval, last_notif, since_time)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	rows, err := p.db.QueryContext(
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
	} else if rows != nil && rows.Err() != nil {
		return 0, rows.Err()
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows: ", err)
		}
	}()

	var lastID int64
	for rows.Next() {
		if err := rows.Scan(&lastID); err != nil {
			return 0, err
		}
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

	rows, err := p.db.QueryContext(
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
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows: ", err)
		}
	}()

	var lastID int64
	for rows.Next() {
		if err := rows.Scan(&lastID); err != nil {
			return 0, err
		}
	}

	return lastID, nil
}
