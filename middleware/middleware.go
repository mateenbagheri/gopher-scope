package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RequestInfoFillerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			reqID := generateRequestID()
			c.Set("request_id", reqID)
			c.Response().Header().Set("X-Request-ID", reqID)
			err := next(c)

			return err
		}
	}
}

func StructuredLoggingMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			start := time.Now()

			reqID := c.Get("request_id").(string)

			reqLogger := logger.With(
				zap.String("request_id", reqID),
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
			)

			c.Set("logger", reqLogger)

			reqLogger.Info("request_started",
				zap.String("user_agent", c.Request().UserAgent()),
			)

			err := next(c)
			duration := time.Since(start)

			resp, err := echo.UnwrapResponse(c.Response())
			var status int
			status = resp.Status

			logFields := []zapcore.Field{
				zap.Int("status", status),
				zap.Duration("duration", duration),
				zap.Int64("duration_ms", duration.Microseconds()),
			}

			if status >= 500 {
				reqLogger.Error("request_completed", logFields...)
			} else if status >= 400 {
				reqLogger.Warn("request_completed", logFields...)
			} else {
				reqLogger.Info("request_completed", logFields...)
			}

			return err
		}
	}
}

func generateRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
