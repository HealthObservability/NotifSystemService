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
	ToID           int64      `json:"to_id"`
	Message        string     `json:"message"`
	RepeatInterval string     `json:"repeat_interval"`
	LastSent       *time.Time `json:"last_sent_timestamp"`
	Since          time.Time  `json:"since_timestamp"`
}
