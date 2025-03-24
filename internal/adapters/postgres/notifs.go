package postgres

import (
	"context"
	"fmt"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
	"time"
)

func (p *Postgres) GetNotifs(ctx context.Context, ids []int64) ([]domains.Notif, error) {
	q := `
			SELECT id, creation_timestamp, id_to_send, text,
			       repeat_interval, last_scheduled_time, since_time 
			FROM notif_info
	`

	if len(ids) > 0 {
		q += `WHERE id = ANY($1)`
	}

	var rows pgx.Rows
	var err error
	if len(ids) > 0 {
		rows, err = p.pool.Query(ctx, q, pq.Array(ids))
	} else {
		rows, err = p.pool.Query(ctx, q)
	}
	if err != nil {
		return nil, err
	} else if rows != nil && rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var notifs []domains.Notif
	for rows != nil && rows.Next() {
		var notif domains.Notif
		err = rows.Scan(
			&notif.ID,
			&notif.CreatedAt,
			&notif.ToID,
			&notif.Message,
			&notif.RepeatInterval,
			&notif.LastScheduled,
			&notif.Since,
		)
		if err != nil {
			return nil, err
		}

		notifs = append(notifs, notif)
	}

	return notifs, nil
}

func (p *Postgres) StatesByStatus(ctx context.Context, status string) ([]domains.NotifState, error) {
	// log := logger.GetLogger().WithField("op", "Postgres.StatesByStatus")

	q := `
			SELECT * 
			FROM notif_states
			WHERE send_due <= $1 AND status = $2
	`

	rows, err := p.pool.Query(ctx, q, time.Now().UTC(), status)
	if err != nil {
		return nil, err
	} else if rows != nil && rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var notifs []domains.NotifState
	for rows != nil && rows.Next() {
		var notif domains.NotifState
		err := rows.Scan(
			&notif.ID,
			&notif.NotifID,
			&notif.SendDue,
			&notif.CreateDate,
			&notif.Status,
		)
		if err != nil {
			return nil, err
		}
		notifs = append(notifs, notif)
	}

	return notifs, nil
}

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
