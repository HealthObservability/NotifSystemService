package callbackservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
	"github.com/HealthObservability/NotifSystemService/internal/services/notifservice"
)

type CallbackService interface {
	ProcessCallback(ctx context.Context, callback domains.SenderCallback) error
}

type Service struct {
	adapters     *adapters.Adapters
	notifService notifservice.NotifService
}

func New(a *adapters.Adapters, notifS notifservice.NotifService) CallbackService {
	return &Service{adapters: a, notifService: notifS}
}

func (s *Service) ProcessCallback(ctx context.Context, callback domains.SenderCallback) error {
	//TODO implement me
	panic("implement me")
}
