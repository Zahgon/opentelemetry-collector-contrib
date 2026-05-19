// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinv2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv2"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type jsonUnmarshaler struct {
	toTranslator ToTranslator
}

// UnmarshalTraces from JSON bytes.
func (j jsonUnmarshaler) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// NewJSONTracesUnmarshaler returns an unmarshaler for JSON bytes.
func NewJSONTracesUnmarshaler(parseStringTags bool) ptrace.Unmarshaler {
	_ = "STUB: not implemented"
	return *new(ptrace.Unmarshaler)
}

// NewJSONTracesMarshaler returns a marshaler to JSON bytes.
func NewJSONTracesMarshaler() ptrace.Marshaler {
	_ = "STUB: not implemented"
	return *new(ptrace.Marshaler)
}
