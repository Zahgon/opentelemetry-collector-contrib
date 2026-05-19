// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datapointstorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor/internal/datapointstorage"

import (
	"sync"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// AttributeHash is used to store a hash of attributes for a metric. See pdatautil.MapHash for more details.
type AttributeHash [16]byte

// TimeseriesInfo contains the information necessary to adjust from the initial point and to detect resets.
type TimeseriesInfo struct {
	Mark bool

	Number               NumberInfo
	Histogram            HistogramInfo
	ExponentialHistogram ExponentialHistogramInfo
	Summary              SummaryInfo
}

type NumberInfo struct {
	StartTime           pcommon.Timestamp
	PreviousDoubleValue float64
	PreviousIntValue    int64

	// These are the optional reference values for strategies that need to cache
	// additional data from the initial points.
	// For example - storing the initial point for the subtract_initial_point strategy.
	RefDoubleValue float64
	RefIntValue    int64
}

type HistogramInfo struct {
	StartTime            pcommon.Timestamp
	PreviousCount        uint64
	PreviousSum          float64
	PreviousBucketCounts []uint64
	ExplicitBounds       []float64

	// These are the optional reference values for strategies that need to cache
	// additional data from the initial points.
	// For example - storing the initial point for the subtract_initial_point strategy.
	RefCount        uint64
	RefSum          float64
	RefBucketCounts []uint64
}

type ExponentialHistogramInfo struct {
	StartTime         pcommon.Timestamp
	PreviousCount     uint64
	PreviousSum       float64
	PreviousZeroCount uint64
	Scale             int32
	PreviousPositive  ExponentialHistogramBucketInfo
	PreviousNegative  ExponentialHistogramBucketInfo

	// These are the optional reference values for strategies that need to cache
	// additional data from the initial points.
	// For example - storing the initial point for the subtract_initial_point strategy.
	RefCount     uint64
	RefSum       float64
	RefZeroCount uint64
	RefPositive  ExponentialHistogramBucketInfo
	RefNegative  ExponentialHistogramBucketInfo
}

type ExponentialHistogramBucketInfo struct {
	Offset       int32
	BucketCounts []uint64
}

func NewExponentialHistogramBucketInfo(ehdpb pmetric.ExponentialHistogramDataPointBuckets) ExponentialHistogramBucketInfo {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramBucketInfo)
}

type SummaryInfo struct {
	StartTime     pcommon.Timestamp
	PreviousCount uint64
	PreviousSum   float64

	// These are the optional reference values for strategies that need to cache
	// additional data from the initial points.
	// For example - storing the initial point for the subtract_initial_point strategy.
	RefCount uint64
	RefSum   float64
}

type TimeseriesKey struct {
	Name           string
	Attributes     [16]byte
	AggTemporality pmetric.AggregationTemporality
}

// TimeseriesMap maps from a timeseries instance (metric * label values) to the timeseries info for
// the instance.
type TimeseriesMap struct {
	sync.RWMutex
	// The mutex is used to protect access to the member fields. It is acquired for the entirety of
	// AdjustMetricSlice() and also acquired by gc().

	Mark   bool
	TsiMap map[TimeseriesKey]*TimeseriesInfo
}

// Get the TimeseriesInfo for the timeseries associated with the metric and label values.
func (tsm *TimeseriesMap) Get(metric pmetric.Metric, kv pcommon.Map) (*TimeseriesInfo, bool) {
	_ = "STUB: not implemented"
	// This should only be invoked be functions called (directly or indirectly) by AdjustMetricSlice().
	// The lock protecting tsm.tsiMap is acquired there.
	return nil, false
}

// There are 2 types of Histograms whose aggregation temporality needs distinguishing:
// * CumulativeHistogram
// * GaugeHistogram

// There are 2 types of ExponentialHistograms whose aggregation temporality needs distinguishing:
// * CumulativeHistogram
// * GaugeHistogram

// Remove timeseries that have aged out.
func (tsm *TimeseriesMap) GC() { _ = "STUB: not implemented"; return }

// IsResetHistogram compares the given histogram datapoint h, to ref
// and determines whether the metric has been reset based on the values.  It is
// a reset if any of the bucket boundaries have changed, if any of the bucket
// counts have decreased or if the total sum or count have decreased.
func (ref *TimeseriesInfo) IsResetHistogram(h pmetric.HistogramDataPoint) bool {
	_ = "STUB: not implemented"
	return false
}

// Guard against bucket boundaries changes.

// We need to check individual buckets to make sure the counts are all increasing.

// IsResetExponentialHistogram compares the given exponential histogram
// datapoint eh, to ref and determines whether the metric
// has been reset based on the values.  It is a reset if any of the bucket
// boundaries have changed, if any of the bucket counts have decreased or if the
// total sum or count have decreased.
func (ref *TimeseriesInfo) IsResetExponentialHistogram(eh pmetric.ExponentialHistogramDataPoint) bool {
	_ = "STUB: not implemented"
	// Same as the histogram implementation
	return false
}

// Guard against bucket boundaries changes.

// We need to check individual buckets to make sure the counts are all increasing.

// IsResetSummary compares the given summary datapoint s to ref and
// determines whether the metric has been reset based on the values.  It is a
// reset if the count or sum has decreased.
func (ref *TimeseriesInfo) IsResetSummary(s pmetric.SummaryDataPoint) bool {
	_ = "STUB: not implemented"
	return false
}

// IsResetSum compares the given number datapoint s to ref and determines
// whether the metric has been reset based on the values.  It is a reset if the
// value has decreased.
func (ref *TimeseriesInfo) IsResetSum(s pmetric.NumberDataPoint) bool {
	_ = "STUB: not implemented"
	return false
}

func newTimeseriesMap() *TimeseriesMap { _ = "STUB: not implemented"; return nil }
