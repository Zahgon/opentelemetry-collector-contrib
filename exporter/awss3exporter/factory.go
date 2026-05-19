// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3exporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
)

// TODO: Find a place for this to be shared.
type baseMetricsExporter struct {
	component.Component
	consumer.Metrics
}

// TODO: Find a place for this to be shared.
type baseLogsExporter struct {
	component.Component
	consumer.Logs
}

// TODO: Find a place for this to be shared.
type baseTracesExporter struct {
	component.Component
	consumer.Traces
}

// NewFactory creates a factory for S3 exporter.
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

// checkAndCastConfig checks the configuration type and casts it to the S3 exporter Config struct.
func checkAndCastConfig(c component.Config) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
