package app

import "github.com/mateenbagheri/gopher-scope/metrics"

func initMetrics() *metrics.Metrics {
	return metrics.NewMetrics("gopherscope")
}
