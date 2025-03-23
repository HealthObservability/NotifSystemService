package main

import (
	"context"
	"github.com/HealthObservability/NotifSystemService/internal/adapters"
	"github.com/HealthObservability/NotifSystemService/internal/config"
	"github.com/HealthObservability/NotifSystemService/internal/handlers"
	"github.com/HealthObservability/NotifSystemService/internal/services"
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
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	defer stop()
	
	logger.InitDefaultLogger(logger.WithDebug())

	cfg := config.MustConfigure("./config.yaml")

	ad := adapters.NewAdapters(ctx, cfg)
	srv := services.MustNew(ad.Database, ad)
	h := handlers.New(srv)

	h.Run(ctx, ad.Database, ad.HTTPServ, ad.Looper)
}
