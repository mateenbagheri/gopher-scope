package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			reqID := generateRequestID()
			c.Set("request_id", reqID)
			c.Response().Header().Set("X-Request-ID", reqID)
			return next(c)
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

			var status int
			if err != nil {
				if httpErr, ok := err.(*echo.HTTPError); ok {
					status = httpErr.Code
				} else {
					status = 500
				}
			}

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
