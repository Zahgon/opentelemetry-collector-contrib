// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	exp "go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
)

func newLogsExporter(cfg component.Config, set exp.Settings) (*logsExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type logsExporter struct {
	grpcLogsExporter plogotlp.GRPCClient
	httpLogsExporter httpLogsExporter
	*signalExporter
}

func (e *logsExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsExporter) pushLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsExporter) enhanceContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
