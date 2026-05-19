// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
)

type dataConsumer interface {
	consume(ctx context.Context, event *azureEvent) error
	setNextLogsConsumer(nextLogsConsumer consumer.Logs)
	setNextMetricsConsumer(nextLogsConsumer consumer.Metrics)
	setNextTracesConsumer(nextTracesConsumer consumer.Traces)
}

type eventLogsUnmarshaler interface {
	UnmarshalLogs(event *azureEvent) (plog.Logs, error)
}

type eventMetricsUnmarshaler interface {
	UnmarshalMetrics(event *azureEvent) (pmetric.Metrics, error)
}

type eventTracesUnmarshaler interface {
	UnmarshalTraces(event *azureEvent) (ptrace.Traces, error)
}

type eventhubReceiver struct {
	eventHandler        *eventhubHandler
	signal              pipeline.Signal
	logger              *zap.Logger
	logsUnmarshaler     eventLogsUnmarshaler
	metricsUnmarshaler  eventMetricsUnmarshaler
	tracesUnmarshaler   eventTracesUnmarshaler
	nextLogsConsumer    consumer.Logs
	nextMetricsConsumer consumer.Metrics
	nextTracesConsumer  consumer.Traces
	obsrecv             *receiverhelper.ObsReport
}

func (receiver *eventhubReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *eventhubReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *eventhubReceiver) setNextLogsConsumer(nextLogsConsumer consumer.Logs) {
	_ = "STUB: not implemented"
	return
}

func (receiver *eventhubReceiver) setNextMetricsConsumer(nextMetricsConsumer consumer.Metrics) {
	_ = "STUB: not implemented"
	return
}

func (receiver *eventhubReceiver) setNextTracesConsumer(nextTracesConsumer consumer.Traces) {
	_ = "STUB: not implemented"
	return
}

func (receiver *eventhubReceiver) consume(ctx context.Context, event *azureEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *eventhubReceiver) consumeLogs(ctx context.Context, event *azureEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *eventhubReceiver) consumeMetrics(ctx context.Context, event *azureEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *eventhubReceiver) consumeTraces(ctx context.Context, event *azureEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func newReceiver(
	signal pipeline.Signal,
	logsUnmarshaler eventLogsUnmarshaler,
	metricsUnmarshaler eventMetricsUnmarshaler,
	tracesUnmarshaler eventTracesUnmarshaler,
	eventHandler *eventhubHandler,
	settings receiver.Settings,
) (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}
