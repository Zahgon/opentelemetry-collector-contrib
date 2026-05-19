// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger"

import (
	"github.com/jaegertracing/jaeger-idl/model/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var blankJaegerProtoSpan = new(model.Span)

// ProtoToTraces converts multiple Jaeger proto batches to internal traces
func ProtoToTraces(batches []*model.Batch) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func regroup(batches []*model.Batch) []*model.Batch {
	_ = "STUB: not implemented"
	// Re-group batches
	// This is needed as there might be a Process within Batch and Span at the same
	// time, with the span one taking precedence.
	// As we only have it at one level in OpenTelemetry, ResourceSpans, we split
	// each batch into potentially multiple other batches, with the sum of their
	// processes as the key to a map.
	// Step 1) iterate over the batches
	// Step 2) for each batch, calculate the batch's process checksum and store
	// it on a map, with the checksum as the key and the process as the value
	// Step 3) iterate the spans for a batch: if a given span has its own process,
	// calculate the checksum for the process and store it on the same map
	// Step 4) each entry on the map becomes a ResourceSpan
	return nil
}

func batchForProcess(registry map[uint64]*model.Batch, p *model.Process) *model.Batch {
	_ = "STUB: not implemented"
	return nil
}

func checksum(process *model.Process) uint64 {
	_ = "STUB: not implemented"
	// this will get all the keys and values, plus service name, into this buffer
	// this is potentially dangerous, as a batch/span with a big enough processes
	// might cause the collector to allocate this extra big information
	// for this reason, we hash it as an integer and return it, instead of keeping
	// all the hashes for all the processes for all batches in memory
	return 0
}

// this effectively means that all spans from batches with nil processes
// will be grouped together
// this should only ever happen in unit tests
// this implementation never returns an error according to the Hash interface

func protoBatchToResourceSpans(batch model.Batch, dest ptrace.ResourceSpans) {
	_ = "STUB: not implemented"
	return
}

func jProcessToInternalResource(process *model.Process, dest pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

// Handle special keys translations.

// translateHostnameAttr translates "hostname" atttribute
func translateHostnameAttr(attrs pcommon.Map) { _ = "STUB: not implemented"; return }

// translateHostnameAttr translates "jaeger.version" atttribute
func translateJaegerVersionAttr(attrs pcommon.Map) { _ = "STUB: not implemented"; return }

type scope struct {
	name, version string
}

func jSpansToInternal(spans []*model.Span, dest ptrace.ScopeSpansSlice) {
	_ = "STUB: not implemented"
	return
}

func jSpanToInternal(span *model.Span, dest ptrace.Span) { _ = "STUB: not implemented"; return }

// drop the attributes slice if all of them were replaced during translation

func jTagsToInternalAttributes(tags []model.KeyValue, dest pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func setInternalSpanStatus(attrs pcommon.Map, span ptrace.Span) { _ = "STUB: not implemented"; return }

// The error tag is the ultimate truth for a Jaeger spans' error
// status. Only parse the otel.status_code tag if the error tag is
// not set to true.

// Regardless of error tag value, remove the otel.status_code tag. The
// otel.status_message tag will have already been removed if
// statusExists is true.

// Fallback to introspecting if this span represents a failed HTTP
// request or response, but again, only do so if the `error` tag was
// not set to true and no explicit status was sent.

// extractStatusDescFromAttr returns the OTel status description from attrs
// along with true if it is set. Otherwise, an empty string and false are
// returned. The OTel status description attribute is deleted from attrs in
// the process.
func extractStatusDescFromAttr(attrs pcommon.Map) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// codeFromAttr returns the integer code value from attrVal. An error is
// returned if the code is not represented by an integer or string value in
// the attrVal or the value is outside the bounds of an int representation.
func codeFromAttr(attrVal pcommon.Value) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func getStatusCodeFromHTTPStatusAttr(attrVal pcommon.Value, kind ptrace.SpanKind) (ptrace.StatusCode, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.StatusCode), nil
}

// For HTTP status codes in the 4xx range span status MUST be left unset
// in case of SpanKind.SERVER and MUST be set to Error in case of SpanKind.CLIENT.
// For HTTP status codes in the 5xx range, as well as any other code the client
// failed to interpret, span status MUST be set to Error.

func jSpanKindToInternal(spanKind string) ptrace.SpanKind {
	_ = "STUB: not implemented"
	return *new(ptrace.SpanKind)
}

func jLogsToSpanEvents(logs []model.Log, dest ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	return
}

// jReferencesToSpanLinks sets internal span links based on jaeger span references skipping excludeParentID
func jReferencesToSpanLinks(refs []model.SpanRef, excludeParentID model.SpanID, dest ptrace.SpanLinkSlice) {
	_ = "STUB: not implemented"
	return
}

func getTraceStateFromAttrs(attrs pcommon.Map) string {
	_ = "STUB: not implemented"

	// TODO Bring this inline with solution for jaegertracing/jaeger-client-java #702 once available
	return ""
}

func getScope(span *model.Span) scope { _ = "STUB: not implemented"; return *new(scope) }

func getAndDeleteTag(span *model.Span, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func jRefTypeToAttribute(ref model.SpanRefType) string { _ = "STUB: not implemented"; return "" }
