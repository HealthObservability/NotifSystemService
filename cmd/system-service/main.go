package main

import (
	"context"
	"flag"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/app"
	"github.com/HealthObservability/NotifSystemService/internal/config"
	"github.com/HealthObservability/NotifSystemService/internal/services"
	"github.com/HealthObservability/NotifSystemService/internal/storages"
	"github.com/HealthObservability/NotifSystemService/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

var (
	isLocal = flag.Bool("local", false, "is it local? can make logs pretty")
	isDebug = flag.Bool("debug", false, "wanna see debug mode?")
)

func main() {
	logger.InitLogger(isLocal, isDebug)

	cfg := config.MustConfigure("./config.yaml")

	adapt := adapters.NewAdapters(cfg)
	storage := storages.MustNew(adapt)
	service := services.MustNew(storage)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	defer stop()
	
	app.Run(ctx, adapt, service)
}
