package notifservice

import (
	"github.com/HealthObservability/NotifSystemService/internal/storages"
)

const (
	NotSentStatus = "not_sent"
	SentStatus    = "sent"
	PendingStatus = "pending"
)

type Service struct {
	storage *storages.Storage
}

func NewNotifService(s *storages.Storage) *Service {
	return &Service{s}
}
