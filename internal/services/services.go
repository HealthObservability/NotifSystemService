package services

import (
	"context"
	"notifservice/internal/services/checker"
	"notifservice/internal/storages"
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

func New(s *storages.Storage) *Service {
	return &Service{
		Checker: checker.NewCheckerService(),
		s:       s,
	}
}
