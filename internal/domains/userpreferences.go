package domains

import "time"

type UserPrefs struct {
	ID               *int64    `json:"id"`
	UserID           *int64    `json:"user_id"`
	DontDisturbStart time.Time `json:"dont_disturb_start"`
	DontDisturbEnd   time.Time `json:"dont_disturb_end"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
