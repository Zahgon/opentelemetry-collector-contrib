// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/zipkinencodingextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

const (
	zipkinProtobufEncoding = "zipkin_proto"
	zipkinJSONEncoding     = "zipkin_json"
	zipkinThriftEncoding   = "zipkin_thrift"
	v1                     = "v1"
	v2                     = "v2"
)

var (
	_ encoding.TracesMarshalerExtension   = (*zipkinExtension)(nil)
	_ encoding.TracesUnmarshalerExtension = (*zipkinExtension)(nil)
)

type zipkinExtension struct {
	config      *Config
	marshaler   ptrace.Marshaler
	unmarshaler ptrace.Unmarshaler
}

func newExtension(config *Config) (*zipkinExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*zipkinExtension) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*zipkinExtension) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (ex *zipkinExtension) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (ex *zipkinExtension) MarshalTraces(td ptrace.Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
