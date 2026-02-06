package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/mateenbagheri/gopher-scope/logging"
	"github.com/mateenbagheri/gopher-scope/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	customMiddleware "github.com/mateenbagheri/gopher-scope/middleware"
)

func main() {
	err := logging.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logging.SyncLogger()

	logger := logging.GetLogger()
	logger.Info("app is starting ...")

	promMetrics := metrics.NewMetrics("mytestapp")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(customMiddleware.RequestIDMiddleware())
	e.Use(customMiddleware.StructuredLoggingMiddleware(logger))
	e.Use(promMetrics.MetricsMiddleware())

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	go func() {
		port := ":1323"
		logger.Info("starting server", zap.String("port", port))
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			logger.Fatal("shutting down the server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	logger.Info("server exited properly")
}
