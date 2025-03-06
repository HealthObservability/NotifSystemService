package domains

type SenderCallback struct {
	NotifID int64  `json:"notif_id" validate:"required"`
	Status  string `json:"status" validate:"required"`
}
