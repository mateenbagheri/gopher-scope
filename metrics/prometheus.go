package metrics

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	RequestCounter  *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
	ActiveRequests  prometheus.Gauge
}

func NewMetrics(namespace string) *Metrics {
	constLabels := prometheus.Labels{"service": namespace}

	m := &Metrics{
		RequestCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace:   namespace,
				Name:        "http_request_total",
				Help:        "total request count for HTTP requests",
				ConstLabels: constLabels,
			},
			[]string{"method", "path", "status"},
		),

		RequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace:   namespace,
				Name:        "http_request_duration_in_seconds",
				Help:        "duration of HTTP requests in second",
				ConstLabels: constLabels,
				Buckets:     prometheus.DefBuckets,
			},
			[]string{"method", "path", "status"},
		),

		ActiveRequests: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Name:        "http_request_active",
				Help:        "number of current active requests",
				ConstLabels: constLabels,
			},
		),
	}

	return m
}

func (m *Metrics) MetricsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			start := time.Now()
			m.ActiveRequests.Inc()

			err := next(c)
			duration := time.Since(start).Seconds()

			var status int
			if err != nil {
				if httpErr, ok := err.(*echo.HTTPError); ok {
					status = httpErr.Code
				} else {
					status = 500
				}
			}

			statusStr := strconv.Itoa(status)

			// Recording process now: adding metrics to prometheus basically
			m.RequestCounter.WithLabelValues(
				c.Request().Method,
				c.Path(),
				statusStr,
			).Inc()

			m.RequestDuration.WithLabelValues(
				c.Request().Method,
				c.Path(),
				statusStr,
			).Observe(duration)

			m.ActiveRequests.Dec()
			return err
		}
	}
}
