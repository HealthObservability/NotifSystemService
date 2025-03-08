package notifservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

func (s *Service) GetNotifs(ctx context.Context, n []domains.NotifState) ([]domains.Notif, error) {
	length := len(n)

	if length == 0 {
		return []domains.Notif{}, nil
	}

	ids := make([]int64, length)
	for i := 0; i < length; i++ {
		ids[i] = n[i].NotifID
	}

	notifs, err := s.storage.GetNotifs(ctx, ids)
	if err != nil {
		return nil, err
	}

	if len(notifs) == 0 {
		return []domains.Notif{}, nil
	}

	return notifs, nil
}

func (s *Service) GetNotifStates(ctx context.Context) ([]domains.NotifState, error) {
	notifs, err := s.storage.NotificationStates.GetNotifStates(ctx, NotSentStatus)
	if err != nil {
		return nil, err
	}

	if notifs == nil {
		notifs = make([]domains.NotifState, 0)
	}

	return notifs, nil
}
