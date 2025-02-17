package checker

import (
	"context"
	"errors"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

type Service struct {
	storage *storages.Storage
}

func NewNotifService(s *storages.Storage) *Service {
	return &Service{s}
}

func (s *Service) GetNotifications(ctx context.Context) ([]any, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) UpdateNotifications(ctx context.Context, id, msg string) error {
	return errors.New("not implemented")
}
