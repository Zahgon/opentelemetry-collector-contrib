// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalkingencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/skywalkingencodingextension"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type skywalkingProtobufTrace struct{}

func (skywalkingProtobufTrace) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}
