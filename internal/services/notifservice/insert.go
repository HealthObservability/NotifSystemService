package notifservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"time"
)

// no-repeat minute half-hour hour day week month check it at domains.Notif
const (
	IntervalNoRepeat = "no-repeat"
	IntervalMinute   = "minute"
	IntervalHalfHour = "half-hour"
	IntervalHour     = "hour"
	IntervalDay      = "day"
	IntervalWeek     = "week"
	IntervalMonth    = "month"
)

func (s *service) AddNotification(ctx context.Context, n domains.Notif) (int64, error) {
	n.CreatedAt = time.Now().UTC()
	n.Since = n.Since.UTC()
	n.LastScheduled = n.Since

	fmt.Println("n.Since:", n.Since)
	fmt.Println("n.LastScheduled:", n.LastScheduled)

	if n.Since.Before(n.CreatedAt) {
		return 0, errors.New("since time is in the past")
	}

	notifID, err := s.Db.InsertNotif(ctx, n)
	if err != nil {
		return 0, err
	}

	var notifState domains.NotifState
	notifState.NotifID = notifID
	notifState.CreateDate = n.CreatedAt
	notifState.SendDue = n.LastScheduled
	notifState.Status = NotSentStatus

	_, err = s.Db.InsertState(ctx, notifState)
	if err != nil {
		return 0, err
	}

	return notifID, nil
}
