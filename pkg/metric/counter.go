package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Counter interface {
	With(lvs ...string) Counter
	Inc()
	Add(delta float64)
}

var _ Counter = (*counter)(nil)

type counter struct {
	cv  *prometheus.CounterVec
	lvs []string
}

// Counter new a prometheus counter and returns Counter.
func NewRegisterCounter(cv *prometheus.CounterVec) Counter {
	prometheus.MustRegister(cv)
	return &counter{
		cv: cv,
	}
}

func (c *counter) With(lvs ...string) Counter {
	return &counter{
		cv:  c.cv,
		lvs: lvs,
	}
}

func (c *counter) Inc() {
	c.cv.WithLabelValues(c.lvs...).Inc()
}

func (c *counter) Add(delta float64) {
	c.cv.WithLabelValues(c.lvs...).Add(delta)
}
