// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultEncoding     = "otlp_proto"
	defaultTraceTopic   = "otlp_spans"
	defaultMetricsTopic = "otlp_metrics"
	defaultLogsTopic    = "otlp_logs"
	defaultConsumerName = ""
	defaultSubscription = "otlp_subscription"
	defaultServiceURL   = "pulsar://localhost:6650"
)

// FactoryOption applies changes to PulsarExporterFactory.
type FactoryOption func(factory *pulsarReceiverFactory)

// withTracesUnmarshalers adds Unmarshalers.
func withTracesUnmarshalers(tracesUnmarshalers ...TracesUnmarshaler) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// withMetricsUnmarshalers adds MetricsUnmarshalers.
func withMetricsUnmarshalers(metricsUnmarshalers ...MetricsUnmarshaler) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// withLogsUnmarshalers adds LogsUnmarshalers.
func withLogsUnmarshalers(logsUnmarshalers ...LogsUnmarshaler) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

// NewFactory creates Pulsar receiver factory.
func NewFactory(options ...FactoryOption) receiver.Factory {
	_ = "STUB: not implemented"
	return *new(receiver.Factory)
}

type pulsarReceiverFactory struct {
	tracesUnmarshalers  map[string]TracesUnmarshaler
	metricsUnmarshalers map[string]MetricsUnmarshaler
	logsUnmarshalers    map[string]LogsUnmarshaler
}

func (f *pulsarReceiverFactory) createTracesReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func (f *pulsarReceiverFactory) createMetricsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func (f *pulsarReceiverFactory) createLogsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
