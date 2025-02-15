package checker

import (
	"context"
	"notifservice/pkg/logger"
	"time"
)

const repeatTime = 1 * time.Second

type Service struct {
}

func NewCheckerService() *Service {
	return &Service{}
}

func (s *Service) GetNotifications(ctx context.Context) {
	// TODO implement me
	if ctx.Err() != nil {
		panic("implement me")
	}
}

func (s *Service) LoopRepeater(ctx context.Context) error {
	log := logger.Logger.WithField("op", "checker.LoopRepeater")
	// TODO implement me
	// panic("implement me")
	ticker := time.NewTicker(repeatTime)

	for {
		select {
		case <-ctx.Done():
			log.Info("stop loop repeater")
			return nil
		case <-ticker.C:
			log.Info("mock checker repeater")
		}
	}
}
