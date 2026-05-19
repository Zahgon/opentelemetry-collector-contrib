// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package rabbitmqexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/rabbitmqexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	defaultConnectionTimeout          = time.Second * 10
	defaultConnectionHeartbeat        = time.Second * 5
	defaultPublishConfirmationTimeout = time.Second * 5

	spansRoutingKey   = "otlp_spans"
	metricsRoutingKey = "otlp_metrics"
	logsRoutingKey    = "otlp_logs"

	defaultSpansConnectionName   = "otel-collector-spans"
	defaultMetricsConnectionName = "otel-collector-metrics"
	defaultLogsConnectionName    = "otel-collector-logs"
)

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
	return *new(exporter.Traces), nil
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func getRoutingKeyOrDefault(config *Config, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

func newPublisherFactory(set exporter.Settings) publisherFactory {
	_ = "STUB: not implemented"
	return *new(publisherFactory)
}

func newTLSFactory(config *Config) tlsFactory { _ = "STUB: not implemented"; return *new(tlsFactory) }
