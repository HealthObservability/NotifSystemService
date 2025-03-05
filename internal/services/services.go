package services

import (
	"github.com/HealthObservability/NotifSystemService/internal/services/notifservice"
	"github.com/HealthObservability/NotifSystemService/internal/services/senderservice"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

type Service struct {
	notifservice.NotifService
	senderservice.SenderService
}

func MustNew(s *storages.Storage) *Service {
	return &Service{
		NotifService:  notifservice.NewNotifService(s),
		SenderService: senderservice.New(),
	}
}
