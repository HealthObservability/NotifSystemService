package services

import (
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/services/callbackservice"
	"github.com/HealthObservability/NotifSystemService/internal/services/notifservice"
	"github.com/HealthObservability/NotifSystemService/internal/services/senderservice"
	"github.com/HealthObservability/NotifSystemService/internal/services/userprefservice"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

type Service struct {
	notifservice.NotifService
	senderservice.SenderService
	callbackservice.CallbackService
	userprefservice.UserPreferences
}

func MustNew(s *storages.Storage, a *adapters.Adapters) *Service {
	notifS := notifservice.NewNotifService(s)
	return &Service{
		NotifService:    notifS,
		SenderService:   senderservice.New(a),
		CallbackService: callbackservice.New(a, notifS),
		UserPreferences: userprefservice.New(a),
	}
}
