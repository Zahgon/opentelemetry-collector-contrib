// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package googlemanagedprometheusexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlemanagedprometheusexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	defaultTimeout = 12 * time.Second // Consistent with Cloud Monitoring's timeout
)

// NewFactory creates a factory for the googlemanagedprometheus exporter
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// createDefaultConfig creates the default configuration for exporter.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetricsExporter creates a metrics exporter based on this config.
func createMetricsExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

// Disable exporterhelper Timeout, since we are using a custom mechanism
// within exporter itself
