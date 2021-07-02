package metric

import (
	"sync/atomic"
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
			metricName,
			nil, nil,
		),
	}
}

func (m *QpsMetric) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.metric
}

func (m *QpsMetric) Collect(ch chan<- prometheus.Metric) {
	metricValue := float64(atomic.LoadUint64(&qps))
	ch <- prometheus.MustNewConstMetric(m.metric, prometheus.GaugeValue, metricValue)
}

var (
	qps uint64
)

func IncQps() {
	atomic.AddUint64(&qps, 1)
}

func cal() {
	tick := time.NewTicker(1 * time.Second)
	for {
		//@TODO clean up
		select {
		case <-tick.C:
			atomic.StoreUint64(&qps, 0)
		}
	}
}
