// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package marshaler // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/marshaler"

import (
	jaegerproto "github.com/jaegertracing/jaeger-idl/model/v1"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	_ TracesMarshaler = JaegerProtoSpanMarshaler{}
	_ TracesMarshaler = JaegerJSONSpanMarshaler{}
)

type JaegerProtoSpanMarshaler struct{}

type JaegerJSONSpanMarshaler struct{}

func (JaegerProtoSpanMarshaler) MarshalTraces(traces ptrace.Traces, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

func (JaegerJSONSpanMarshaler) MarshalTraces(traces ptrace.Traces, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

func marshalJaeger(traces ptrace.Traces, yield func(key, value []byte), marshal marshalJaegerSpanFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// continue to process spans that can be serialized

type marshalJaegerSpanFunc func(*jaegerproto.Span) ([]byte, error)

func marshalJaegerSpanProto(span *jaegerproto.Span) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalJaegerSpanJSON(span *jaegerproto.Span) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
