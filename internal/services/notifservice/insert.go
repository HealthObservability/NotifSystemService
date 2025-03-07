package notifservice

import (
	"context"
	"errors"
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

func createSendDue(lastSent time.Time, interval string) (time.Time, error) {
	sendDue := lastSent
	switch interval {
	case IntervalNoRepeat:
		sendDue = lastSent
	case IntervalMinute:
		sendDue = lastSent.Add(time.Minute)
	case IntervalHalfHour:
		sendDue = lastSent.Add(30 * time.Minute)
	case IntervalHour:
		sendDue = lastSent.Add(time.Hour)
	case IntervalDay:
		sendDue = lastSent.AddDate(0, 0, 1)
	case IntervalWeek:
		sendDue = lastSent.AddDate(0, 0, 7)
	case IntervalMonth:
		sendDue = lastSent.AddDate(0, 1, 0)
	default:
		return time.Time{}, errors.New("unknown interval")
	}

	return sendDue, nil
}

func (s *Service) AddNotification(ctx context.Context, n domains.Notif) (int64, error) {
	n.CreatedAt = time.Now()

	notifID, err := s.storage.Notifications.InsertNotification(ctx, n)
	if err != nil {
		return 0, err
	}

	var notifState domains.NotifState
	notifState.NotifID = notifID
	notifState.CreateDate = n.CreatedAt
	notifState.SendDue = n.Since
	notifState.Status = NotSentStatus

	_, err = s.storage.Notifications.InsertNotifState(ctx, notifState)
	if err != nil {
		return 0, err
	}

	return notifID, nil
}
