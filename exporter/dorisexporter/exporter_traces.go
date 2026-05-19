// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"context"
	_ "embed" // for SQL file embedding

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

//go:embed sql/traces_ddl.sql
var tracesDDL string

//go:embed sql/traces_view.sql
var tracesView string

//go:embed sql/traces_graph_ddl.sql
var tracesGraphDDL string

//go:embed sql/traces_graph_job.sql
var tracesGraphJob string

// dTrace Trace to Doris
type dTrace struct {
	ServiceName        string         `json:"service_name"`
	Timestamp          string         `json:"timestamp"`
	ServiceInstanceID  string         `json:"service_instance_id"`
	TraceID            string         `json:"trace_id"`
	SpanID             string         `json:"span_id"`
	TraceState         string         `json:"trace_state"`
	ParentSpanID       string         `json:"parent_span_id"`
	SpanName           string         `json:"span_name"`
	SpanKind           string         `json:"span_kind"`
	EndTime            string         `json:"end_time"`
	Duration           int64          `json:"duration"`
	SpanAttributes     map[string]any `json:"span_attributes"`
	Events             []*dEvent      `json:"events"`
	Links              []*dLink       `json:"links"`
	StatusMessage      string         `json:"status_message"`
	StatusCode         string         `json:"status_code"`
	ResourceAttributes map[string]any `json:"resource_attributes"`
	ScopeName          string         `json:"scope_name"`
	ScopeVersion       string         `json:"scope_version"`
}

// dEvent Event to Doris
type dEvent struct {
	Timestamp  string         `json:"timestamp"`
	Name       string         `json:"name"`
	Attributes map[string]any `json:"attributes"`
}

// dLink Link to Doris
type dLink struct {
	TraceID    string         `json:"trace_id"`
	SpanID     string         `json:"span_id"`
	TraceState string         `json:"trace_state"`
	Attributes map[string]any `json:"attributes"`
}

type tracesExporter struct {
	*commonExporter
}

func newTracesExporter(logger *zap.Logger, cfg *Config, set component.TelemetrySettings) *tracesExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *tracesExporter) pushTraceData(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) pushTraceDataInternal(ctx context.Context, traces []*dTrace, label string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracesExporter) formatDropTraceGraphJob() string { _ = "STUB: not implemented"; return "" }

func (e *tracesExporter) formatTraceGraphJob() string { _ = "STUB: not implemented"; return "" }
