package services

import (
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/ports"
	"github.com/HealthObservability/NotifSystemService/internal/services/callbacksvc"
	"github.com/HealthObservability/NotifSystemService/internal/services/notifservice"
	"github.com/HealthObservability/NotifSystemService/internal/services/sendersvc"
	"github.com/HealthObservability/NotifSystemService/internal/services/userprefsvc"
)

type Service struct {
	notifservice.NotifService
	sendersvc.SenderService
	callbacksvc.CallbackService
	userprefsvc.UserPref
}

func MustNew(db ports.Database, a *adapters.Adapters) *Service {
	userPrefSvc := userprefsvc.New(db)
	notifS := notifservice.NewNotifService(db, userPrefSvc)
	return &Service{
		NotifService:    notifS,
		SenderService:   sendersvc.New(),
		CallbackService: callbacksvc.New(notifS),
		UserPref:        userPrefSvc,
	}
}
