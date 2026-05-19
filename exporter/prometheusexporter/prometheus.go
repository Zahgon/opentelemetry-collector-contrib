// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"

import (
	"context"
	"errors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type prometheusExporter struct {
	config       Config
	name         string
	endpoint     string
	shutdownFunc func(ctx context.Context) error
	handler      http.Handler
	collector    *collector
	registry     *prometheus.Registry
	settings     component.TelemetrySettings
	stopCh       chan struct{} // signals the background metric cleanup goroutine to stop
}

var errBlankPrometheusAddress = errors.New("expecting a non-blank address to run the Prometheus metrics handler")

func newPrometheusExporter(config *Config, set exporter.Settings) (*prometheusExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return error early because newCollector
// will call logger.Error if it fails to build the namespace.

func (pe *prometheusExporter) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Start a background goroutine that periodically evicts expired metric families.
// Without this, cleanup only happens during Collect() (i.e. when Prometheus scrapes).
// If no scraper is active, stale entries in metricFamilies accumulate indefinitely,
// causing unbounded memory growth. See https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/41123

func (pe *prometheusExporter) ConsumeMetrics(_ context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (pe *prometheusExporter) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
