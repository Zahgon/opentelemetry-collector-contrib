// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exceptionsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/exceptionsconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil"
)

type logsConnector struct {
	config Config

	// Additional dimensions to add to logs.
	dimensions []pdatautil.Dimension

	logsConsumer consumer.Logs
	component.StartFunc
	component.ShutdownFunc

	logger *zap.Logger
}

func newLogsConnector(logger *zap.Logger, config component.Config) *logsConnector {
	_ = "STUB: not implemented"
	return nil
}

// Capabilities implements the consumer interface.
func (*logsConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeTraces implements the consumer.Traces interface.
// It aggregates the trace data to generate logs.
func (c *logsConnector) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *logsConnector) exportLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*logsConnector) newScopeLogs(ld plog.Logs) plog.ScopeLogs {
	_ = "STUB: not implemented"
	return *new(plog.ScopeLogs)
}

func (c *logsConnector) attrToLogRecord(sl plog.ScopeLogs, serviceName string, span ptrace.Span, event ptrace.SpanEvent, resourceAttrs pcommon.Map) plog.LogRecord {
	_ = "STUB: not implemented"
	return *new(plog.LogRecord)
}

// Copy span attributes to the log record.

// Add common attributes to the log record.

// Add configured dimension attributes to the log record.

// Add stacktrace to the log record.
