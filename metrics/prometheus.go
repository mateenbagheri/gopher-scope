package metrics

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	RequestCounter   *prometheus.CounterVec
	RequestDuration  *prometheus.HistogramVec
	ActiveRequests   prometheus.Gauge
	HTTPErrorCounter *prometheus.CounterVec
}

func NewMetrics(namespace string) *Metrics {
	constLabels := prometheus.Labels{"service": namespace}

	metric := &Metrics{
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

		HTTPErrorCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace:   namespace,
				Name:        "http_errors_total",
				Help:        "totla number of http errors",
				ConstLabels: constLabels,
			},
			[]string{"method", "path", "status"},
		),
	}

	return metric
}

func (m *Metrics) MetricsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			log.Println("METRICS MIDDLEWARE HIT:", c.Path())

			if c.Path() == "/metrics" {
				return next(c)
			}

			start := time.Now()
			m.ActiveRequests.Inc()
			defer m.ActiveRequests.Dec()

			err := next(c)

			statusStr := c.Request().Response.Status
			status, err := strconv.Atoi(statusStr)
			if status == 0 {
				status = http.StatusOK
			}

			method := c.Request().Method
			path := c.Path()

			m.RequestCounter.WithLabelValues(
				method, path, statusStr,
			).Inc()

			m.RequestDuration.WithLabelValues(
				method, path, statusStr,
			).Observe(time.Since(start).Seconds())

			log.Println("Status HIT:", statusStr)
			if status >= 400 {
				m.HTTPErrorCounter.WithLabelValues(
					method, path, statusStr,
				).Inc()
			}

			return err
		}
	}
}
