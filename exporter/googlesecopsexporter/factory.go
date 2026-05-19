// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlesecopsexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlesecopsexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a new Google SecOps exporter factory.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// createDefaultConfig creates the default configuration for the google secops exporter.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createLogsExporter creates a new log exporter based on this config.
func createLogsExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg component.Config,
) (exp exporter.Logs, err error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
