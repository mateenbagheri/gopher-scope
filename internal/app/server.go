package app

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"go.uber.org/zap"

	httpx "github.com/mateenbagheri/gopher-scope/internal/http"
	"github.com/mateenbagheri/gopher-scope/metrics"
)

func newServer(
	logger *zap.Logger,
	promMetrics *metrics.Metrics,
) *echo.Echo {
	e := echo.New()

	httpx.RegisterMiddleware(e, logger, promMetrics)
	httpx.RegisterRoutes(e)

	return e
}

func startServer(e *echo.Echo, logger *zap.Logger) {
	const port = ":1323"

	logger.Info("starting server", zap.String("port", port))

	if err := e.Start(port); err != nil && err != http.ErrServerClosed {
		logger.Fatal("server failed", zap.Error(err))
	}
}
