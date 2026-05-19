// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package logzioexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a factory for Logz.io exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Default to gzip compression

// We almost read 0 bytes, so no need to tune ReadBufferSize.

func getListenerURL(region, dataType string) string { _ = "STUB: not implemented"; return "" }

func generateEndpoint(cfg *Config, dataType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createTracesExporter(_ context.Context, params exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createLogsExporter(_ context.Context, params exporter.Settings, cfg component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
