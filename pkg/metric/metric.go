package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "metric"
	labels    = []string{"env", "service", "protocol", "path", "method", "status"}
	// 请求数统计
	ReqCount = NewRegisterCounter(
		prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "requests_total",
			Help:      "Total number of Requests.",
		}, []string{"env", "service", "protocol", "path", "method"}),
	)

	RespCount = NewRegisterCounter(
		prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "responses_total",
			Help:      "Total number of Responses.",
		}, labels),
	)

	// 响应时间分布统计
	RespDurationHistogram = NewRegisterHistogram(
		prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "responses_duration_histogram",
			Help:      "responses latencies in histogram seconds.",
		}, labels),
	)

	// 响应最大时间统计
	RespDurationGauge = NewRegisterGauge(
		prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "responses_duration_gauge",
			Help:      "responses latencies in gauge seconds",
		}, labels),
	)

	Count = NewRegisterCounter(
		prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "count_by_name_total",
			Help:      "Total number of Requests.",
		}, []string{"name"}),
	)
)
