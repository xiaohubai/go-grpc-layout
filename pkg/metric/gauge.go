package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Gauge is metrics gauge.
type Gauge interface {
	With(lvs ...string) Gauge
	Inc()
	Set(value float64)
	Add(delta float64)
	Sub(delta float64)
}

var _ Gauge = (*gauge)(nil)

type gauge struct {
	gv  *prometheus.GaugeVec
	lvs []string
}

// NewGauge new a prometheus gauge and returns Gauge.
func NewRegisterGauge(gv *prometheus.GaugeVec) Gauge {
	prometheus.MustRegister(gv)
	return &gauge{
		gv: gv,
	}
}

func (g *gauge) With(lvs ...string) Gauge {
	return &gauge{
		gv:  g.gv,
		lvs: lvs,
	}
}

func (g *gauge) Set(value float64) {
	g.gv.WithLabelValues(g.lvs...).Set(value)
}

func (g *gauge) Inc() {
	g.gv.WithLabelValues(g.lvs...).Inc()
}

func (g *gauge) Add(delta float64) {
	g.gv.WithLabelValues(g.lvs...).Add(delta)
}

func (g *gauge) Sub(delta float64) {
	g.gv.WithLabelValues(g.lvs...).Sub(delta)
}
