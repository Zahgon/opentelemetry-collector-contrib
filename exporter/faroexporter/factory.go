// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/faroexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a factory for Faro exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTraces(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createLogs(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
