package app

import (
	"context"
	"fmt"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/services"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"golang.org/x/sync/errgroup"
	"time"
)

const timeout = time.Second * 10

func Run(ctx context.Context, a *adapters.Adapters, s *services.Service) {
	log := logger.Logger.WithField("op", "app.Run")

	if err := s.LoopRepeater(ctx); err != nil {
		log.Fatal(err.Error())
	}

	g, _ := errgroup.WithContext(ctx)

	// g.Go(func() error {
	// 	<-gCtx.Done()
	// 	log.Info("shutting down postgres...")
	//
	// 	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()
	//
	// 	if err := a.Postgres.Shutdown(); err != nil {
	// 		return fmt.Errorf("failed to shutdown HTTP adapter: %w", err)
	// 	}
	//
	// 	log.Info("http adapter shutdown completed")
	// 	return nil
	// })

	if err := g.Wait(); err != nil {
		log.Errorf("shutdown error: %v", err)
	}

	g.Go(func() error {
		tCtx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := a.Postgres.Shutdown(tCtx); err != nil {
			return fmt.Errorf("failed to shutdown postgres adapter: %w", err)
		}
		log.Info("postgres adapter shutdown completed")
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Errorf("shutdown error: %v", err)
	}
}
