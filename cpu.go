package simdog

import (
	"github.com/linnv/logx"
	"github.com/prometheus/client_golang/prometheus"
)

type CpuMetric struct {
	metric *prometheus.Desc
}

func NewCpuMetric(metricName string) *CpuMetric {
	return &CpuMetric{
		metric: prometheus.NewDesc("cpu_"+metricName,
			metricName+" cpu count",
			nil, nil,
		),
	}
}

func (m *CpuMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.metric
}

func (m *CpuMetric) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64

	if v, err := CurProc.CPUPercent(); err == nil {
		metricValue = float64(v)
	} else {
		logx.Debugf("cpu err: %+v\n", v)
	}
	ch <- prometheus.MustNewConstMetric(m.metric, prometheus.GaugeValue, metricValue)
}
