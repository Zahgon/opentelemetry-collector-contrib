// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package awskinesisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	defaultEncoding    = "otlp"
	defaultCompression = "none"
)

// NewFactory creates a factory for Kinesis exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func newTracesExporter(ctx context.Context, params exporter.Settings, conf component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func newMetricsExporter(ctx context.Context, params exporter.Settings, conf component.Config) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func newLogsExporter(ctx context.Context, params exporter.Settings, conf component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
