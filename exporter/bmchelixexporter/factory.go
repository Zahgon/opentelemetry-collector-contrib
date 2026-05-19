// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package bmchelixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// create BMC Helix Exporter factory
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// creates the default configuration for the BMC Helix exporter
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// creates an exporter.Metrics that records observability metrics for BMC Helix
func createMetricsExporter(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}
