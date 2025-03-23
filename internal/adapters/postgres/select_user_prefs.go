package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (p *Postgres) UserPrefs(ctx context.Context, f domains.UserPrefs) ([]domains.UserPrefs, error) {
	// log := logger.GetLogger().WithField("op", "Postgres.UserPrefs")

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
