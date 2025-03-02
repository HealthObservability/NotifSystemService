package notifservice

import (
	"context"
	"errors"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

type Service struct {
	storage *storages.Storage
}

func NewNotifService(s *storages.Storage) *Service {
	return &Service{s}
}

func (s *Service) GetNotifStates(ctx context.Context) ([]domains.NotifStatus, error) {
	notifs, err := s.storage.NotificationStates.GetNotifications(ctx)
	if err != nil {
		return nil, err
	}

	if notifs == nil {
		notifs = make([]domains.NotifStatus, 0)
	}

	return notifs, nil
}

func (s *Service) GetFullNotifs(context.Context, []domains.NotifStatus) ([]domains.Notification, error) {
	return []domains.Notification{}, nil
}

func (s *Service) UpdateNotifications(ctx context.Context, id, msg string) error {
	return errors.New("not implemented")
}
