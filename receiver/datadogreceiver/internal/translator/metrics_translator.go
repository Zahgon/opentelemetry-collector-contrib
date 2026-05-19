// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
)

type MetricsTranslator struct {
	sync.RWMutex
	buildInfo         component.BuildInfo
	lastTs            map[identity.Stream]pcommon.Timestamp
	stringPool        *StringPool
	idleSeriesTimeout time.Duration
}

func NewMetricsTranslator(buildInfo component.BuildInfo, idleSeriesTimeout time.Duration) *MetricsTranslator {
	_ = "STUB: not implemented"
	return nil
}

func (mt *MetricsTranslator) streamHasTimestamp(stream identity.Stream) (pcommon.Timestamp, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp), false
}

func (mt *MetricsTranslator) updateLastTsForStream(stream identity.Stream, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// Prune recreates the map keeping only the recent items.
// It returns the number of removed items.
func (mt *MetricsTranslator) Prune() int {
	_ = "STUB: not implemented"
	// If the timeout is 0, the feature is disabled.
	// Return 0 immediately to preserve legacy behavior (keep all series).
	return 0
}

// Full Lock is required here because we are swapping the entire map reference.
// During this process, no one can read or write.

// Optimization: if the map is empty, do nothing.

// Create a new map.
// We let it grow organically to avoid allocating memory for the stale data.

// Convert pcommon timestamp (nanos) to time.Time.

// If the age is less than the max idle time, keep it.

// Swap the pointer. The old map (mt.lastTs) loses the reference.
// The Go Garbage Collector will detect this and release all memory allocated
// by the old buckets.

// Reset the string pool as well since we cleaned up the streams.

// Return the number of removed items.
