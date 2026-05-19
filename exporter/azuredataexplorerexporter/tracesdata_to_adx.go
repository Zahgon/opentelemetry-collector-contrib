// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuredataexplorerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type adxTrace struct {
	TraceID            string         // TraceID associated to the Trace
	SpanID             string         // SpanID associated to the Trace
	ParentID           string         // ParentID associated to the Trace
	SpanName           string         // The SpanName of the Trace
	SpanStatus         string         // The SpanStatus Code associated to the Trace
	SpanStatusMessage  string         // The SpanStatusMessage associated to the Trace
	SpanKind           string         // The SpanKind of the Trace
	StartTime          string         // The start time of the occurrence. Formatted into string as RFC3339Nano
	EndTime            string         // The end time of the occurrence. Formatted into string as RFC3339Nano
	ResourceAttributes map[string]any // JSON Resource attributes that can then be parsed.
	TraceAttributes    map[string]any // JSON attributes that can then be parsed.
	Events             []*event       // Array containing the events in a span
	Links              []*link        // Array containing the link in a span
}

type event struct {
	EventName       string
	Timestamp       string
	EventAttributes map[string]any
}

type link struct {
	TraceID            string
	SpanID             string
	TraceState         string
	SpanLinkAttributes map[string]any
}

func mapToAdxTrace(resource pcommon.Resource, scope pcommon.InstrumentationScope, spanData ptrace.Span) *adxTrace {
	_ = "STUB: not implemented"
	return nil
}

func getEventsData(sd ptrace.Span) []*event { _ = "STUB: not implemented"; return nil }

func getLinksData(sd ptrace.Span) []*link { _ = "STUB: not implemented"; return nil }
