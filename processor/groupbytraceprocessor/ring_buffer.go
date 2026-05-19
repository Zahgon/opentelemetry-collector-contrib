// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package groupbytraceprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"

import "go.opentelemetry.io/collector/pdata/pcommon"

// ringBuffer keeps an in-memory bounded buffer with the in-flight trace IDs
type ringBuffer struct {
	index     int
	size      int
	ids       []pcommon.TraceID
	idToIndex map[pcommon.TraceID]int // key is traceID, value is the index on the 'ids' slice
}

func newRingBuffer(size int) *ringBuffer { _ = "STUB: not implemented"; return nil }

// the first span to be received will be placed at position '0'

func (r *ringBuffer) put(traceID pcommon.TraceID) pcommon.TraceID {
	_ = "STUB: not implemented"
	// calculates the item in the ring that we'll store the trace
	return *new(pcommon.TraceID)
}

// see if the ring has an item already

// clear space for the new item

// place the traceID in memory

func (r *ringBuffer) contains(traceID pcommon.TraceID) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *ringBuffer) delete(traceID pcommon.TraceID) bool { _ = "STUB: not implemented"; return false }
