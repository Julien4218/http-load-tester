package observability

import (
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	nrApp   *newrelic.Application
	promReg *prometheus.Registry
)

func Init() {
	// Initialize New Relic agent.
	appName := os.Getenv("NEW_RELIC_APP_NAME")
	if appName == "" {
		appName = "http-load-tester (local)"
	}
	licenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if licenseKey == "" {
		log.Warn("environment variable NEW_RELIC_LICENSE_KEY not set, skipping instrumentation")
		return
	}

	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		func(cfg *newrelic.Config) {
			// Set specific Config fields inside a custom ConfigOption.
			cfg.Attributes.Enabled = false
			cfg.ApplicationLogging.Enabled = true
			cfg.Logger = nrlogrus.StandardLogger()
		},
	)
	if err != nil {
		log.Error(err)
		return
	}

	promReg = prometheus.NewRegistry()

	// Wrap the handler with New Relic instrumentation.
	http.HandleFunc(newrelic.WrapHandleFunc(nrApp, "/metrics", promhttp.Handler().ServeHTTP))
}

func AddGaugeMetric(gauge prometheus.Gauge) {
	if promReg != nil {
		promReg.MustRegister(gauge)
	}
}

func AddCounterMetric(counter prometheus.Counter) {
	promReg.MustRegister(counter)
	{
		promReg.MustRegister(counter)
	}
}
