// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaegerencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/jaegerencodingextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var (
	_ encoding.TracesUnmarshalerExtension = &jaegerExtension{}
	_ ptrace.Unmarshaler                  = &jaegerExtension{}
)

type jaegerExtension struct {
	config      *Config
	unmarshaler ptrace.Unmarshaler
}

func (e *jaegerExtension) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (e *jaegerExtension) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*jaegerExtension) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
