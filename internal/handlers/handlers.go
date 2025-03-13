package handlers

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

type Handler struct {
	service *services.Service
}

func New(s *services.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Run(ctx context.Context, a *adapters.Adapters) {
	log := logger.GetLogger().WithField("op", "app.Run")

	a.HTTPServ.SetHandlers(h)

	// run here all handlers and services
	go a.Looper.Start(ctx, h.handleLoop)
	go a.HTTPServ.MustRun()

	g, gCtx := errgroup.WithContext(ctx)
	<-gCtx.Done()

	// graceful shutdown
	tCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	g.Go(func() error {
		<-gCtx.Done()
		log.Info("shutting down http server")

		if err := a.HTTPServ.Shutdown(tCtx); err != nil {
			return fmt.Errorf("failed to shutdown HTTP adapter: %w", err)
		}

		log.Info("http adapter shutdown completed")
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Errorf("shutdown handlers error: %v", err)
	}

	g.Go(func() error {
		if err := a.Database.Shutdown(tCtx); err != nil {
			return fmt.Errorf("failed to shutdown postgres adapter: %w", err)
		}
		log.Info("postgres adapter shutdown completed")
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Errorf("shutdown error: %v", err)
	}
}
