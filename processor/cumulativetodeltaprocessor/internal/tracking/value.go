// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracking // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor/internal/tracking"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ValuePoint struct {
	ObservedTimestamp         pcommon.Timestamp
	FloatValue                float64
	IntValue                  int64
	HistogramValue            *HistogramPoint
	ExponentialHistogramValue *ExponentialHistogramPoint
}

type HistogramPoint struct {
	Count        uint64
	Sum          float64
	BucketBounds []float64
	BucketCounts []uint64
}

func (point *HistogramPoint) Clone() HistogramPoint {
	_ = "STUB: not implemented"
	return *new(HistogramPoint)
}

type ExponentialBuckets struct {
	Offset       int32
	BucketCounts []uint64
}

// Coarsen reduces an exponential histogram's scale by bitsLost,
// which amounts to dividing bucket indices by 2**bitsLost and merging buckets with the same resulting index.
func (buckets *ExponentialBuckets) Coarsen(bitsLost int32) (out ExponentialBuckets) {
	_ = "STUB: not implemented"
	return *new(ExponentialBuckets)
}

// TrimZeros removes buckets below a given bucket index, returning the sum of their counts.
// This is used when increasing the zero threshold: the removed count will be added to the zero count.
func (buckets *ExponentialBuckets) TrimZeros(thresholdBucket int32) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// Diff computes the delta between two sets of buckets with the same scale.
func (buckets *ExponentialBuckets) Diff(old *ExponentialBuckets) (out ExponentialBuckets) {
	_ = "STUB: not implemented"
	return *new(ExponentialBuckets)
}

// bucketCount < oldCount is a monotonicity error, we'll consider the diff to be zero as fallback

type ExponentialHistogramPoint struct {
	Count         uint64
	Sum           float64
	ZeroCount     uint64
	ZeroThreshold float64
	Scale         int32
	Positive      ExponentialBuckets
	Negative      ExponentialBuckets
}

func (point *ExponentialHistogramPoint) Clone() ExponentialHistogramPoint {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramPoint)
}
