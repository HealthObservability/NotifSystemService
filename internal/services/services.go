package services

import (
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/ports"
	"github.com/HealthObservability/NotifSystemService/internal/services/callbackservice"
	"github.com/HealthObservability/NotifSystemService/internal/services/notifservice"
	"github.com/HealthObservability/NotifSystemService/internal/services/senderservice"
	"github.com/HealthObservability/NotifSystemService/internal/services/userprefservice"
)

type Service struct {
	notifservice.NotifService
	senderservice.SenderService
	callbackservice.CallbackService
	userprefservice.UserPreferences
}

func MustNew(db ports.Database, a *adapters.Adapters) *Service {
	notifS := notifservice.NewNotifService(db)
	return &Service{
		NotifService:    notifS,
		SenderService:   senderservice.New(),
		CallbackService: callbackservice.New(notifS),
		UserPreferences: userprefservice.New(db),
	}
}
