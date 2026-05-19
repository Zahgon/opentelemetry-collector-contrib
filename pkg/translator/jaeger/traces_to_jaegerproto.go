// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger"

import (
	"github.com/jaegertracing/jaeger-idl/model/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// ProtoFromTraces translates internal trace data into the Jaeger Proto for GRPC.
// Returns slice of translated Jaeger batches and error if translation failed.
func ProtoFromTraces(td ptrace.Traces) []*model.Batch { _ = "STUB: not implemented"; return nil }

func resourceSpansToJaegerProto(rs ptrace.ResourceSpans) *model.Batch {
	_ = "STUB: not implemented"
	return nil
}

// Approximate the number of the spans as the number of the spans in the first
// instrumentation library info.

func resourceToJaegerProtoProcess(resource pcommon.Resource) *model.Process {
	_ = "STUB: not implemented"
	return nil
}

func appendTagsFromResourceAttributes(dest []model.KeyValue, attrs pcommon.Map) []model.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func appendTagsFromAttributes(dest []model.KeyValue, attrs pcommon.Map) []model.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func attributeToJaegerProtoTag(key string, attr pcommon.Value) model.KeyValue {
	_ = "STUB: not implemented"
	return *new(model.KeyValue)
}

func spanToJaegerProto(span ptrace.Span, libraryTags pcommon.InstrumentationScope) *model.Span {
	_ = "STUB: not implemented"
	return nil
}

func getJaegerProtoSpanTags(span ptrace.Span, scope pcommon.InstrumentationScope) []model.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func traceIDToJaegerProto(traceID pcommon.TraceID) model.TraceID {
	_ = "STUB: not implemented"
	return *new(model.TraceID)
}

func spanIDToJaegerProto(spanID pcommon.SpanID) model.SpanID {
	_ = "STUB: not implemented"
	return *new(model.SpanID)
}

// makeJaegerProtoReferences constructs jaeger span references based on parent span ID and span links.
// The parent span ID is used to add a CHILD_OF reference, _unless_ it is referenced from one of the links.
func makeJaegerProtoReferences(links ptrace.SpanLinkSlice, parentSpanID model.SpanID, traceID model.TraceID) []model.SpanRef {
	_ = "STUB: not implemented"
	return nil
}

// Put parent span ID at the first place because usually backends look for it
// as the first CHILD_OF item in the model.SpanRef slice.

// We already added a reference to this span, but maybe with the wrong type, so override.

func spanEventsToJaegerProtoLogs(events ptrace.SpanEventSlice) []model.Log {
	_ = "STUB: not implemented"
	return nil
}

func getTagFromSpanKind(spanKind ptrace.SpanKind) (model.KeyValue, bool) {
	_ = "STUB: not implemented"
	return *new(model.KeyValue), false
}

func getTagFromStatusCode(statusCode ptrace.StatusCode) (model.KeyValue, bool) {
	_ = "STUB: not implemented"
	return *new(model.KeyValue), false
}

func getErrorTagFromStatusCode(statusCode ptrace.StatusCode) (model.KeyValue, bool) {
	_ = "STUB: not implemented"
	return *new(model.KeyValue), false
}

func getTagFromStatusMsg(statusMsg string) (model.KeyValue, bool) {
	_ = "STUB: not implemented"
	return *new(model.KeyValue), false
}

func getTagsFromTraceState(traceState string) ([]model.KeyValue, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// TODO Bring this inline with solution for jaegertracing/jaeger-client-java #702 once available

func getTagsFromInstrumentationLibrary(il pcommon.InstrumentationScope) ([]model.KeyValue, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func refTypeFromLink(link ptrace.SpanLink) model.SpanRefType {
	_ = "STUB: not implemented"
	return *new(model.SpanRefType)
}

func strToJRefType(attr string) model.SpanRefType {
	_ = "STUB: not implemented"
	return *new(model.SpanRefType)
}

// There are only 2 types of SpanRefType we assume that everything
// that's not a model.ChildOf is a model.FollowsFrom
