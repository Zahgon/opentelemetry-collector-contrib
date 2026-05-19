// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azureblobexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	// the format of encoded telemetry data
	formatTypeJSON  = "json"
	formatTypeProto = "proto"
)

// NewFactory creates a factory for Azure Blob exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createLogsExporter(ctx context.Context,
	params exporter.Settings,
	config component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func createMetricsExporter(ctx context.Context,
	params exporter.Settings,
	config component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createTracesExporter(ctx context.Context,
	params exporter.Settings,
	config component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}
