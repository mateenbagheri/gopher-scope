package handlers

import (
	nethttp "net/http"

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
