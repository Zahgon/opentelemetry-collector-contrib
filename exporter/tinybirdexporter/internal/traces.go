// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tinybirdexporter/internal"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type traceSignal struct {
	ResourceSchemaURL  string            `json:"resource_schema_url"`
	ResourceAttributes map[string]string `json:"resource_attributes"`
	ServiceName        string            `json:"service_name"`
	ScopeSchemaURL     string            `json:"scope_schema_url"`
	ScopeName          string            `json:"scope_name"`
	ScopeVersion       string            `json:"scope_version"`
	ScopeAttributes    map[string]string `json:"scope_attributes"`
	TraceID            string            `json:"trace_id"`
	SpanID             string            `json:"span_id"`
	ParentSpanID       string            `json:"parent_span_id"`
	TraceState         string            `json:"trace_state"`
	TraceFlags         uint32            `json:"trace_flags"`
	SpanName           string            `json:"span_name"`
	SpanKind           string            `json:"span_kind"`
	SpanAttributes     map[string]string `json:"span_attributes"`
	StartTime          string            `json:"start_time"`
	// used when users choose the StartTime-to-EndTime approach
	EndTime string `json:"end_time,omitempty"`
	// used when users choose the StartTime-plus-Duration approach
	Duration         int64               `json:"duration,omitempty"`
	StatusCode       string              `json:"status_code"`
	StatusMessage    string              `json:"status_message"`
	EventsTimestamp  []string            `json:"events_timestamp"`
	EventsName       []string            `json:"events_name"`
	EventsAttributes []map[string]string `json:"events_attributes"`
	LinksTraceID     []string            `json:"links_trace_id"`
	LinksSpanID      []string            `json:"links_span_id"`
	LinksTraceState  []string            `json:"links_trace_state"`
	LinksAttributes  []map[string]string `json:"links_attributes"`
}

func convertEvents(events ptrace.SpanEventSlice) (timestamps, names []string, attributes []map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func convertLinks(links ptrace.SpanLinkSlice) (traceIDs, spanIDs, states []string, attrs []map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func ConvertTraces(td ptrace.Traces, encoder Encoder) error { _ = "STUB: not implemented"; return nil }
