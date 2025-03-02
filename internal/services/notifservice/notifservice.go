package notifservice

import (
	"context"
	"errors"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

const SentStatus = "sent"

type Service struct {
	storage *storages.Storage
}

func NewNotifService(s *storages.Storage) *Service {
	return &Service{s}
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

func (s *Service) ChangeNotifState(context.Context, uint64, string) error {
	// todo implement
	return nil
}

func (s *Service) GetFullNotifs(context.Context, []domains.NotifState) ([]domains.Notification, error) {
	return []domains.Notification{}, nil
}

func (s *Service) UpdateNotifications(ctx context.Context, id, msg string) error {
	return errors.New("not implemented")
}
