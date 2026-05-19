// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinv2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv2"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type protobufUnmarshaler struct {
	// debugWasSet toggles the Debug field of each Span. It is usually set to true if
	// the "X-B3-Flags" header is set to 1 on the request.
	debugWasSet bool

	toTranslator ToTranslator
}

// UnmarshalTraces from protobuf bytes.
func (p protobufUnmarshaler) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// NewProtobufTracesUnmarshaler returns an ptrace.Unmarshaler of protobuf bytes.
func NewProtobufTracesUnmarshaler(debugWasSet, parseStringTags bool) ptrace.Unmarshaler {
	_ = "STUB: not implemented"
	return *new(ptrace.Unmarshaler)
}

// NewProtobufTracesMarshaler returns a new ptrace.Marshaler to protobuf bytes.
func NewProtobufTracesMarshaler() ptrace.Marshaler {
	_ = "STUB: not implemented"
	return *new(ptrace.Marshaler)
}
