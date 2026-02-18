package http

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/zap"

	"github.com/mateenbagheri/gopher-scope/metrics"
	custom "github.com/mateenbagheri/gopher-scope/middleware"
)

func RegisterMiddleware(
	e *echo.Echo,
	logger *zap.Logger,
	metrics *metrics.Metrics,
) {
	e.Use(middleware.Recover())
	e.Use(custom.RequestIDMiddleware())
	e.Use(custom.StructuredLoggingMiddleware(logger))
	e.Use(metrics.MetricsMiddleware())
}
