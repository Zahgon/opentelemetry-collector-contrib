// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alertmanagerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alertmanagerexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a factory for Alertmanager exporter
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(ctx context.Context, set exporter.Settings, config component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}
