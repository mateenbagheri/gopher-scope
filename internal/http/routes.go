package http

import (
	"github.com/labstack/echo/v5"

	"github.com/mateenbagheri/gopher-scope/internal/http/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	h := handlers.New()

	e.GET("/", h.Root)
	e.GET("/metrics", h.Metrics)
	e.GET("/unreliable", h.UnreliableEndpoint)
	e.GET("/slow", h.SlowEndpoint)
}
