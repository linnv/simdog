package simdog

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type QpsMetric struct {
	metric *prometheus.Desc
}

func NewQpsMetric(metricName string) *QpsMetric {
	go cal()
	return &QpsMetric{
		metric: prometheus.NewDesc("http_qps_"+metricName,
			metricName+" http qps",
			nil, nil,
		),
	}
}

func (m *QpsMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.metric
}

func (m *QpsMetric) Collect(ch chan<- prometheus.Metric) {
	metricValue := float64(qps)
	ch <- prometheus.MustNewConstMetric(m.metric, prometheus.GaugeValue, metricValue)
}

var (
	gIntCur  uint64
	gIntLast uint64
	qps      uint64
)

func GetQps() uint64 {
	qps = gIntCur - gIntLast
	return qps
}

func IncQps() {
	gIntCur++
}

func cal() {
	tick := time.NewTicker(1 * time.Second)
	for {
		//@TODO clean up
		//@TODO mutex
		select {
		case <-tick.C:
			qps = gIntCur - gIntLast
			gIntLast = gIntCur
		}
	}
}
