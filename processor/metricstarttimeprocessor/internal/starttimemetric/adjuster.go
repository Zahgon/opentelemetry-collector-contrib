// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package starttimemetric // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor/internal/starttimemetric"

import (
	"context"
	"errors"
	"regexp"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor/internal/datapointstorage"
)

const (
	// Type is the value users can use to configure the start time metric adjuster.
	Type                = "start_time_metric"
	startTimeMetricName = "process_start_time_seconds"
)

var (
	errNoStartTimeMetrics             = errors.New("start_time metric is missing")
	errNoDataPointsStartTimeMetric    = errors.New("start time metric with no data points")
	errUnsupportedTypeStartTimeMetric = errors.New("unsupported data type for start time metric")
	// approximateCollectorStartTime is the approximate start time of the
	// collector. Used as a fallback start time for metrics when the start time
	// metric is not found. Set when the component is initialized.
	approximateCollectorStartTime time.Time
)

func init() {
	approximateCollectorStartTime = time.Now()
}

type Adjuster struct {
	referenceValueCache  *datapointstorage.Cache
	startTimeMetricRegex *regexp.Regexp
	set                  component.TelemetrySettings
}

// NewAdjuster returns a new Adjuster which adjust metrics' start times based on the initial received points.
func NewAdjuster(set component.TelemetrySettings, startTimeMetricRegex *regexp.Regexp, gcInterval time.Duration) *Adjuster {
	_ = "STUB: not implemented"
	return nil
}

// AdjustMetrics adjusts the start time of metrics based on a different metric in the batch.
func (a *Adjuster) AdjustMetrics(_ context.Context, metrics pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// The lock on the relevant timeseriesMap is held throughout the adjustment process to ensure that
// nothing else can modify the data used for adjustment.

// Report point as is if the start timestamp is already set.

// Report point as is if the start timestamp is already set.

// Report point as is if the start timestamp is already set.

// Report point as is if the start timestamp is already set.

func timestampFromFloat64(ts float64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (a *Adjuster) getStartTime(metrics pmetric.Metrics) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (a *Adjuster) matchStartTimeMetric(metricName string) bool {
	_ = "STUB: not implemented"
	return false
}
