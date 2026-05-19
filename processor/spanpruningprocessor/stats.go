// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package spanpruningprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanpruningprocessor"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// aggregationData tracks statistics and time ranges for a group of spans in
// a single pass, replacing separate calculations for efficiency.
type aggregationData struct {
	count         int64
	minDuration   time.Duration
	maxDuration   time.Duration
	sumDuration   time.Duration
	bucketCounts  []int64
	earliestStart pcommon.Timestamp
	latestEnd     pcommon.Timestamp
}

// calculateAggregationData derives span counts and duration stats for the
// provided nodes in one traversal.
func (p *spanPruningProcessor) calculateAggregationData(nodes []*spanNode) aggregationData {
	_ = "STUB: not implemented"
	return *new(aggregationData)
}

// updateWithSpan incorporates a single span into the aggregation statistics,
// tracking min/max durations and time ranges.
func (data *aggregationData) updateWithSpan(span ptrace.Span, isFirst bool, histogramBuckets []time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Calculate duration statistics
