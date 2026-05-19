// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unmarshaler // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver/internal/unmarshaler"

import (
	jaegerproto "github.com/jaegertracing/jaeger-idl/model/v1"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	_ ptrace.Unmarshaler = JaegerProtoSpanUnmarshaler{}
	_ ptrace.Unmarshaler = JaegerJSONSpanUnmarshaler{}
)

type JaegerProtoSpanUnmarshaler struct{}

func (JaegerProtoSpanUnmarshaler) UnmarshalTraces(bytes []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

type JaegerJSONSpanUnmarshaler struct{}

func (JaegerJSONSpanUnmarshaler) UnmarshalTraces(data []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func jaegerSpanToTraces(span *jaegerproto.Span) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
