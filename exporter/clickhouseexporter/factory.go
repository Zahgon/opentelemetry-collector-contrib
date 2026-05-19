// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package clickhouseexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/featuregate"
	"go.uber.org/zap"
)

// Deprecated: Use the `json` config option instead. This feature gate will be removed in a future version.
var featureGateJSON = featuregate.GlobalRegistry().MustRegister(
	"clickhouse.json",
	featuregate.StageDeprecated,
	featuregate.WithRegisterDescription("Deprecated: Use the `json` config option instead."),
	featuregate.WithRegisterToVersion("v0.149.0"),
)

// NewFactory creates a factory for the ClickHouse exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func useJSON(logger *zap.Logger, c *Config) bool { _ = "STUB: not implemented"; return false }

func createMetricExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}
