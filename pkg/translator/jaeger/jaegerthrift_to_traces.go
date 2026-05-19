// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger"

import (
	"github.com/jaegertracing/jaeger-idl/thrift-gen/jaeger"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var blankJaegerThriftSpan = new(jaeger.Span)

// ThriftToTraces transforms a Thrift trace batch into ptrace.Traces.
func ThriftToTraces(batches *jaeger.Batch) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func jThriftProcessToInternalResource(process *jaeger.Process, dest pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

// Handle special keys translations.

func jThriftSpansToInternal(spans []*jaeger.Span, dest ptrace.SpanSlice) {
	_ = "STUB: not implemented"
	return
}

// jThriftSpanParentID infers the parent span ID for a given span.
// Based on https://github.com/jaegertracing/jaeger/blob/8c61b6561f9057a199c1504606d8e68319ee7b31/model/span.go#L143
func jThriftSpanParentID(span *jaeger.Span) int64 { _ = "STUB: not implemented"; return 0 }

// If span.ParentSpanId undefined but there are references to the same trace,
// they can also be considered a parent, with CHILD_OF being higher priority.

// must be from the same trace

func jThriftSpanToInternal(span *jaeger.Span, dest ptrace.Span) { _ = "STUB: not implemented"; return }

// drop the attributes slice if all of them were replaced during translation

// jThriftTagsToInternalAttributes sets internal span links based on jaeger span references skipping excludeParentID
func jThriftTagsToInternalAttributes(tags []*jaeger.Tag, dest pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func jThriftLogsToSpanEvents(logs []*jaeger.Log, dest ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	return
}

func jThriftReferencesToSpanLinks(refs []*jaeger.SpanRef, excludeParentID int64, dest ptrace.SpanLinkSlice) {
	_ = "STUB: not implemented"
	return
}

// microsecondsToUnixNano converts epoch microseconds to pcommon.Timestamp
func microsecondsToUnixNano(ms int64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func jThriftRefTypeToAttribute(ref jaeger.SpanRefType) string { _ = "STUB: not implemented"; return "" }
