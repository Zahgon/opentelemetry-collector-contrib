// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logdedupprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/logdedupprocessor"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/logdedupprocessor/internal/metadata"
)

// Attributes names for first and last observed timestamps
const (
	firstObservedTSAttr = "first_observed_timestamp"
	lastObservedTSAttr  = "last_observed_timestamp"
)

// timeNow can be reassigned for testing
var timeNow = time.Now

// logAggregator tracks the number of times a specific logRecord has been seen.
type logAggregator struct {
	resources         map[uint64]*resourceAggregator
	logCountAttribute string
	timezone          *time.Location
	telemetryBuilder  *metadata.TelemetryBuilder
	dedupFields       []string
}

// newLogAggregator creates a new LogCounter.
func newLogAggregator(logCountAttribute string, timezone *time.Location, telemetryBuilder *metadata.TelemetryBuilder, dedupFields []string) *logAggregator {
	_ = "STUB: not implemented"
	return nil
}

// Export exports the counter as a Logs
func (l *logAggregator) Export(ctx context.Context) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// Record aggregated logs records

// Set log record timestamps

// Add attributes for log count and first/last observed timestamps

// Add adds the logRecord to the resource aggregator that is identified by the resource attributes
func (l *logAggregator) Add(resource pcommon.Resource, scope pcommon.InstrumentationScope, logRecord plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

// Reset resets the counter.
func (l *logAggregator) Reset() { _ = "STUB: not implemented"; return }

// resourceAggregator dimensions the counter by resource.
type resourceAggregator struct {
	resource      pcommon.Resource
	scopeCounters map[uint64]*scopeAggregator
	dedupFields   []string
}

// newResourceAggregator creates a new ResourceCounter.
func newResourceAggregator(resource pcommon.Resource, dedupFields []string) *resourceAggregator {
	_ = "STUB: not implemented"
	return nil
}

// Add increments the counter that the logRecord matches.
func (r *resourceAggregator) Add(scope pcommon.InstrumentationScope, logRecord plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

// scopeAggregator dimensions the counter by scope.
type scopeAggregator struct {
	scope       pcommon.InstrumentationScope
	logCounters map[uint64]*logCounter
	dedupFields []string
}

// newScopeAggregator creates a new ScopeCounter.
func newScopeAggregator(scope pcommon.InstrumentationScope, dedupFields []string) *scopeAggregator {
	_ = "STUB: not implemented"
	return nil
}

// Add increments the counter that the logRecord matches.
func (s *scopeAggregator) Add(logRecord plog.LogRecord) { _ = "STUB: not implemented"; return }

// logCounter is a counter for a log record.
type logCounter struct {
	logRecord              plog.LogRecord
	firstObservedTimestamp time.Time
	lastObservedTimestamp  time.Time
	count                  int64
}

// newLogCounter creates a new AttributeCounter.
func newLogCounter(logRecord plog.LogRecord) *logCounter {
	_ = "STUB: not implemented"
	// Since we always remove the logRecord if we got to this point, we can move it instead of copying.
	return nil
}

// Increment increments the counter.
func (a *logCounter) Increment() { _ = "STUB: not implemented"; return }

// getResourceKey creates a unique hash for the resource to use as a map key
func getResourceKey(resource pcommon.Resource) uint64 { _ = "STUB: not implemented"; return 0 }

// getScopeKey creates a unique hash for the scope to use as a map key
func getScopeKey(scope pcommon.InstrumentationScope) uint64 { _ = "STUB: not implemented"; return 0 }

// getLogKey creates a unique hash for the log record to use as a map key.
// If dedupFields is non-empty, it is used to determine the fields whose values are hashed.
// If no dedupFields are found in the log record, all fields are hashed.
func getLogKey(logRecord plog.LogRecord, dedupFields []string) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func getMap(logRecord plog.LogRecord, leadingPart string) (pcommon.Map, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.Map), false
}

func getKeyValue(valueMap pcommon.Map, keyParts []string) (pcommon.Value, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), false
}

// Look for the value associated with the next key part.
// If we don't find it then return

// No more key parts that means we have found the value

// If the value is a map then recurse through with the remaining parts
