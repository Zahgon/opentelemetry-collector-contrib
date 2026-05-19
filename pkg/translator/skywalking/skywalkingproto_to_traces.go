// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalking // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/skywalking"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"
	common "skywalking.apache.org/repo/goapi/collect/common/v3"
	agentV3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const (
	AttributeRefType                   = "refType"
	AttributeParentService             = "parent.service"
	AttributeParentInstance            = "parent.service.instance"
	AttributeParentEndpoint            = "parent.endpoint"
	AttributeSkywalkingSpanID          = "sw8.span_id"
	AttributeSkywalkingTraceID         = "sw8.trace_id"
	AttributeSkywalkingSegmentID       = "sw8.segment_id"
	AttributeSkywalkingParentSpanID    = "sw8.parent_span_id"
	AttributeSkywalkingParentSegmentID = "sw8.parent_segment_id"
	AttributeNetworkAddressUsedAtPeer  = "network.AddressUsedAtPeer"
)

var otSpanTagsMappingStable = map[string]string{
	"url":         string(conventions.URLFullKey),
	"status_code": string(conventions.HTTPResponseStatusCodeKey),
	"db.type":     string(conventions.DBSystemNameKey),
	"db.instance": string(conventions.DBNamespaceKey),
	"mq.broker":   string(conventions.ServerAddressKey),
}

var otSpanTagsMappingLegacy = map[string]string{
	"url":         "http.url",
	"status_code": "http.status_code",
	"db.type":     "db.system",
	"db.instance": "db.name",
	"mq.broker":   "net.peer.name",
}

// getOtSpanTagsMapping returns the appropriate mapping based on the feature gate
func getOtSpanTagsMapping() map[string]string { _ = "STUB: not implemented"; return nil }

// ProtoToTraces converts multiple skywalking proto batches to internal traces
func ProtoToTraces(segment *agentV3.SegmentObject) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}

func swTagsToInternalResource(span *agentV3.SpanObject, dest pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

func swSpansToSpanSlice(traceID, segmentID string, spans []*agentV3.SpanObject, dest ptrace.SpanSlice) {
	_ = "STUB: not implemented"
	return
}

func swSpanToSpan(traceID, segmentID string, span *agentV3.SpanObject, dest ptrace.Span) {
	_ = "STUB: not implemented"
	return
}

// skywalking defines segmentId + spanId as unique identifier
// so use segmentId to convert to an unique otel-span

// parent spanid = -1, means(root span) no parent span in current skywalking segment, so it is necessary to search for the parent segment.

// TODO: SegmentReference references usually have only one element, but in batch consumer case, such as in MQ or async batch process, it could be multiple.
// We only handle one element for now.

// drop the attributes slice if all of them were replaced during translation

// skywalking: In the across thread and across processes, these references target the parent segments.

func swReferencesToSpanLinks(refs []*agentV3.SegmentReference, dest ptrace.SpanLinkSlice) {
	_ = "STUB: not implemented"
	return
}

func setInternalSpanStatus(span *agentV3.SpanObject, dest ptrace.Status) {
	_ = "STUB: not implemented"
	return
}

func setSwSpanIDToAttributes(span *agentV3.SpanObject, dest pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func swLogsToSpanEvents(logs []*agentV3.Log, dest ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	return
}

func swKvPairsToInternalAttributes(pairs []*common.KeyStringValuePair, dest pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// microsecondsToTimestamp converts epoch microseconds to pcommon.Timestamp
func microsecondsToTimestamp(ms int64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func swTraceIDToTraceID(traceID string) pcommon.TraceID {
	_ = "STUB: not implemented"
	// skywalking traceid format:
	// de5980b8-fce3-4a37-aab9-b4ac3af7eedd: from browser/js-sdk/envoy/nginx-lua sdk/py-agent
	// 56a5e1c519ae4c76a2b8b11d92cead7f.12.16563474296430001: from java-agent
	return *new(pcommon.TraceID)
}

// 36: uuid length (rfc4122)

func segmentIDToSpanID(segmentID string, spanID uint32) pcommon.SpanID {
	_ = "STUB: not implemented"
	// skywalking segmentid format:
	// 56a5e1c519ae4c76a2b8b11d92cead7f.12.16563474296430001: from TraceSegmentId
	// 56a5e1c519ae4c76a2b8b11d92cead7f: from ParentTraceSegmentId
	return *new(pcommon.SpanID)
}

func swStringToUUID(s string, extra uint32) (dst [16]byte) {
	_ = "STUB: not implemented"
	// there are 2 possible formats for 's':
	// s format = 56a5e1c519ae4c76a2b8b11d92cead7f.0000000000.000000000000000000
	//
	//	^ start(length=32)               ^ mid(u32) ^ last(u64)
	//
	// uid = UUID(start) XOR ([4]byte(extra) . [4]byte(uint32(mid)) . [8]byte(uint64(last)))
	return nil
}

// s format = 56a5e1c519ae4c76a2b8b11d92cead7f
//            ^ start(length=32)
// uid = UUID(start) XOR [4]byte(extra)

func uuidTo8Bytes(uuid [16]byte) [8]byte {
	_ = "STUB: not implemented"
	// high bit XOR low bit
	return nil
}

func unsafeGetBytes(s string) []byte { _ = "STUB: not implemented"; return nil }
