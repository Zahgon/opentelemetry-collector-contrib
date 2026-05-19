// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
)

func newTracesExporter(cfg component.Config, set exporter.Settings) (*tracesExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type tracesExporter struct {
	grpcTracesExporter ptraceotlp.GRPCClient
	httpTracesExporter httpTracesExporter
	*signalExporter
}

func (e *tracesExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) pushTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) enhanceContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
