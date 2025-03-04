package domains

import "time"

// NotifState is an object contains in postgres db and references to Notif.
type NotifState struct {
	ID         int64     `json:"id"`
	NotifID    int64     `json:"notification_id"`
	SendDue    time.Time `json:"send_due"`
	CreateDate time.Time `json:"create_date"`
	Status     string    `json:"status"` // notSent / Sent / cancelled
}

// Notif is an object that contains in any key-value or document-based db.
type Notif struct {
	ID             int64      `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	LastSent       *time.Time `json:"last_sent_timestamp"`
	ToID           int64      `json:"to_id" validate:"required"`
	Message        string     `json:"message" validate:"required"`
	RepeatInterval string     `json:"repeat_interval" validate:"required,oneof=no-repeat minute half-hour hour day week month"`
	Since          time.Time  `json:"since_timestamp" validate:"required"`
}
