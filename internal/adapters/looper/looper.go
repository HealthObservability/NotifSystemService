package looper

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"time"
)

const loopTime = 1 * time.Second

type Looper struct {
}

func NewLooper() *Looper {
	return &Looper{}
}

func (l *Looper) Start(ctx context.Context, handler func(ctx context.Context)) {
	log := logger.Logger.WithField("op", "Looper.LoopRepeater")
	ticker := time.NewTicker(loopTime)

	for {
		select {
		case <-ctx.Done():
			log.Info("stop l")
			return
		case <-ticker.C:
			handler(ctx)
		}
	}
}
