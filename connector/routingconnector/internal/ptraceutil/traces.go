// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ptraceutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector/internal/ptraceutil"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// MoveResourcesIf calls f sequentially for each ResourceSpans present in the first ptrace.Traces.
// If f returns true, the element is removed from the first ptrace.Traces and added to the second ptrace.Traces.
func MoveResourcesIf(from, to ptrace.Traces, f func(ptrace.ResourceSpans) bool) {
	_ = "STUB: not implemented"
	return
}

// CopyResourcesIf calls f sequentially for each ResourceSpans present in the first ptrace.Traces.
// If f returns true, the element is copied from the first ptrace.Traces to the second ptrace.Traces.
func CopyResourcesIf(from, to ptrace.Traces, f func(ptrace.ResourceSpans) bool) {
	_ = "STUB: not implemented"
	return
}

// MoveSpansWithContextIf calls f sequentially for each Span present in the first ptrace.Traces.
// If f returns true, the element is removed from the first ptrace.Traces and added to the second ptrace.Traces.
// Notably, the Resource and Scope associated with the Span are created in the second ptrace.Traces only once.
// Resources or Scopes are removed from the original if they become empty. All ordering is preserved.
func MoveSpansWithContextIf(from, to ptrace.Traces, f func(ptrace.ResourceSpans, ptrace.ScopeSpans, ptrace.Span) bool) {
	_ = "STUB: not implemented"
	return
}

// CopySpansWithContextIf calls f sequentially for each Span present in the first ptrace.Traces.
// If f returns true, the element is copied from the first ptrace.Traces to the second ptrace.Traces.
// Notably, the Resource and Scope associated with the Span are created in the second ptrace.Traces only once.
func CopySpansWithContextIf(from, to ptrace.Traces, f func(ptrace.ResourceSpans, ptrace.ScopeSpans, ptrace.Span) bool) {
	_ = "STUB: not implemented"
	return
}
