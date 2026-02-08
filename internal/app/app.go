package app

import (
	"github.com/mateenbagheri/gopher-scope/logging"
	"go.uber.org/zap"
)

func Run() {
	logger := initLogger()
	defer logging.SyncLogger()

	metrics := initMetrics()

	server := newServer(logger, metrics)

	go startServer(server, logger)

	waitForShutdown(logger)

	logger.Info("server exited properly")
}

func initLogger() *zap.Logger {
	if err := logging.InitLogger(); err != nil {
		panic(err)
	}

	logger := logging.GetLogger()
	return logger
}
