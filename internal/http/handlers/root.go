package handlers

import (
	"math/rand"
	"net/http"
	nethttp "net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/mateenbagheri/gopher-scope/service"
)

type Handlers struct {
	greeter *service.Greeter
}

func New() *Handlers {
	return &Handlers{
		greeter: service.NewGreeter(),
	}
}

func (h *Handlers) Root(c *echo.Context) error {
	msg := h.greeter.Greet()

	return c.String(nethttp.StatusOK, msg)
}

func (h *Handlers) UnreliableEndpoint(c *echo.Context) error {
	randomNum := rand.Intn(100)

	if randomNum < 70 {
		errors := []struct {
			status  int
			message string
		}{
			{http.StatusBadRequest, "Bad request: invalid parameters"},
			{http.StatusUnauthorized, "Unauthorized: invalid credentials"},
			{http.StatusForbidden, "Forbidden: insufficient permissions"},
			{http.StatusNotFound, "Resource not found"},
			{http.StatusConflict, "Conflict: resource already exists"},
			{http.StatusTooManyRequests, "Too many requests: rate limit exceeded"},
			{http.StatusInternalServerError, "Internal server error"},
			{http.StatusServiceUnavailable, "Service unavailable"},
			{http.StatusGatewayTimeout, "Gateway timeout"},
			{http.StatusTeapot, "I'm a teapot"},
		}

		selectedError := errors[rand.Intn(len(errors))]

		return c.JSON(selectedError.status, map[string]any{
			"error":   selectedError.message,
			"status":  selectedError.status,
			"success": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Request succeeded!",
		"status":  http.StatusOK,
		"success": true,
	})
}

// SlowEndpoint sometimes takes a long time (70% chance of delay between 5-10 seconds)
func (h *Handlers) SlowEndpoint(c *echo.Context) error {
	randomNum := rand.Intn(100)

	if randomNum < 70 {
		delaySeconds := rand.Intn(6) + 5
		delay := time.Duration(delaySeconds) * time.Second

		time.Sleep(delay)

		return c.JSON(http.StatusOK, map[string]any{
			"message":       "Request processed slowly",
			"delay_seconds": delaySeconds,
			"status":        http.StatusOK,
			"success":       true,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Request processed instantly!",
		"status":  http.StatusOK,
		"success": true,
	})
}
