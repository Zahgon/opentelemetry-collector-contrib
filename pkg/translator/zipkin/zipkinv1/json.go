// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinv1 // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv1"

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	// ZipkinV1 friendly conversion errors
	msgZipkinV1JSONUnmarshalError = "zipkinv1"
	msgZipkinV1TraceIDError       = "zipkinV1 span traceId"
	msgZipkinV1SpanIDError        = "zipkinV1 span id"
	msgZipkinV1ParentIDError      = "zipkinV1 span parentId"
	// Generic hex to ID conversion errors
	errHexTraceIDWrongLen = errors.New("hex traceId span has wrong length (expected 16 or 32)")
	errHexTraceIDZero     = errors.New("traceId is zero")
	errHexIDWrongLen      = errors.New("hex Id has wrong length (expected 16)")
	errHexIDZero          = errors.New("ID is zero")
)

type jsonUnmarshaler struct {
	// ParseStringTags should be set to true if tags should be converted to numbers when possible.
	ParseStringTags bool
}

// UnmarshalTraces from JSON bytes.
func (j jsonUnmarshaler) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// NewJSONTracesUnmarshaler returns an unmarshaler for Zipkin JSON.
func NewJSONTracesUnmarshaler(parseStringTags bool) ptrace.Unmarshaler {
	_ = "STUB: not implemented"
	return *new(ptrace.Unmarshaler)
}

// Trace translation from Zipkin V1 is a bit of special case since there is no model
// defined in golang for Zipkin V1 spans and there is no need to define one here, given
// that the jsonSpan defined below is as defined at:
// https://zipkin.io/zipkin-api/zipkin-api.yaml
type jsonSpan struct {
	TraceID           string              `json:"traceId"`
	Name              string              `json:"name,omitempty"`
	ParentID          string              `json:"parentId,omitempty"`
	ID                string              `json:"id"`
	Timestamp         int64               `json:"timestamp"`
	Duration          int64               `json:"duration"`
	Debug             bool                `json:"debug,omitempty"`
	Annotations       []*annotation       `json:"annotations,omitempty"`
	BinaryAnnotations []*binaryAnnotation `json:"binaryAnnotations,omitempty"`
}

// endpoint structure used by jsonSpan.
type endpoint struct {
	ServiceName string `json:"serviceName"`
	IPv4        string `json:"ipv4"`
	IPv6        string `json:"ipv6"`
	Port        int32  `json:"port"`
}

// annotation struct used by jsonSpan.
type annotation struct {
	Timestamp int64     `json:"timestamp"`
	Value     string    `json:"value"`
	Endpoint  *endpoint `json:"endpoint"`
}

// binaryAnnotation used by jsonSpan.
type binaryAnnotation struct {
	Key      string    `json:"key"`
	Value    string    `json:"value"`
	Endpoint *endpoint `json:"endpoint"`
}

// jsonBatchToTraces converts a JSON blob with a list of Zipkin v1 spans to ptrace.Traces.
func jsonBatchToTraces(blob []byte, parseStringTags bool) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// error from internal package function, it already wraps the error to give better context.

type spanAndEndpoint struct {
	span     ptrace.Span
	endpoint *endpoint
}

func zipkinToTraces(spanAndEndpoints []spanAndEndpoint) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *

	// Service to batch maps the service name to the trace request with the corresponding node.
	new(ptrace.Traces), nil
}

func jsonToSpanAndEndpoint(zSpan *jsonSpan, parseStringTags bool) (spanAndEndpoint, error) {
	_ = "STUB: not implemented"
	return *new(spanAndEndpoint), nil
}

func jsonBinAnnotationsToSpanAttributes(span ptrace.Span, binAnnotations []*binaryAnnotation, parseStringTags bool) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO: (@pjanotti) add reference to OpenTracing and change related tags to use them

func parseAnnotationValue(value string, parseStringTags bool) pcommon.Value {
	_ = "STUB: not implemented"
	return *new(pcommon.Value)
}

// Unknown service name works both as a default value and a flag to indicate that a valid endpoint was found.
const unknownServiceName = "unknown-service"

func jsonAnnotationsToSpanAndEndpoint(annotations []*annotation) (ptrace.Span, *endpoint) {
	_ = "STUB: not implemented"
	// Zipkin V1 annotations have a timestamp so they fit well with ptrace.SpanEvent
	return *new(ptrace.Span), nil
}

// We want to set the span kind from the first annotation that contains information
// about the span kind. This flags ensures we only set span kind once from
// the first annotation.

// Check if annotation has span kind information.

// Populate the endpoint if it is not already populated and current endpoint
// has a service name and span kind.

// We have not yet populated span kind, do it now.
// Translate from Zipkin span kind stored in Value field to Kind/ExternalKind
// pair of internal fields.

// Remember that we populated the span kind, so that we don't do it again.

// If this annotation is for the send/receive timestamps, no need to create the annotation

func hexToTraceID(hexStr string) (pcommon.TraceID, error) {
	_ = "STUB: not implemented"
	// Per info at https://zipkin.io/zipkin-api/zipkin-api.yaml it should be 16 or 32 characters
	return *new(pcommon.TraceID), nil
}

func hexToSpanID(hexStr string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	// Per info at https://zipkin.io/zipkin-api/zipkin-api.yaml it should be 16 characters
	return *new(pcommon.SpanID), nil
}

func epochMicrosecondsToTimestamp(msecs int64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func getOrCreateNodeRequest(m map[string]ptrace.SpanSlice, td ptrace.Traces, endpoint *endpoint) ptrace.SpanSlice {
	_ = "STUB: not implemented"
	// this private function assumes that the caller never passes an nil endpoint
	return *new(ptrace.SpanSlice)
}

func (ep *endpoint) string() string { _ = "STUB: not implemented"; return "" }

func (ep *endpoint) setAttributes(dest pcommon.Map) { _ = "STUB: not implemented"; return }

func setTimestampsIfUnset(span ptrace.Span) {
	_ = "STUB: not implemented"
	// zipkin allows timestamp to be unset, but opentelemetry-collector expects it to have a value.
	// If this is unset, the conversion from open census to the internal trace format breaks
	// what should be an identity transformation oc -> internal -> oc
	return
}
