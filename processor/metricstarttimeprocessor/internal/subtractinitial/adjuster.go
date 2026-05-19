// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package subtractinitial // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor/internal/subtractinitial"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor/internal/datapointstorage"
)

// Type is the value users can use to configure the subtract initial point adjuster.
// The subtract initial point adjuster sets the start time of all points in a series by:
//   - Dropping the initial point, and recording its value and timestamp.
//   - Subtracting the initial point from all subsequent points, and using the timestamp of the initial point as the start timestamp.
//
// Note that when a reset is detected (eg: value of a counter is decreasing) - the strategy will set the
// start time of the reset point as point timestamp - 1ms.
const Type = "subtract_initial_point"

type Adjuster struct {
	referenceCache *datapointstorage.Cache
	set            component.TelemetrySettings
}

// NewAdjuster returns a new Adjuster which adjust metrics' start times based on the initial received points.
func NewAdjuster(set component.TelemetrySettings, gcInterval time.Duration) *Adjuster {
	_ = "STUB: not implemented"
	return nil
}

// AdjustMetrics adjusts the start time of metrics based on the initial received
// points.
//
// It uses two caches: referenceCache to store the initial point of each
// timeseries, and previousValueCache to store the previous value of each
// timeseries for reset detection.
//
// If a point has not been seen before, it will be dropped and cached as the
// reference point. For each subsequent point, it will normalize against the
// reference cached point reporting the delta. If a reset is detected, the
// current point will be reported as is, and the reference point will be
// updated. The function returns a new pmetric.Metrics containing the adjusted
// metrics.
func (a *Adjuster) AdjustMetrics(_ context.Context, metrics pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// The lock on the relevant timeseriesMap is held throughout the adjustment process to ensure that
// nothing else can modify the data used for adjustment.

func adjustMetricHistogram(referenceValueTsm *datapointstorage.TimeseriesMap, metric pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

// Only dealing with CumulativeDistributions.

// Report point as is.

// First time we see this point. Skip it and use as a reference point for the next points.

// Adjust the datapoint based on the reference value.

// reset re-initialize everything and use the non adjusted points start time.

// Update the reference value with the metric point.

func adjustMetricExponentialHistogram(referenceValueTsm *datapointstorage.TimeseriesMap, metric pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

// Only dealing with CumulativeDistributions.

// Report point as is.

// First time we see this point. Skip it and use as a reference point for the next points.

// Adjust the datapoint based on the reference value.

// reset re-initialize everything and use the non adjusted points start time.

func adjustMetricSum(referenceValueTsm *datapointstorage.TimeseriesMap, metric pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

// Only handle cumulative temporality sums

// Report point as is.

// First time we see this point. Skip it and use as a reference point for the next points.

// Adjust the datapoint based on the reference value.

// reset re-initialize everything and use the non adjusted points start time.

// Update the currentSum appropriately based on the original value type.

func adjustMetricSummary(referenceValueTsm *datapointstorage.TimeseriesMap, metric pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

// Report point as is.

// First time we see this point. Skip it and use as a reference point for the next points.

// Adjust the datapoint based on the reference value.

// reset re-initialize everything and use the non adjusted points start time.

// subtractHistogramDataPoint subtracts b from a.
func subtractHistogramDataPoint(a pmetric.HistogramDataPoint, ref datapointstorage.HistogramInfo) {
	_ = "STUB: not implemented"
	return
}

// Post reset, the reference histogram will have no buckets.

// subtractExponentialHistogramDataPoint subtracts b from a.
func subtractExponentialHistogramDataPoint(a pmetric.ExponentialHistogramDataPoint, ref datapointstorage.ExponentialHistogramInfo) {
	_ = "STUB: not implemented"
	return
}

// Post reset, the reference histogram will have no buckets.

// subtractExponentialBuckets subtracts b from a.
func subtractExponentialBuckets(a pmetric.ExponentialHistogramDataPointBuckets, b datapointstorage.ExponentialHistogramBucketInfo) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

// if there is no corresponding bucket for the starting BucketCounts, don't normalize
