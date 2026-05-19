// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
)

const (
	zapEndpointKey = "endpoint"
)

// NewFactory creates a factory for the exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// By default we disable resilience options on loadbalancing exporter level
// to maintain compatibility with workflow in previous versions

func buildExporterConfig(cfg *Config, endpoint string) otlpexporter.Config {
	_ = "STUB: not implemented"
	return *new(otlpexporter.Config)
}

func buildExporterSettings(typ component.Type, params exporter.Settings, endpoint string) exporter.Settings {
	_ = "STUB: not implemented"
	return *new(exporter.Settings)
}

func buildExporterResilienceOptions(options []exporterhelper.Option, cfg *Config) []exporterhelper.Option {
	_ = "STUB: not implemented"
	return nil
}

func createTracesExporter(ctx context.Context, params exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createLogsExporter(ctx context.Context, params exporter.Settings, cfg component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func createMetricsExporter(ctx context.Context, params exporter.Settings, cfg component.Config) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}
