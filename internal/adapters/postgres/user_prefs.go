package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (p *Postgres) UserPrefs(ctx context.Context, f domains.UserPrefs) ([]domains.UserPrefs, error) {
	q := `
		SELECT id, user_id, dont_disturb_start, dont_disturb_end, created_at, updated_at
		FROM user_preferences
	`

	rows, err := p.pool.Query(ctx, q, f.UserID)
	if err != nil {
		return []domains.UserPrefs{}, err
	}
	defer rows.Close()

	var userPrefs []domains.UserPrefs
	for rows != nil && rows.Next() {
		var userPref domains.UserPrefs
		err = rows.Scan(
			&userPref.UserID,
			&userPref.DontDisturbStart,
			&userPref.DontDisturbEnd,
			&userPref.CreatedAt,
			&userPref.UpdatedAt,
		)
		if err != nil {
			return []domains.UserPrefs{}, err
		}
	}

	return userPrefs, nil
}

func (p *Postgres) InsertUserPreferences(ctx context.Context, prefs domains.UserPrefs) (int64, error) {
	q := `
		INSERT INTO user_preferences (user_id, dont_disturb_start, dont_disturb_end, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	rows, err := p.pool.Query(ctx, q,
		prefs.UserID, prefs.DontDisturbStart, prefs.DontDisturbEnd, prefs.CreatedAt, prefs.UpdatedAt)
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

func (p *Postgres) DeleteUserPrefs(ctx context.Context, prefsID int64) error {
	q := `DELETE FROM user_preferences WHERE id = $1`

	_, err := p.pool.Exec(ctx, q, prefsID)
	if err != nil {
		return err
	}

	return nil
}
