// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pebbletailstorageextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/tailstorage/pebbletailstorageextension"

import (
	"sync/atomic"

	"github.com/cockroachdb/pebble/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const (
	traceIDSeparator byte = ':'
	traceIDBytes          = len(pcommon.TraceID{})

	// storageVersion is a version to support evolution.
	storageVersion = "v0"
)

type storage struct {
	db          *pebble.DB
	logger      *zap.Logger
	nextSeq     atomic.Uint64
	unmarshaler ptrace.Unmarshaler
	marshaler   ptrace.Marshaler
}

func newStorage(storageDir string, logger *zap.Logger) (*storage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *storage) Close() error { _ = "STUB: not implemented"; return nil }

func (s *storage) Append(traceID pcommon.TraceID, rss ptrace.ResourceSpans) {
	_ = "STUB: not implemented"
	return
}

func (s *storage) Take(traceID pcommon.TraceID) (ptrace.Traces, bool) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), false
}

func (s *storage) Delete(traceID pcommon.TraceID) { _ = "STUB: not implemented"; return }

// Delete all entries for the trace in one range operation instead of
// iterating keys and deleting one-by-one.

func (s *storage) readByTracePrefix(prefix []byte) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}

// SeekPrefixGE enables prefix bloom filter usage when configured in Pebble options.

func tracePrefix(traceID pcommon.TraceID) (prefix [traceIDBytes + 1]byte) {
	_ = "STUB: not implemented"
	return nil
}

func tracePrefixUpperBound(prefix [traceIDBytes + 1]byte) (upper [traceIDBytes + 1]byte) {
	_ = "STUB: not implemented"
	// copy
	return nil
}

func traceEntryKey(traceID pcommon.TraceID, seq uint64) (key [traceIDBytes + 1 + 8]byte) {
	_ = "STUB: not implemented"
	return nil
}
