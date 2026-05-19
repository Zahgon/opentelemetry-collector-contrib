// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinv1 // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv1"

import (
	"errors"

	"github.com/jaegertracing/jaeger-idl/thrift-gen/zipkincore"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type thriftUnmarshaler struct{}

// UnmarshalTraces from Thrift bytes.
func (thriftUnmarshaler) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// NewThriftTracesUnmarshaler returns an unmarshaler for Zipkin Thrift.
func NewThriftTracesUnmarshaler() ptrace.Unmarshaler {
	_ = "STUB: not implemented"
	return *new(ptrace.Unmarshaler)
}

// thriftBatchToTraces converts Zipkin v1 spans to ptrace.Traces.
func thriftBatchToTraces(zSpans []*zipkincore.Span) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func thriftToSpanAndEndpoint(zSpan *zipkincore.Span) spanAndEndpoint {
	_ = "STUB: not implemented"
	return *new(spanAndEndpoint)
}

// TODO: (@pjanotti) ideally we should error here instead of generating invalid Traces
// however per https://go.opentelemetry.io/collector/issues/349
// failures on the receivers in general are silent at this moment, so letting them
// proceed for now. We should validate the traceID, spanID and parentID are good with
// OTLP requirements.

func thriftAnnotationsToSpanAndEndpoint(ztAnnotations []*zipkincore.Annotation) (ptrace.Span, *endpoint) {
	_ = "STUB: not implemented"
	return *new(ptrace.Span), nil
}

func toTranslatorEndpoint(e *zipkincore.Endpoint) *endpoint { _ = "STUB: not implemented"; return nil }

var trueByteSlice = []byte{1}

func thriftBinAnnotationsToSpanAttributes(span ptrace.Span, ztBinAnnotations []*zipkincore.BinaryAnnotation) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO: (@pjanotti) add reference to OpenTracing and change related tags to use them

var errNotEnoughBytes = errors.New("not enough bytes representing the number")

func bytesInt16ToInt64(b []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func bytesInt32ToInt64(b []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func bytesInt64ToInt64(b []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func bytesFloat64ToFloat64(b []byte) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func strAttributeForError(dest pcommon.Value, err error) { _ = "STUB: not implemented"; return }
