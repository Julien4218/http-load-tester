package observability

import (
	"fmt"
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	nrApp *newrelic.Application
)

func Init() {
	nrApp = initAgent()

	if nrApp == nil {
		http.Handle("/metrics", promhttp.Handler())
		go http.ListenAndServe(fmt.Sprintf(":%d", METRIC_PORT), nil)
		log.Infof("metrics endpoint available at localhost:%d", METRIC_PORT)
	} else {
		// Wrap the handler with New Relic instrumentation.
		http.HandleFunc(newrelic.WrapHandleFunc(nrApp, "/metrics", promhttp.Handler().ServeHTTP))
		go http.ListenAndServe(fmt.Sprintf(":%d", METRIC_PORT), nil)
		log.Infof("metrics endpoint available at localhost:%d wrap with newrelic", METRIC_PORT)
	}
}

func initAgent() *newrelic.Application {
	// Initialize New Relic agent.
	appName := os.Getenv("NEW_RELIC_APP_NAME")
	if appName == "" {
		appName = "http-load-tester (local)"
	}
	licenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if licenseKey == "" {
		log.Warn("environment variable NEW_RELIC_LICENSE_KEY not set, skipping instrumentation")
		return nil
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		func(cfg *newrelic.Config) {
			// Set specific Config fields inside a custom ConfigOption.
			cfg.Attributes.Enabled = false
			cfg.ApplicationLogging.Enabled = true
			cfg.Logger = nrlogrus.StandardLogger()
			cfg.ApplicationLogging.Metrics.Enabled = true
			cfg.Logger.DebugEnabled()
		},
	)
	if err != nil {
		log.Error(err)
		return nil
	}
	return app
}
