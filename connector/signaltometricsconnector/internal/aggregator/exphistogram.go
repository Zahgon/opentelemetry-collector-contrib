// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aggregator // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/aggregator"

import (
	"time"

	"github.com/lightstep/go-expohisto/structure"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type exponentialHistogramDP struct {
	attrs pcommon.Map
	data  *structure.Histogram[float64]
}

func newExponentialHistogramDP(attrs pcommon.Map, maxSize int32) *exponentialHistogramDP {
	_ = "STUB: not implemented"
	return nil
}

func (dp *exponentialHistogramDP) Aggregate(value float64, count int64) {
	_ = "STUB: not implemented"
	return
}

func (dp *exponentialHistogramDP) Copy(
	timestamp time.Time,
	dest pmetric.ExponentialHistogramDataPoint,
) {
	_ = "STUB: not implemented"
	return
}

// TODO determine appropriate start time

// copyBucketRange copies a bucket range from exponential histogram
// datastructure to the OTel representation.
func copyBucketRange(
	src *structure.Buckets,
	dest pmetric.ExponentialHistogramDataPointBuckets,
) {
	_ = "STUB: not implemented"
	return
}
