// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"context"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type azureMonitorExporter struct {
	config           *Config
	transportChannel appinsights.TelemetryChannel
	settings         component.TelemetrySettings
	logger           *zap.Logger
	packer           *metricPacker
}

func (exporter *azureMonitorExporter) Start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (exporter *azureMonitorExporter) Shutdown(_ context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (exporter *azureMonitorExporter) consumeLogs(_ context.Context, logData plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Flush the transport channel to force the telemetry to be sent

func (exporter *azureMonitorExporter) consumeMetrics(_ context.Context, metricData pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Flush the transport channel to force the telemetry to be sent

type traceVisitor struct {
	processed int
	err       error
	exporter  *azureMonitorExporter
}

// Called for each tuple of Resource, InstrumentationScope, and Span
func (v *traceVisitor) visit(
	resource pcommon.Resource,
	scope pcommon.InstrumentationScope,
	span ptrace.Span,
) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// record the error and short-circuit

// This is a fire and forget operation

func (exporter *azureMonitorExporter) consumeTraces(_ context.Context, traceData ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// Flush the transport channel to force the telemetry to be sent
