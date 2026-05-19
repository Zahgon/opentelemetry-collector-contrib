// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logicmonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"

	traces "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter/internal/traces"
)

type tracesExporter struct {
	config   *Config
	sender   *traces.Sender
	settings component.TelemetrySettings
	cancel   context.CancelFunc
}

// newTracesExporter creates new Logicmonitor Traces Exporter.
func newTracesExporter(_ context.Context, cfg component.Config, set exporter.Settings) *tracesExporter {
	_ = "STUB: not implemented"
	return nil

	// client construction is deferred to start
}

func (e *tracesExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) PushTraceData(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }
