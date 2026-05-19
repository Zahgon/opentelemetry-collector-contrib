// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// factory.go - OpenTelemetry Collector factory implementation
package isolationforestprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/isolationforestprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"
)

const (
	typeStr   = "isolationforest"
	stability = component.StabilityLevelAlpha
)

func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createTracesProcessor(
	_ context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func createMetricsProcessor(
	_ context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func createLogsProcessor(
	_ context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

type tracesProcessor struct {
	*isolationForestProcessor
	nextConsumer consumer.Traces
	logger       *zap.Logger
}

func (tp *tracesProcessor) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (tp *tracesProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (tp *tracesProcessor) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (*tracesProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

type metricsProcessor struct {
	*isolationForestProcessor
	nextConsumer consumer.Metrics
	logger       *zap.Logger
}

func (mp *metricsProcessor) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *metricsProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *metricsProcessor) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (*metricsProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

type logsProcessor struct {
	*isolationForestProcessor
	nextConsumer consumer.Logs
	logger       *zap.Logger
}

func (lp *logsProcessor) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (lp *logsProcessor) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (lp *logsProcessor) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*logsProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}
