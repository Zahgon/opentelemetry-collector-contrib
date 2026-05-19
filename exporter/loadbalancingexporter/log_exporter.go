// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

var _ exporter.Logs = (*logExporterImp)(nil)

type logExporterImp struct {
	loadBalancer *loadBalancer

	logger     *zap.Logger
	started    bool
	shutdownWg sync.WaitGroup
	telemetry  *metadata.TelemetryBuilder
}

// Create new logs exporter
func newLogsExporter(params exporter.Settings, cfg component.Config) (*logExporterImp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*logExporterImp) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (e *logExporterImp) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logExporterImp) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *logExporterImp) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logExporterImp) consumeLog(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// every log may not contain a traceID
// generate a random traceID as balancingKey
// so the log can be routed to a random backend

func traceIDFromLogs(ld plog.Logs) pcommon.TraceID {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID)
}

func random() pcommon.TraceID { _ = "STUB: not implemented"; return *new(pcommon.TraceID) }
