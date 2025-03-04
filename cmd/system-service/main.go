package main

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/config"
	"github.com/HealthObservability/NotifSystemService/internal/handlers"
	"github.com/HealthObservability/NotifSystemService/internal/services"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

// var (
//	isLocal = flag.Bool("local", false, "is it local? can make logs pretty")
//	isDebug = flag.Bool("debug", true, "wanna see debug mode?")
// )

func main() {
	logger.InitDefaultLogger(logger.WithDebug())

	cfg := config.MustConfigure("./config.yaml")

	ad := adapters.NewAdapters(cfg)
	st := storages.MustNew(ad)
	srv := services.MustNew(st)
	h := handlers.New(srv)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	defer stop()

	h.Run(ctx, ad)
}
