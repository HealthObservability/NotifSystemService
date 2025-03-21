package postgres

import "context"

func (p *Postgres) DeleteUserPrefs(ctx context.Context, prefsID int64) error {
	q := `DELETE FROM user_preferences WHERE id = $1`

	_, err := p.db.ExecContext(ctx, q, prefsID)
	if err != nil {
		return err
	}

	return nil
}
