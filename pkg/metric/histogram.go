package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var _ Histogram = (*histogram)(nil)

type histogram struct {
	hv  *prometheus.HistogramVec
	lvs []string
}

// Histogram is metrics histogram.
type Histogram interface {
	With(lvs ...string) Histogram
	Observe(float64)
}

// NewHistogram new a prometheus histogram and returns Histogram.
func NewRegisterHistogram(hv *prometheus.HistogramVec) Histogram {
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
