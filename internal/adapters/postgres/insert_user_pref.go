package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"log"
)

func (p *Postgres) InsertUserPreferences(ctx context.Context, prefs domains.UserPreferences) (int64, error) {
	q := `
		INSERT INTO user_preferences (user_id, dont_disturb_start, dont_disturb_end, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	rows, err := p.db.QueryContext(ctx, q,
		prefs.UserID, prefs.DontDisturbStart, prefs.DontDisturbEnd, prefs.CreatedAt, prefs.UpdatedAt)
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
