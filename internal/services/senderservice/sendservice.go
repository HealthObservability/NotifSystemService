package senderservice

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/domains"
)

type SenderService interface {
	SendNotifications(ctx context.Context, notifs []domains.Notif) error
}

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s Service) SendNotifications(ctx context.Context, notifs []domains.Notif) error {
	//TODO implement me
	panic("implement me")

	// TODO get each notif
	// TODO SEND each notif to sender service

	// TODO BRAINSTORM
	// MAY BE I CAN JUST SEND ALL NOTIFS TO NOTIFICATION SERVICE AND IT WILL SEND A MESSAGES TO EACH USER.
	// S?
}
