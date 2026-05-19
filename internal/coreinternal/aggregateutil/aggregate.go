// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aggregateutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/aggregateutil"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func CopyMetricDetails(from, to pmetric.Metric) { _ = "STUB: not implemented"; return }

//exhaustive:enforce

func FilterAttrs(metric pmetric.Metric, filterAttrKeys []string) {
	_ = "STUB: not implemented"
	// filterAttrKeys being nil means the filter is to be skipped.
	return
}

// filterAttrKeys being empty means it is explicitly expected to filter
// against an empty label set, which is functionally the same as removing
// all attributes.

// filterAttrKeys having provided attributes means the filter continues
// as normal.

func GroupDataPoints(metric pmetric.Metric, ag *AggGroups) { _ = "STUB: not implemented"; return }

func MergeDataPoints(to pmetric.Metric, aggType AggregationType, ag AggGroups) {
	_ = "STUB: not implemented"
	return
}

// RangeDataPointAttributes calls f sequentially on attributes of every metric data point.
// The iteration terminates if f returns false.
func RangeDataPointAttributes(metric pmetric.Metric, f func(pcommon.Map) bool) {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return
}

func mergeNumberDataPoints(dpsMap map[string]pmetric.NumberDataPointSlice, agg AggregationType, to pmetric.NumberDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func doubleVal(dp pmetric.NumberDataPoint) float64 { _ = "STUB: not implemented"; return 0 }

func intVal(dp pmetric.NumberDataPoint) int64 { _ = "STUB: not implemented"; return 0 }

func mergeHistogramDataPoints(dpsMap map[string]pmetric.HistogramDataPointSlice, to pmetric.HistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func mergeExponentialHistogramDataPoints(dpsMap map[string]pmetric.ExponentialHistogramDataPointSlice,
	to pmetric.ExponentialHistogramDataPointSlice,
) {
	_ = "STUB: not implemented"
	return
}

// Check if offsets are different, indicating we need to adjust dp offsets and bucket counts after merging

// Only adjust offsets and remove leading zero buckets if we had different offsets and the buckets are not empty

func mergeExponentialHistogramBuckets(tgt, src pcommon.UInt64Slice, tgtOff, srcOff int32) {
	_ = "STUB: not implemented"
	// Both data points have the same offset - simple element-wise addition
	return
}

// Source offset is less than target offset - source data point covers lower values

// Source offset is greater than target offset - source data point covers higher values

func trimBuckets(buckets pcommon.UInt64Slice) { _ = "STUB: not implemented"; return }

func groupNumberDataPoints(dps pmetric.NumberDataPointSlice, useStartTime bool,
	dpsByAttrsAndTs map[string]pmetric.NumberDataPointSlice,
) {
	_ = "STUB: not implemented"
	return
}

func groupHistogramDataPoints(dps pmetric.HistogramDataPointSlice, useStartTime bool,
	dpsByAttrsAndTs map[string]pmetric.HistogramDataPointSlice,
) {
	_ = "STUB: not implemented"
	return
}

func groupExponentialHistogramDataPoints(dps pmetric.ExponentialHistogramDataPointSlice, useStartTime bool,
	dpsByAttrsAndTs map[string]pmetric.ExponentialHistogramDataPointSlice,
) {
	_ = "STUB: not implemented"
	return
}

func dataPointHashKey(atts pcommon.Map, ts pcommon.Timestamp, other ...any) string {
	_ = "STUB: not implemented"
	return ""
}
