// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tailstorageextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/tailstorageextension"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var _ TailStorage = (*inMemoryTailStorage)(nil)

type inMemoryTailStorage struct {
	idToSpans map[pcommon.TraceID]ptrace.Traces
}

func NewInMemoryTailStorage() TailStorage { _ = "STUB: not implemented"; return *new(TailStorage) }

func (s *inMemoryTailStorage) Append(traceID pcommon.TraceID, rss ptrace.ResourceSpans) {
	_ = "STUB: not implemented"
	return
}

func (s *inMemoryTailStorage) Take(traceID pcommon.TraceID) (ptrace.Traces, bool) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), false
}

func (s *inMemoryTailStorage) Delete(traceID pcommon.TraceID) { _ = "STUB: not implemented"; return }
