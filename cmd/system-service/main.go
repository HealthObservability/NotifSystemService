package main

import (
	"context"
	"flag"
	"notifservice/internal/services"
	"notifservice/pkg/logger"
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

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	defer stop()

	s := services.New(nil)

	logger.Logger.Info("starting system service")

	if err := s.LoopRepeater(ctx); err != nil {
		logger.Logger.Fatal(err)
	}
}
