// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package goldendataset // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/goldendataset"

import (
	"io"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var statusCodeMap = map[PICTInputStatus]ptrace.StatusCode{
	SpanStatusUnset: ptrace.StatusCodeUnset,
	SpanStatusOk:    ptrace.StatusCodeOk,
	SpanStatusError: ptrace.StatusCodeError,
}

var statusMsgMap = map[PICTInputStatus]string{
	SpanStatusUnset: "Unset",
	SpanStatusOk:    "Ok",
	SpanStatusError: "Error",
}

// appendSpans appends to the ptrace.SpanSlice objects the number of spans specified by the count input
// parameter. The random parameter injects the random number generator to use in generating IDs and other random values.
// Using a random number generator with the same seed value enables reproducible tests.
//
// If err is not nil, the spans slice will have nil values.
func appendSpans(count int, pictFile string, random io.Reader, spanList ptrace.SpanSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// use existing if available

func generateSpanName(spanInputs *PICTSpanInputs) string { _ = "STUB: not implemented"; return "" }

// fillSpan generates a single ptrace.Span based on the input values provided. They are:
//
//	traceID - the trace ID to use, should not be nil
//	parentID - the parent span ID or nil if it is a root span
//	spanName - the span name, should not be blank
//	spanInputs - the pairwise combination of field value variations for this span
//	random - the random number generator to use in generating ID values
//
// The generated span is returned.
func fillSpan(traceID pcommon.TraceID, parentID pcommon.SpanID, spanName string, spanInputs *PICTSpanInputs, random io.Reader, span ptrace.Span) {
	_ = "STUB: not implemented"
	return
}

func generateTraceState(tracestate PICTInputTracestate) string {
	_ = "STUB: not implemented"
	return ""
}

func lookupSpanKind(kind PICTInputKind) ptrace.SpanKind {
	_ = "STUB: not implemented"
	return *new(ptrace.SpanKind)
}

func appendSpanAttributes(spanTypeID PICTInputAttributes, statusStr PICTInputStatus, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func fillStatus(statusStr PICTInputStatus, spanStatus ptrace.Status) {
	_ = "STUB: not implemented"
	return
}

func appendDatabaseSQLAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendDatabaseNoSQLAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendFaaSDatasourceAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendFaaSHTTPAttributes(includeStatus bool, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func appendFaaSPubSubAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendFaaSTimerAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendFaaSOtherAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendHTTPClientAttributes(includeStatus bool, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func appendHTTPServerAttributes(includeStatus bool, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func appendMessagingProducerAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendMessagingConsumerAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendGRPCClientAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendGRPCServerAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendInternalAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendMaxCountAttributes(includeStatus bool, attrMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func appendSpanEvents(eventCnt PICTInputSpanChild, spanEvents ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	return
}

func appendSpanLinks(linkCnt PICTInputSpanChild, random io.Reader, spanLinks ptrace.SpanLinkSlice) {
	_ = "STUB: not implemented"
	return
}

func calculateListSize(listCnt PICTInputSpanChild) int { _ = "STUB: not implemented"; return 0 }

func appendSpanEvent(index int, spanEvents ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	return
}

func appendSpanLink(random io.Reader, index int, spanLinks ptrace.SpanLinkSlice) {
	_ = "STUB: not implemented"
	return
}

func generateTraceID(random io.Reader) pcommon.TraceID {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID)
}

func generateSpanID(random io.Reader) pcommon.SpanID {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID)
}
