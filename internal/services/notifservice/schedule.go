package notifservice

import (
	"context"
	"errors"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"golang.org/x/sync/errgroup"
	"time"
)

func calculateScheduledTime(lastScheduled time.Time, interval string) (time.Time, error) {
	log := logger.GetLogger().WithField("op", "calculateScheduledTime")
	now := time.Now().UTC()
	lastScheduled = lastScheduled.UTC()

	timeDiff := lastScheduled.Sub(now)
	if timeDiff < 0 {
		timeDiff = -1 * timeDiff
	}
	log.Debugf("Time difference: %v", timeDiff)

	var scheduledTime time.Time

	switch interval {
	case IntervalNoRepeat:
		return time.Time{}, nil

	case IntervalMinute:
		if timeDiff >= time.Minute {
			elapsedIntervals := int(timeDiff / time.Minute)
			log.Debugf("Elapsed intervals: %v", elapsedIntervals)
			scheduledTime = lastScheduled.Add(time.Duration(elapsedIntervals) * time.Minute)
		} else {
			scheduledTime = lastScheduled.Add(time.Minute)
		}

	case IntervalHalfHour:
		if timeDiff >= 30*time.Minute {
			elapsedIntervals := int(timeDiff / (30 * time.Minute))
			scheduledTime = lastScheduled.Add(time.Duration(elapsedIntervals) * 30 * time.Minute)
		} else {
			scheduledTime = lastScheduled.Add(30 * time.Minute)
		}

	case IntervalHour:
		if timeDiff >= time.Hour {
			elapsedIntervals := int(timeDiff / time.Hour)
			scheduledTime = lastScheduled.Add(time.Duration(elapsedIntervals) * time.Hour)
		} else {
			scheduledTime = lastScheduled.Add(time.Hour)
		}

	case IntervalDay:
		if timeDiff >= 24*time.Hour {
			elapsedIntervals := int(timeDiff / (24 * time.Hour))
			scheduledTime = lastScheduled.AddDate(0, 0, elapsedIntervals)
		} else {
			scheduledTime = lastScheduled.AddDate(0, 0, 1)
		}

	case IntervalWeek:
		if timeDiff >= 7*24*time.Hour {
			elapsedIntervals := int(timeDiff / (7 * 24 * time.Hour))
			scheduledTime = lastScheduled.AddDate(0, 0, elapsedIntervals*7)
		} else {
			scheduledTime = lastScheduled.AddDate(0, 0, 7)
		}

	case IntervalMonth:
		if timeDiff >= 30*24*time.Hour {
			elapsedIntervals := int(timeDiff / (30 * 24 * time.Hour))
			scheduledTime = lastScheduled.AddDate(0, elapsedIntervals, 0)
		} else {
			scheduledTime = lastScheduled.AddDate(0, 1, 0)
		}

	default:
		return time.Time{}, errors.New("unknown interval")
	}

	log.Debugf("New scheduled time: %v", scheduledTime)

	return scheduledTime, nil
}

func (s *Service) ScheduleNotifications(ctx context.Context, n []domains.Notif) error {
	g, ctx := errgroup.WithContext(ctx)
	for _, notif := range n {
		notif := notif
		g.Go(func() error {
			lastScheduled := notif.LastScheduled
			interval := notif.RepeatInterval
			if interval == IntervalNoRepeat {
				return nil
			}

			scheduledTime, err := calculateScheduledTime(lastScheduled, interval)
			if err != nil {
				return err
			}

			state := domains.NotifState{
				NotifID:    notif.ID,
				SendDue:    scheduledTime,
				CreateDate: time.Now().UTC(),
				Status:     NotSentStatus,
			}

			if err = s.storage.Notifications.ScheduleNotification(ctx, state); err != nil {
				return err
			}

			return nil
		})
	}

	return g.Wait()
}
