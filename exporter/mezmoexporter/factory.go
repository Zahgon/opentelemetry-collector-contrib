// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mezmoexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/mezmoexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a factory for Mezmo exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// Create a default Memzo config
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Create a log exporter for exporting to Mezmo
func createLogsExporter(ctx context.Context, settings exporter.Settings, exporterConfig component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// explicitly disable since we rely on http.Client timeout logic.
