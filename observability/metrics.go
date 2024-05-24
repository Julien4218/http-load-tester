package observability

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
	log "github.com/sirupsen/logrus"
)

const (
	METRIC_PORT int = 5000
)

const (
	metric_namespace = "httploadtester"
	metric_subsystem = "results"
)

type NewRelicCounter interface {
	Inc()
}

type NewRelicGauge interface {
	Set(value float64)
}

type NewRelicMetric struct {
	name string
}

type Metrics struct {
	TestCount     NewRelicCounter
	SuccessCount  NewRelicCounter
	FailCount     NewRelicCounter
	ElapsedTimeMs NewRelicGauge
	RpmPace       NewRelicGauge
}

var (
	metrics  *Metrics
	hostname string
	pid      int
)

func init() {
	hostname, _ = os.Hostname()
	pid = os.Getpid()
	metrics = &Metrics{
		TestCount:     createCounter("test_count"),
		SuccessCount:  createCounter("success_count"),
		FailCount:     createCounter("fail_count"),
		ElapsedTimeMs: createGauge("elapsed_time_ms"),
		RpmPace:       createGauge("rpm_pace"),
	}
	log.Infof("Metrics created")
}

func createCounter(name string) NewRelicCounter {
	return &NewRelicMetric{fmt.Sprintf("%s_%s_%s", metric_namespace, metric_subsystem, name)}
}

func createGauge(name string) NewRelicGauge {
	return &NewRelicMetric{fmt.Sprintf("Custom/%s_%s_%s", metric_namespace, metric_subsystem, name)}
}

func GetMetrics() *Metrics {
	return metrics
}

func (m *NewRelicMetric) Inc() {
	harvester.RecordMetric(telemetry.Count{
		Name:      m.name,
		Value:     1,
		Timestamp: time.Now(),
		Attributes: map[string]interface{}{
			"source": hostname,
			"pid":    pid,
		},
	})
}

func (m *NewRelicMetric) Set(value float64) {
	harvester.RecordMetric(telemetry.Gauge{
		Name:      m.name,
		Value:     value,
		Timestamp: time.Now(),
		Attributes: map[string]interface{}{
			"source": hostname,
			"pid":    pid,
		},
	})
}

func HarvestNow() {
	harvester.HarvestNow(context.Background())
}
