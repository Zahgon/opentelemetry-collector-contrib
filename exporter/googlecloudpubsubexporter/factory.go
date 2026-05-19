// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package googlecloudpubsubexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	defaultTimeout = 12 * time.Second
)

// NewFactory creates a factory for Google Cloud Pub/Sub exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

var exporters = map[*Config]*pubsubExporter{}

func ensureExporter(params exporter.Settings, pCfg *Config) *pubsubExporter {
	_ = "STUB: not implemented"
	return nil
}

// we ignore the error here as the config is already validated with the same method

// createDefaultConfig creates the default configuration for exporter.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createMetricsExporter(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createLogsExporter(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
