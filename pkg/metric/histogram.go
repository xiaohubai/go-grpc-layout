package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var _ Histogram = (*histogram)(nil)

type histogram struct {
	hv  *prometheus.HistogramVec
	lvs []string
}

// NewHistogram new a prometheus histogram and returns Histogram.
func NewHistogram(hv *prometheus.HistogramVec) Histogram {
	prometheus.MustRegister(hv)
	return &histogram{
		hv: hv,
	}
}

func (h *histogram) With(lvs ...string) Histogram {
	return &histogram{
		hv:  h.hv,
		lvs: lvs,
	}
}

func (h *histogram) Observe(value float64) {
	h.hv.WithLabelValues(h.lvs...).Observe(value)
}
