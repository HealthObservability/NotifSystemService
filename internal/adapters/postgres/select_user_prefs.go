package postgres

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
)

func (p *Postgres) UserPrefs(ctx context.Context, f domains.UserPrefs) ([]domains.UserPrefs, error) {
	log := logger.GetLogger().WithField("op", "Postgres.UserPrefs")

	q := `
		SELECT id, user_id, dont_disturb_start, dont_disturb_end, created_at, updated_at
		FROM user_preferences
	`

	rows, err := p.db.QueryContext(ctx, q, f.UserID)
	if err != nil {
		return []domains.UserPrefs{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error(err.Error())
		}
	}()

	var userPrefs []domains.UserPrefs
	for rows.Next() {
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
