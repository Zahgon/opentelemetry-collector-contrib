// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuredataexplorerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.uber.org/zap"
)

const (
	// The value of "type" key in configuration.
	managedIngestType  = "managed"
	queuedIngestTest   = "queued"
	otelDb             = "oteldb"
	defaultMetricTable = "OTELMetrics"
	defaultLogTable    = "OTELLogs"
	defaultTraceTable  = "OTELTraces"
	metricsType        = 1
	logsType           = 2
	tracesType         = 3
)

// Creates a factory for the ADX Exporter
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

// Create default configurations
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

// call the common exporter function in baseexporter. This ensures that the client and the ingest
// are initialized and the metrics struct are available for operations

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// call the common exporter function in baseexporter. This ensures that the client and the ingest
// are initialized and the metrics struct are available for operations

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config,
) (exp exporter.Logs, err error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// call the common exporter function in baseexporter. This ensures that the client and the ingest
// are initialized and the metrics struct are available for operations

func setDefaultIngestionType(config *Config, logger *zap.Logger) {
	_ = "STUB: not implemented"
	// If ingestion type is not set , it falls back to queued ingestion.
	// This form of ingestion is always available on all clusters
	return
}
