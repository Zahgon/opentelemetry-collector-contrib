// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package failoverconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type wrappedTracesConnector struct {
	consumer     consumer.Traces
	failoverCore *tracesFailover
}

type wrappedMetricsConnector struct {
	consumer     consumer.Metrics
	failoverCore *metricsFailover
}

type wrappedLogsConnector struct {
	consumer     consumer.Logs
	failoverCore *logsFailover
}

func (w *wrappedTracesConnector) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedTracesConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (w *wrappedTracesConnector) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedMetricsConnector) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedMetricsConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (w *wrappedMetricsConnector) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedLogsConnector) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedLogsConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (w *wrappedLogsConnector) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedTracesConnector) GetFailoverRouter() *tracesRouter {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedMetricsConnector) GetFailoverRouter() *metricsRouter {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedLogsConnector) GetFailoverRouter() *logsRouter {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedTracesConnector) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedMetricsConnector) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wrappedLogsConnector) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func newWrappedTracesConnector(consumer consumer.Traces, failoverCore *tracesFailover) *wrappedTracesConnector {
	_ = "STUB: not implemented"
	return nil
}

func newWrappedMetricsConnector(consumer consumer.Metrics, failoverCore *metricsFailover) *wrappedMetricsConnector {
	_ = "STUB: not implemented"
	return nil
}

func newWrappedLogsConnector(consumer consumer.Logs, failoverCore *logsFailover) *wrappedLogsConnector {
	_ = "STUB: not implemented"
	return nil
}
