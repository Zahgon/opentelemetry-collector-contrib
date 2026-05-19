// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/exporterhelper/xexporterhelper"
	"go.opentelemetry.io/collector/exporter/xexporter"
)

const (
	defaultLogsTopic        = "otlp_logs"
	defaultLogsEncoding     = "otlp_proto"
	defaultMetricsTopic     = "otlp_metrics"
	defaultMetricsEncoding  = "otlp_proto"
	defaultTracesTopic      = "otlp_spans"
	defaultTracesEncoding   = "otlp_proto"
	defaultProfilesTopic    = "otlp_profiles"
	defaultProfilesEncoding = "otlp_proto"

	// partitioning metrics by resource attributes is disabled by default
	defaultPartitionMetricsByResourceAttributesEnabled = false
	// partitioning logs by resource attributes is disabled by default
	defaultPartitionLogsByResourceAttributesEnabled = false
	// partitioning logs by trace id is disabled by default
	defaultPartitionLogsByTraceIDEnabled = false
)

// NewFactory creates Kafka exporter factory.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *
	// Clone the config
	new(exporter.Traces), nil
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *
	// Clone the config
	new(exporter.Metrics), nil
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *
	// Clone the config
	new(exporter.Logs), nil
}

func createProfilesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (xexporter.Profiles, error) {
	_ = "STUB: not implemented"
	return *
	// Clone the config
	new(xexporter.Profiles), nil
}

func exporterhelperOptions(
	cfg Config,
	qbs xexporterhelper.QueueBatchSettings,
	startFunc component.StartFunc,
	shutdownFunc component.ShutdownFunc,
) []exporterhelper.Option {
	_ = "STUB: not implemented"
	return nil
}
