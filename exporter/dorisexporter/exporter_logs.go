// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"context"
	_ "embed" // for SQL file embedding

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

//go:embed sql/logs_ddl.sql
var logsDDL string

//go:embed sql/logs_view.sql
var logsView string

// dLog Log to Doris
type dLog struct {
	ServiceName        string         `json:"service_name"`
	Timestamp          string         `json:"timestamp"`
	ServiceInstanceID  string         `json:"service_instance_id"`
	TraceID            string         `json:"trace_id"`
	SpanID             string         `json:"span_id"`
	SeverityNumber     int32          `json:"severity_number"`
	SeverityText       string         `json:"severity_text"`
	Body               string         `json:"body"`
	ResourceAttributes map[string]any `json:"resource_attributes"`
	LogAttributes      map[string]any `json:"log_attributes"`
	ScopeName          string         `json:"scope_name"`
	ScopeVersion       string         `json:"scope_version"`
}

type logsExporter struct {
	*commonExporter
}

func newLogsExporter(logger *zap.Logger, cfg *Config, set component.TelemetrySettings) *logsExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *logsExporter) pushLogData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *logsExporter) pushLogDataInternal(ctx context.Context, logs []*dLog, label string) error {
	_ = "STUB: not implemented"
	return nil
}
