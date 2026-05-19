// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package xidutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/core/xidutils"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// UInt64ToTraceID converts the pair of uint64 representation of a TraceID to pcommon.TraceID.
func UInt64ToTraceID(high, low uint64) pcommon.TraceID {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID)
}

// TraceIDToUInt64Pair converts the pcommon.TraceID to a pair of uint64 representation.
func TraceIDToUInt64Pair(traceID pcommon.TraceID) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// UInt64ToSpanID converts the uint64 representation of a SpanID to pcommon.SpanID.
func UInt64ToSpanID(id uint64) pcommon.SpanID {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID)
}

// SpanIDToUInt64 converts the pcommon.SpanID to uint64 representation.
func SpanIDToUInt64(spanID pcommon.SpanID) uint64 { _ = "STUB: not implemented"; return 0 }
