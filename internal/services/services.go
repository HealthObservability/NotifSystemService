package services

import "context"

type Checker interface {
	LookUpNotifications()
}

type Updater interface {
	UpdateNotifications(ctx context.Context, id, message string)
}

type Service struct {
	Checker
	Updater
}

func New() *Service {
	return &Service{}
}
