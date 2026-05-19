// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package parser // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/parser"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/libhoneyevent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/response"
)

// IndexMapping tracks which original libhoney event indices became logs vs traces
type IndexMapping struct {
	LogIndices   []int // Original event indices that became logs
	TraceIndices []int // Original event indices that became traces/spans/span events
}

// GetDatasetFromRequest extracts the dataset name from the request path
func GetDatasetFromRequest(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ToPdata converts a list of LibhoneyEvents to a Pdata Logs object and tracks which original indices became what
func ToPdata(dataset string, lhes []libhoneyevent.LibhoneyEvent, cfg libhoneyevent.FieldMapConfig, logger zap.Logger) (plog.Logs, ptrace.Traces, IndexMapping, []response.ResponseInBatch) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), *new(ptrace.Traces), *new(IndexMapping), nil
}

// Initialize index mapping to track which original events become logs vs traces

// Initialize parsing results to track success/failure per original event

// a list of already seen scopes

// adds a new found scope if needed

// Track successful span for consumer processing

// adds a new found scope if needed

// Track successful log for consumer processing

// Span events are processed later, so we need index mapping for them

// Span links are processed later, so we need index mapping for them

func addSpanEventsToSpan(sp ptrace.Span, events []libhoneyevent.LibhoneyEvent, alreadyUsedFields []string, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Handle cases where MsgPackTimestamp might be nil (e.g., JSON data from Refinery)

// Use current time if timestamp is not available

func addSpanLinksToSpan(sp ptrace.Span, links []libhoneyevent.LibhoneyEvent, alreadyUsedFields []string, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Convert slice to [16]byte array

// Convert slice to [8]byte array
