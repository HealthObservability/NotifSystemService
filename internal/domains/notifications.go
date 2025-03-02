package domains

import "time"

// NotifState is an object contains in postgres db and references to Notification.
type NotifState struct {
	ID             uint64    `json:"id"`
	NotificationID string    `json:"notification_id"`
	SendDue        time.Time `json:"send_due"`
	CreateDate     time.Time `json:"create_date"`
	Status         string    `json:"status"` // notSent / Sent / cancelled
}

// Notification is an object that contains in any key-value or document-based db.
type Notification struct {
	ID        string `json:"id"`
	ToID      string `json:"to_id"`
	Message   string `json:"message"`
	Type      string `json:"type"` // alert or something
	CreatedAt uint64 `json:"created_at"`
	SendDue   uint64 `json:"send_due"`
}
