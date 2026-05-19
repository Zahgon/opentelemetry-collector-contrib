// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bmchelixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	om "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter/internal/operationsmanagement"
)

// metricsExporter is responsible for exporting metrics to BMC Helix
type metricsExporter struct {
	config            *Config
	logger            *zap.Logger
	version           string
	telemetrySettings component.TelemetrySettings
	producer          *om.MetricsProducer
	client            *om.MetricsClient
}

// newMetricsExporter instantiates a new metrics exporter for BMC Helix
func newMetricsExporter(config *Config, createSettings exporter.Settings) (*metricsExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pushMetrics is invoked by the OpenTelemetry Collector to push metrics to BMC Helix
func (me *metricsExporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// start is invoked during service start
func (me *metricsExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize and store the MetricsProducer

// Initialize and store the MetricsClient
