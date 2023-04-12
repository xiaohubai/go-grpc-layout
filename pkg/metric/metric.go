package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Counter interface {
	With(lvs ...string) Counter
	Inc()
	Add(delta float64)
}

// Gauge is metrics gauge.
type Gauge interface {
	With(lvs ...string) Gauge
	Inc()
	Set(value float64)
	Add(delta float64)
	Sub(delta float64)
}

// Histogram is metrics histogram.
type Histogram interface {
	With(lvs ...string) Histogram
	Observe(float64)
}

var (
	namespace = "metric"
	labels    = []string{"service", "status", "path", "method"}

	// QPS
	ReqCount = NewCounter(
		prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests made.",
		}, labels),
	)

	// 接口响应时间
	ReqDuration = NewHistogram(
		prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
		}, labels),
	)

	// 当前正在处理请求的QPS
	CurReqCount = NewGauge(
		prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "http_requests_in_flight",
			Help:      "Current number of http requests in flight.",
		}, labels),
	)
)
