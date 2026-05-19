// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunk // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/splunk"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// hecEvent is a data structure holding a span event to export explicitly to Splunk HEC.
type hecEvent struct {
	Attributes map[string]any    `json:"attributes,omitempty"`
	Name       string            `json:"name"`
	Timestamp  pcommon.Timestamp `json:"timestamp"`
}

// hecLink is a data structure holding a span link to export explicitly to Splunk HEC.
type hecLink struct {
	Attributes map[string]any `json:"attributes,omitempty"`
	TraceID    string         `json:"trace_id"`
	SpanID     string         `json:"span_id"`
	TraceState string         `json:"trace_state"`
}

// hecSpanStatus is a data structure holding the status of a span to export explicitly to Splunk HEC.
type hecSpanStatus struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// hecSpan is a data structure used to export explicitly a span to Splunk HEC.
type hecSpan struct {
	TraceID    string            `json:"trace_id"`
	SpanID     string            `json:"span_id"`
	ParentSpan string            `json:"parent_span_id"`
	Name       string            `json:"name"`
	Attributes map[string]any    `json:"attributes,omitempty"`
	EndTime    pcommon.Timestamp `json:"end_time"`
	Kind       string            `json:"kind"`
	Status     hecSpanStatus     `json:"status"`
	StartTime  pcommon.Timestamp `json:"start_time"`
	Events     []hecEvent        `json:"events,omitempty"`
	Links      []hecLink         `json:"links,omitempty"`
}

func SpanToSplunkEvent(resource pcommon.Resource, span ptrace.Span, mapping HecToOtelAttrs, source, sourceType, index string) *Event {
	_ = "STUB: not implemented"
	return nil
}

// ignore

func toHecSpan(span ptrace.Span) hecSpan { _ = "STUB: not implemented"; return *new(hecSpan) }
