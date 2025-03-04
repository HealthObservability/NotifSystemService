package notifservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"time"
)

func (s *Service) AddNotification(ctx context.Context, n domains.Notif) (int64, error) {
	n.CreatedAt = time.Now()

	notifID, err := s.storage.Notifications.InsertNotification(ctx, n)
	if err != nil {
		return 0, err
	}

	var notifState domains.NotifState
	notifState.NotifID = notifID
	notifState.CreateDate = n.CreatedAt
	notifState.Status = NotSentStatus

	_, err = s.storage.Notifications.InsertNotifState(ctx, notifState)
	if err != nil {
		return 0, err
	}

	return notifID, nil
}
