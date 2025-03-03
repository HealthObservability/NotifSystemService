package notifservice

import (
	"context"
	"errors"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

const (
	SentStatus    = "sent"
	PendingStatus = "pending"
)

type Service struct {
	storage *storages.Storage
}

func NewNotifService(s *storages.Storage) *Service {
	return &Service{s}
}

func (s *Service) GetNotifs(ctx context.Context, n []domains.NotifState) ([]domains.Notif, error) {
	length := len(n)
	ids := make([]uint64, length)
	for i := 0; i < length; i++ {
		ids[i] = n[i].NotificationID
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
	notifs, err := s.storage.NotificationStates.GetNotifStates(ctx)
	if err != nil {
		return nil, err
	}

	if notifs == nil {
		notifs = make([]domains.NotifState, 0)
	}

	return notifs, nil
}

func (s *Service) UpdateNotifications(context.Context, string, string) error {
	return errors.New("not implemented")
}
