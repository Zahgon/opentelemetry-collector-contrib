// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package influxdbexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/influxdbexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a factory for InfluxDB exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// defaults per suggested:
// https://docs.influxdata.com/influxdb/cloud-serverless/write-data/best-practices/optimize-writes/#batch-writes

func createTraceExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createMetricsExporter(ctx context.Context, set exporter.Settings, config component.Config) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createLogsExporter(ctx context.Context, set exporter.Settings, config component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
