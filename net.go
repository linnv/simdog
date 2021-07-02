package simdog

import (
	"github.com/linnv/logx"
	"github.com/prometheus/client_golang/prometheus"
)

type NetMetric struct {
	metric *prometheus.Desc
}

func NewNetMetric(metricName string) *NetMetric {
	return &NetMetric{
		metric: prometheus.NewDesc("socket_tcp_"+metricName,
			metricName+" tcp connection count",
			nil, nil,
		),
	}
}

func (m *NetMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.metric
}

func (m *NetMetric) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64

	if v, err := CurProc.Connections(); err == nil {
		metricValue = float64(len(v))
	} else {
		logx.Debugf("net err: %+v\n", v)
	}
	ch <- prometheus.MustNewConstMetric(m.metric, prometheus.GaugeValue, metricValue)
}
