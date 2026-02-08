package app

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func waitForShutdown(logger *zap.Logger) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")
}
