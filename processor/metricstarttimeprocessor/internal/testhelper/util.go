// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testhelper // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor/internal/testhelper"

import (
	"context"
	"testing"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TimestampFromMs(timeAtMs int64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

type KV struct {
	Key, Value string
}

func Metrics(metrics ...pmetric.Metric) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func MetricsFromResourceMetrics(metrics ...pmetric.ResourceMetrics) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func ResourceMetrics(job, instance string, metrics ...pmetric.Metric) pmetric.ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ResourceMetrics)
}

func HistogramPointRaw(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp) pmetric.HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.HistogramDataPoint)
}

func HistogramPoint(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp, bounds []float64, counts []uint64) pmetric.HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.HistogramDataPoint)
}

func HistogramPointNoValue(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp) pmetric.HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.HistogramDataPoint)
}

func HistogramMetric(name string, points ...pmetric.HistogramDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// By default the AggregationTemporality is Cumulative until it'll be changed by the caller.

func ExponentialHistogramMetric(name string, points ...pmetric.ExponentialHistogramDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// By default the AggregationTemporality is Cumulative until it'll be changed by the caller.

func ExponentialHistogramPointRaw(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp) pmetric.ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.ExponentialHistogramDataPoint)
}

func ExponentialHistogramPoint(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp, scale int32, zeroCount uint64, negativeOffset int32, negativeBuckets []uint64, positiveOffset int32, positiveBuckets []uint64) pmetric.ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.ExponentialHistogramDataPoint)
}

func ExponentialHistogramPointNoValue(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp) pmetric.ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.ExponentialHistogramDataPoint)
}

// exponentialHistogramPointSimplified let's you define an exponential
// histogram with just a few parameters.
// Scale and ZeroCount are set to the provided values.
// Positive and negative buckets are generated using the offset and bucketCount
// parameters by adding buckets from offset in both positive and negative
// directions. Bucket counts start from 1 and increase by 1 for each bucket.
// Sum and Count will be proportional to the bucket count.
func ExponentialHistogramPointSimplified(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp, scale int32, zeroCount uint64, offset int32, bucketCount int) pmetric.ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.ExponentialHistogramDataPoint)
}

func NumberPointRaw(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp) pmetric.NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPoint)
}

func DoublePoint(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp, value float64) pmetric.NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPoint)
}

func IntPoint(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp, value int64) pmetric.NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPoint)
}

func DoublePointNoValue(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp) pmetric.NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPoint)
}

func GaugeMetric(name string, points ...pmetric.NumberDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func SumMetric(name string, points ...pmetric.NumberDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func SummaryPointRaw(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp) pmetric.SummaryDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.SummaryDataPoint)
}

func SummaryPoint(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp, count uint64, sum float64, quantiles, values []float64) pmetric.SummaryDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.SummaryDataPoint)
}

func SummaryPointNoValue(attributes []*KV, startTimestamp, timestamp pcommon.Timestamp) pmetric.SummaryDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.SummaryDataPoint)
}

func SummaryMetric(name string, points ...pmetric.SummaryDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

type Adjuster interface {
	AdjustMetrics(context.Context, pmetric.Metrics) (pmetric.Metrics, error)
}

type MetricsAdjusterTest struct {
	Description string
	Metrics     pmetric.Metrics
	Adjusted    pmetric.Metrics
}

func RunScript(t *testing.T, ma Adjuster, tests []*MetricsAdjusterTest, additionalResourceAttrs ...string) {
	_ = "STUB: not implemented"
	return
}

// Add the instance/job to the input metrics if they aren't already present.

// Add the instance/job to the expected metrics as well if they aren't already present.
