// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func (c *prometheusConverterV2) addExponentialHistogramDataPoints(dataPoints pmetric.ExponentialHistogramDataPointSlice,
	resource pcommon.Resource, scope pcommon.InstrumentationScope, settings Settings, name string, metadata metadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

// exponentialToNativeHistogramV2 translates OTel Exponential Histogram data point
// to Prometheus Native Histogram.
func exponentialToNativeHistogramV2(p pmetric.ExponentialHistogramDataPoint) (writev2.Histogram, error) {
	_ = "STUB: not implemented"
	return *new(writev2.Histogram), nil
}

// The counter reset detection must be compatible with Prometheus to
// safely set ResetHint to NO. This is not ensured currently.
// Sending a sample that triggers counter reset but with ResetHint==NO
// would lead to Prometheus panic as it does not double check the hint.
// Thus we're explicitly saying UNKNOWN here, which is always safe.
// TODO: using created time stamp should be accurate, but we
// need to know here if it was used for the detection.
// Ref: https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/28663#issuecomment-1810577303
// Counter reset detection in Prometheus: https://github.com/prometheus/prometheus/blob/f997c72f294c0f18ca13fa06d51889af04135195/tsdb/chunkenc/histogram.go#L232

// convertBucketsLayoutV2 translates OTel Exponential Histogram dense buckets
// representation to Prometheus Native Histogram sparse bucket representation.
//
// The translation logic is taken from the client_golang `histogram.go#makeBuckets`
// function, see `makeBuckets` https://github.com/prometheus/client_golang/blob/main/prometheus/histogram.go
// The bucket indexes conversion was adjusted, since OTel exp. histogram bucket
// index 0 corresponds to the range (1, base] while Prometheus bucket index 0
// to the range (base 1].
//
// scaleDown is the factor by which the buckets are scaled down. In other words 2^scaleDown buckets will be merged into one.
func convertBucketsLayoutV2(buckets pmetric.ExponentialHistogramDataPointBuckets, scaleDown int32) ([]writev2.BucketSpan, []int64) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Let the compiler figure out that this is const during this function by
// moving it into a local variable.

// The offset is scaled and adjusted by 1 as described above.

// The offset is scaled and adjusted by 1 as described above.

// We have not collected enough buckets to merge yet.

// We have to create a new span, because we have found a gap
// of more than two buckets. The constant 2 is copied from the logic in
// https://github.com/prometheus/client_golang/blob/27f0506d6ebbb117b6b697d0552ee5be2502c5f2/prometheus/histogram.go#L1296

// We have found a small gap (or no gap at all).
// Insert empty buckets as needed.

// Need to use the last item's index. The offset is scaled and adjusted by 1 as described above.

// We have to create a new span, because we have found a gap
// of more than two buckets. The constant 2 is copied from the logic in
// https://github.com/prometheus/client_golang/blob/27f0506d6ebbb117b6b697d0552ee5be2502c5f2/prometheus/histogram.go#L1296

// We have found a small gap (or no gap at all).
// Insert empty buckets as needed.
