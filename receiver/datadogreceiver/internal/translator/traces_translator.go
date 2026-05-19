// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"bytes"
	"net/http"
	"sync"

	pb "github.com/DataDog/datadog-agent/pkg/proto/pbgo/trace"
	lru "github.com/hashicorp/golang-lru/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const (
	datadogSpanKindKey = "span.kind"
	// The datadog trace id
	//
	// Type: string
	// Requirement Level: Optional
	// Examples: '6249785623524942554'
	attributeDatadogTraceID = "datadog.trace.id"
	// The datadog span id
	//
	// Type: string
	// Requirement Level: Optional
	// Examples: '228114450199004348'
	attributeDatadogSpanID = "datadog.span.id"
)

var spanProcessor = map[string]func(*pb.Span, *ptrace.Span){
	// HTTP
	"servlet.request": processHTTPSpan,
	"http.request":    processHTTPSpan,
	"web.request":     processHTTPSpan,

	// Internal
	"spring.handler": processInternalSpan,

	// Database
	"postgresql.query": processDBSpan,
	"redis.query":      processDBSpan,

	// GRPC
	"grpc.server": processGRPCSpan,
	"grpc.client": processGRPCSpan,

	// AWS
	"aws.request": processAWSSdkSpan,
	"aws.command": processAWSSdkSpan,
}

func upsertHeadersAttributes(req *http.Request, attrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// traceID64to128 reconstructs the 128 bits TraceID, if available or cached.
//
// Datadog traces split a 128 bits trace id in two parts: TraceID and Tags._dd_p_tid. This happens if the
// instrumented service received a TraceContext from an OTel instrumented service. When it happens, we need
// to concatenate the two into newSpan.TraceID.
// The traceIDCache keeps track of the TraceIDs we process as only the first span has the upper 64 bits from the 128
// bits trace ID.
//
// Note: This may not be resilient to related spans being flushed separately in datadog's tracing libraries.
//
//	It might also not work if multiple datadog instrumented services are chained.
//
// This is currently gated by a feature gate (receiver.datadogreceiver.Enable128BitTraceID). If we don't get a cache
// in traceIDCache, we don't enable this behavior.
func traceID64to128(span *pb.Span, traceIDCache *lru.Cache[uint64, pcommon.TraceID]) (pcommon.TraceID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID), nil
}

// Child spans don't have _dd.p.tid, we cache it.

func processInternalSpan(span *pb.Span, newSpan *ptrace.Span) { _ = "STUB: not implemented"; return }

func processHTTPSpan(span *pb.Span, newSpan *ptrace.Span) {
	_ = "STUB: not implemented"
	// https://opentelemetry.io/docs/specs/semconv/http/http-spans/#name
	// We assume that http.route coming from datadog is low cardinality
	return
}

func processDBSpan(span *pb.Span, newSpan *ptrace.Span) {
	_ = "STUB: not implemented"
	// references:
	// https://github.com/DataDog/documentation/blob/master/content/en/tracing/guide/ignoring_apm_resources.md#database
	// https://opentelemetry.io/docs/specs/semconv/database/database-spans/#name
	return
}

func processGRPCSpan(span *pb.Span, newSpan *ptrace.Span) {
	_ = "STUB: not implemented"
	// references:
	// https://github.com/DataDog/documentation/blob/master/content/en/tracing/guide/ignoring_apm_resources.md#remote-procedure-calls
	// https://opentelemetry.io/docs/specs/semconv/rpc/rpc-spans/
	return
}

// ddSpan.Attributes["grpc.status.code"] contains the gRPC status code name (eg "OK")
// not the numeric value (eg "0")
// it's ddSpan.error that indicates holds the gRPC status code numeric value

// "rpc.method" is used by dd-trace-rb, check dd-trace-php

// "rpc.grpc.full_method" is used by dd-trace-go & dd-trace-rb, they also set span.Resource to the full method name
// format: /$package.$service/$method

// format: /$package.$service/$method
// "grpc.method.name" is used by dd-trace-dotnet dd-trace-python & dd-trace-go

// unexpected format

// resource is used by dd-trace-java

func processAWSSdkSpan(span *pb.Span, newSpan *ptrace.Span) {
	_ = "STUB: not implemented"
	// https://opentelemetry.io/docs/specs/semconv/cloud-providers/aws-sdk/
	return
}

func processSpanByName(span *pb.Span, newSpan *ptrace.Span) { _ = "STUB: not implemented"; return }

func traceChunkSamplingPriority(traceChunk *pb.TraceChunk) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func ToTraces(logger *zap.Logger, payload *pb.TracerPayload, req *http.Request, traceIDCache *lru.Cache[uint64, pcommon.TraceID]) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// Creating a map of service spans to slices
// since the expectation is that `service.name`
// is added as a resource attribute in most systems
// now instead of being a span level attribute.

// Restore base service name as the service name.
// Without this, internal spans such as postgresql queries have a service.name set to postgresql

// Preserve original per-span service name so the DD exporter
// can recover it via span-level service.name precedence

// Try to get the 128-bit traceID, if available.

// For client/producer/consumer spans, if we have `peer.hostname`, and `server.address` is unset, set
// `server.address` to `peer.hostname`.

// Some spans need specific processing (http, db, grpc...)

// DDSpanLink represents the structure of each JSON object
type DDSpanLink struct {
	TraceID    string         `json:"trace_id"`
	SpanID     string         `json:"span_id"`
	Tracestate string         `json:"tracestate"`
	Attributes map[string]any `json:"attributes"`
}

func tagsToSpanLinks(tags map[string]string, dest ptrace.SpanLinkSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert trace id.

// Convert span id.

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func PutBuffer(buffer *bytes.Buffer) { _ = "STUB: not implemented"; return }

func HandleTracesPayload(req *http.Request) (tp []*pb.TracerPayload, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeRequest(req *http.Request, dest *pb.Traces) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// do our best

func traceChunksFromSpans(spans []pb.Span) []*pb.TraceChunk { _ = "STUB: not implemented"; return nil }

func traceChunksFromTraces(traces pb.Traces) []*pb.TraceChunk {
	_ = "STUB: not implemented"
	return nil
}

func appVersionFromTraceChunks(traces []*pb.TraceChunk) string {
	_ = "STUB: not implemented"
	return ""
}

func getMediaType(req *http.Request) string { _ = "STUB: not implemented"; return "" }

func uInt64ToTraceID(high, low uint64) pcommon.TraceID {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID)
}

func uInt64ToSpanID(id uint64) pcommon.SpanID {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID)
}
