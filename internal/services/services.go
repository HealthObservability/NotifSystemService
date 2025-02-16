package services

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/services/checker"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

type Checker interface {
	GetNotifications(context.Context)
	LoopRepeater(context.Context) error
}

type Updater interface {
	UpdateNotifications(ctx context.Context, id, message string)
}

type Service struct {
	Checker
	Updater

	s *storages.Storage
}

func MustNew(s *storages.Storage) *Service {
	return &Service{
		Checker: checker.NewCheckerService(s),
		s:       s,
	}
}
