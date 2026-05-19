// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package parser // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/parser"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/protocol"
)

var statsDDefaultPercentiles = []float64{0, 10, 50, 90, 95, 100}

func buildCounterMetric(parsedMetric statsDMetric, isMonotonicCounter bool, counterType protocol.CounterType) pmetric.ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ScopeMetrics)
}

// setCounterValue sets the counter data point value based on the configured counter type.
func setCounterValue(dp pmetric.NumberDataPoint, value float64, counterType protocol.CounterType) {
	_ = "STUB: not implemented"
	return
}

// protocol.CounterTypeInt or empty (default)

// aggregateCounterValue adds the value to an existing counter data point.
func aggregateCounterValue(dp pmetric.NumberDataPoint, value float64, counterType protocol.CounterType) {
	_ = "STUB: not implemented"
	return
}

// protocol.CounterTypeInt or empty (default)

// stochasticRound performs probabilistic rounding where the probability of rounding
// up equals the fractional part. This maintains integer type while being statistically
// accurate over time.
func stochasticRound(value float64) int64 { _ = "STUB: not implemented"; return 0 }

func setTimestampsForCounterMetric(ilm pmetric.ScopeMetrics, startTime, timeNow time.Time) {
	_ = "STUB: not implemented"
	return
}

func buildGaugeMetric(parsedMetric statsDMetric, timeNow time.Time) pmetric.ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ScopeMetrics)
}

func buildSummaryMetric(desc statsDMetricDescription, summary summaryMetric, startTime, timeNow time.Time, percentiles []float64, ilm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

// Note: count is rounded here, see note in counterValue().

func buildExplicitBucketHistogramMetric(desc statsDMetricDescription, histogram histogramMetric, startTime, timeNow time.Time, ilm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

// +1 to give space for the +Inf bucket

func buildExponentialBucketHistogramMetric(desc statsDMetricDescription, histogram histogramMetric, startTime, timeNow time.Time, ilm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func buildHistogramMetric(desc statsDMetricDescription, histogram histogramMetric, startTime, timeNow time.Time, ilm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func (s statsDMetric) counterValue() float64 {
	_ = "STUB: not implemented"

	// Note statsD counters are traditionally represented as integers, but
	// we'll preserve the floating point precision to avoid truncating
	// fractional values to zero.
	return 0
}

func (s statsDMetric) gaugeValue() float64 {
	_ = "STUB: not implemented"
	// sampleRate does not have effect for gauge points.
	return 0
}

func (s statsDMetric) sampleValue() sampleValue {
	_ = "STUB: not implemented"
	return *new(sampleValue)
}

type dualSorter struct {
	values, weights []float64
}

func (d dualSorter) Len() int { _ = "STUB: not implemented"; return 0 }

func (d dualSorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (d dualSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
