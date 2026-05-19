// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"

import (
	jaegerproto "github.com/jaegertracing/jaeger-idl/model/v1"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// copy from kafka receiver
type jaegerProtoSpanUnmarshaler struct{}

var _ TracesUnmarshaler = (*jaegerProtoSpanUnmarshaler)(nil)

func (jaegerProtoSpanUnmarshaler) Unmarshal(bytes []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (jaegerProtoSpanUnmarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

type jaegerJSONSpanUnmarshaler struct{}

var _ TracesUnmarshaler = (*jaegerJSONSpanUnmarshaler)(nil)

func (jaegerJSONSpanUnmarshaler) Unmarshal(data []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (jaegerJSONSpanUnmarshaler) Encoding() string { _ = "STUB: not implemented"; return "" }

func jaegerSpanToTraces(span *jaegerproto.Span) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
