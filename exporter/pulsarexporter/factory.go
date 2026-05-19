// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

//go:build !aix

package pulsarexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	defaultTracesTopic  = "otlp_spans"
	defaultMetricsTopic = "otlp_metrics"
	defaultLogsTopic    = "otlp_logs"
	defaultEncoding     = "otlp_proto"
	defaultBroker       = "pulsar://localhost:6650"
)

// FactoryOption applies changes to pulsarExporterFactory.
type FactoryOption func(factory *pulsarExporterFactory)

// withTracesMarshalers adds tracesMarshalers.
func withTracesMarshalers(tracesMarshalers ...TracesMarshaler) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// NewFactory creates Pulsar exporter factory.
func NewFactory(options ...FactoryOption) exporter.Factory {
	_ = "STUB: not implemented"
	return *new(exporter.Factory)
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// using an empty topic to track when it has not been set by user, default is based on traces or metrics.

type pulsarExporterFactory struct {
	tracesMarshalers  map[string]TracesMarshaler
	metricsMarshalers map[string]MetricsMarshaler
	logsMarshalers    map[string]LogsMarshaler
}

func (f *pulsarExporterFactory) createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// Disable exporterhelper Timeout, because we cannot pass a Context to the Producer,
// and will rely on the Pulsar Producer Timeout logic.

func (f *pulsarExporterFactory) createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

// Disable exporterhelper Timeout, because we cannot pass a Context to the Producer,
// and will rely on the sarama Pulsar Timeout logic.

func (f *pulsarExporterFactory) createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// Disable exporterhelper Timeout, because we cannot pass a Context to the Producer,
// and will rely on the Pulsar Producer Timeout logic.
