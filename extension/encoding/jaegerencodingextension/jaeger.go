// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaegerencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/jaegerencodingextension"

import (
	jaegerproto "github.com/jaegertracing/jaeger-idl/model/v1"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type jaegerProtobufTrace struct{}

func (jaegerProtobufTrace) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

type jaegerJSONTrace struct{}

func (jaegerJSONTrace) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func jaegerSpanToTraces(span *jaegerproto.Span) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
