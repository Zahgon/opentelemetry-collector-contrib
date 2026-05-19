// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package truereset // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor/internal/truereset"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor/internal/datapointstorage"
)

// Type is the value users can use to configure the true reset point adjuster.
// The true reset point adjuster sets the start time of all points in a series following:
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/data-model.md#cumulative-streams-inserting-true-reset-points.
// This involves setting the start time using the following strategy:
//   - The initial point in a series has its start time set to that point's end time.
//   - All subsequent points in the series have their start time set to the initial point's end time.
//
// Note that when a reset is detected (eg: value of a counter is decreasing) - the strategy will set the
// start time of the reset point as point timestamp - 1ms.
const Type = "true_reset_point"

// Adjuster takes a map from a metric instance to the initial point in the metrics instance
// and provides AdjustMetric, which takes a sequence of metrics and adjust their start times based on
// the initial points.
type Adjuster struct {
	startTimeCache *datapointstorage.Cache
	set            component.TelemetrySettings
}

// NewAdjuster returns a new Adjuster which adjust metrics' start times based on the initial received points.
func NewAdjuster(set component.TelemetrySettings, gcInterval time.Duration) *Adjuster {
	_ = "STUB: not implemented"
	return nil
}

// AdjustMetrics takes a sequence of metrics and adjust their start times based on the initial and
// previous points in the timeseriesMap.
func (a *Adjuster) AdjustMetrics(_ context.Context, metrics pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// The lock on the relevant timeseriesMap is held throughout the adjustment process to ensure that
// nothing else can modify the data used for adjustment.

// gauges don't need to be adjusted so no additional processing is necessary

// this shouldn't happen

func (*Adjuster) adjustMetricHistogram(tsm *datapointstorage.TimeseriesMap, current pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

// Only dealing with CumulativeDistributions.

// Report point as is if the start timestamp is already set.

// initialize everything.

// For the first point, set the start time as the point timestamp.

// TODO: Investigate why this does not reset.

// reset re-initialize everything.

// Update only previous values.

func (*Adjuster) adjustMetricExponentialHistogram(tsm *datapointstorage.TimeseriesMap, current pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

// Only dealing with CumulativeDistributions.

// Report point as is if the start timestamp is already set.

// initialize everything.

// For the first point, set the start time as the point timestamp.

// TODO: Investigate why this does not reset.

// reset re-initialize everything.

// Update only previous values.

func (*Adjuster) adjustMetricSum(tsm *datapointstorage.TimeseriesMap, current pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

// Report point as is if the start timestamp is already set.

// initialize everything.

// For the first point, set the start time as the point timestamp.

// TODO: Investigate why this does not reset.

// reset re-initialize everything.

// Update only previous values.

func (*Adjuster) adjustMetricSummary(tsm *datapointstorage.TimeseriesMap, current pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

// Report point as is if the start timestamp is already set.

// initialize everything.

// For the first point, set the start time as the point timestamp.

// TODO: Investigate why this does not reset.

// reset re-initialize everything.

// Update only previous values.
