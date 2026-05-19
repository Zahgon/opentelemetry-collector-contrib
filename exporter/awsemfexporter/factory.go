// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsemfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/featuregate"
)

var defaultNoRollupfg = featuregate.GlobalRegistry().MustRegister("awsemf.nodimrollupdefault", featuregate.StageAlpha,
	featuregate.WithRegisterFromVersion("v0.83.0"),
	featuregate.WithRegisterDescription("Changes the default AWS EMF Exporter Dimension rollup option to "+
		"NoDimensionRollup"))

// NewFactory creates a factory for AWS EMF exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// CreateDefaultConfig creates the default configuration for exporter.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createMetricsExporter creates a metrics exporter based on this config.
func createMetricsExporter(ctx context.Context, params exporter.Settings, config component.Config) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}
