package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
)

const (
	METRIC_PORT int = 5000
)

const (
	metric_namespace = "httploadtester"
	metric_subsystem = "results"
)

type Metrics struct {
	TestCount     prometheus.Counter
	SuccessCount  prometheus.Counter
	FailCount     prometheus.Counter
	ElapsedTimeMs prometheus.Gauge
}

var (
	metrics *Metrics
)

func init() {
	metrics = &Metrics{
		TestCount:     createCounter("test_count"),
		SuccessCount:  createCounter("success_count"),
		FailCount:     createCounter("fail_count"),
		ElapsedTimeMs: createGauge("elapsed_time_ms"),
	}
	log.Infof("Metrics created")
}

func createCounter(name string) prometheus.Counter {
	counter := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metric_namespace,
		Subsystem: metric_subsystem,
		Name:      name})
	return counter
}

func createGauge(name string) prometheus.Gauge {
	gauge := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: metric_namespace,
		Subsystem: metric_subsystem,
		Name:      name})
	return gauge
}

func GetMetrics() *Metrics {
	return metrics
}
