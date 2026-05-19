// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory creates a factory for AlibabaCloud LogService exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// CreateDefaultConfig creates the default configuration for exporter.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(
	_ context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createMetricsExporter(
	_ context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exp exporter.Metrics, err error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createLogsExporter(
	_ context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exp exporter.Logs, err error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
