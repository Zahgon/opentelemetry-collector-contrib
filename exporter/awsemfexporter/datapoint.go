// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsemfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	aws "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"
)

const (
	summaryCountSuffix = "_count"
	summarySumSuffix   = "_sum"
)

type emfCalculators struct {
	delta   aws.MetricCalculator
	summary aws.MetricCalculator
}

func calculateSummaryDelta(prev *aws.MetricValue, val any, _ time.Time) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// dataPoint represents a processed metric data point
type dataPoint struct {
	name        string
	value       any
	labels      map[string]string
	timestampMs int64
}

// dataPoints is a wrapper interface for:
//   - pmetric.NumberDataPointSlice
//   - pmetric.HistogramDataPointSlice
//   - pmetric.SummaryDataPointSlice
type dataPoints interface {
	Len() int
	// CalculateDeltaDatapoints calculates the delta datapoint from the DataPointSlice at i-th index
	// for some type (Counter, Summary)
	// dataPoint: the adjusted data point
	// retained: indicates whether the data point is valid for further process
	// NOTE: It is an expensive call as it calculates the metric value.
	CalculateDeltaDatapoints(i int, instrumentationScopeName string, detailedMetrics bool, calculators *emfCalculators) (dataPoint []dataPoint, retained bool)
	// IsStaleNaNInf returns true if metric value has NoRecordedValue flag set or if any metric value contains a NaN or Inf.
	// When return value is true, IsStaleNaNInf also returns the attributes attached to the metric which can be used for
	// logging purposes.
	IsStaleNaNInf(i int) (bool, pcommon.Map)
}

// deltaMetricMetadata contains the metadata required to perform rate/delta calculation
type deltaMetricMetadata struct {
	adjustToDelta              bool
	retainInitialValueForDelta bool
	metricName                 string
	namespace                  string
	logGroup                   string
	logStream                  string
}

// numberDataPointSlice is a wrapper for pmetric.NumberDataPointSlice
type numberDataPointSlice struct {
	deltaMetricMetadata
	pmetric.NumberDataPointSlice
}

// histogramDataPointSlice is a wrapper for pmetric.HistogramDataPointSlice
type histogramDataPointSlice struct {
	// Todo:(khanhntd) Calculate delta value for count and sum value with histogram
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/18245
	deltaMetricMetadata
	pmetric.HistogramDataPointSlice
}

type exponentialHistogramDataPointSlice struct {
	// TODO: Calculate delta value for count and sum value with exponential histogram
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/18245
	deltaMetricMetadata
	pmetric.ExponentialHistogramDataPointSlice
}

// summaryDataPointSlice is a wrapper for pmetric.SummaryDataPointSlice
type summaryDataPointSlice struct {
	deltaMetricMetadata
	pmetric.SummaryDataPointSlice
}

type summaryMetricEntry struct {
	sum   float64
	count uint64
}

// dataPointSplit is a structure used to manage segments of data points split from a histogram.
// It is not safe for concurrent use.
type dataPointSplit struct {
	cWMetricHistogram *cWMetricHistogram
	length            int
	capacity          int
}

func (split *dataPointSplit) isFull() bool { _ = "STUB: not implemented"; return false }

func (split *dataPointSplit) setMax(maxVal float64) { _ = "STUB: not implemented"; return }

func (split *dataPointSplit) setMin(minVal float64) { _ = "STUB: not implemented"; return }

func (split *dataPointSplit) appendMetricData(metricVal float64, count uint64) {
	_ = "STUB: not implemented"
	return
}

// CalculateDeltaDatapoints retrieves the NumberDataPoint at the given index and performs rate/delta calculation if necessary.
func (dps numberDataPointSlice) CalculateDeltaDatapoints(i int, instrumentationScopeName string, _ bool, calculators *emfCalculators) ([]dataPoint, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// If a delta to the previous data point could not be computed use the current metric value instead

// It should not happen in practice that the previous metric value is smaller than the current one.
// If it happens, we assume that the metric is reset for some reason.

func (dps numberDataPointSlice) IsStaleNaNInf(i int) (bool, pcommon.Map) {
	_ = "STUB: not implemented"
	return false, *new(pcommon.Map)
}

// CalculateDeltaDatapoints retrieves the HistogramDataPoint at the given index.
func (dps histogramDataPointSlice) CalculateDeltaDatapoints(i int, instrumentationScopeName string, _ bool, _ *emfCalculators) ([]dataPoint, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (dps histogramDataPointSlice) IsStaleNaNInf(i int) (bool, pcommon.Map) {
	_ = "STUB: not implemented"
	return false, *new(pcommon.Map)
}

// CalculateDeltaDatapoints retrieves the ExponentialHistogramDataPoint at the given index.
// As CloudWatch EMF logs allows in maximum of 100 target members, the exponential histogram metric are split into multiple data points as needed,
// each containing a maximum of 100 buckets, to comply with CloudWatch EMF log constraints.
// Note that the number of values and counts in each split may not be less than splitThreshold as we are only adding non-zero bucket counts.
//
// For each split data point:
// - Min and Max values are recalculated based on the bucket boundary within that specific split.
// - Sum is only assigned to the first split to ensure the total sum of the datapoints after aggregation is correct.
// - Count is accumulated based on the bucket counts within each split.
func (dps exponentialHistogramDataPointSlice) CalculateDeltaDatapoints(idx int, instrumentationScopeName string, _ bool, _ *emfCalculators) ([]dataPoint, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Create a new dataPointSplit with a capacity of up to splitThreshold buckets

// Only assign `Sum` if this is the first split to make sure the total sum of the datapoints after aggregation is correct.

// Set collect values from positive buckets and save into split.

// Set collect values from zero buckets and save into split.

// Set collect values from negative buckets and save into split.

// Add the current split to the datapoints list

// Override the min and max values of the first and last splits with the raw data of the metric.

func collectDatapointsWithPositiveBuckets(split *dataPointSplit, metric pmetric.ExponentialHistogramDataPoint, currentBucketIndex, currentPositiveIndex int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// The value are append from high to low, set Max from the first bucket (highest value) and Min from the last bucket (lowest value)

func collectDatapointsWithZeroBucket(split *dataPointSplit, metric pmetric.ExponentialHistogramDataPoint, currentBucketIndex, currentZeroIndex int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// The value are append from high to low, set Max from the first bucket (highest value) and Min from the last bucket (lowest value)

func collectDatapointsWithNegativeBuckets(split *dataPointSplit, metric pmetric.ExponentialHistogramDataPoint, currentBucketIndex, currentNegativeIndex int) (int, int) {
	_ = "STUB: not implemented"
	// According to metrics spec, the value in histogram is expected to be non-negative.
	// https://opentelemetry.io/docs/specs/otel/metrics/api/#histogram
	// However, the negative support is defined in metrics data model.
	// https://opentelemetry.io/docs/specs/otel/metrics/data-model/#exponentialhistogram
	// The negative is also supported but only verified with unit test.
	return 0, 0
}

// The value are append from high to low, set Max from the first bucket (highest value) and Min from the last bucket (lowest value)

func (dps exponentialHistogramDataPointSlice) IsStaleNaNInf(i int) (bool, pcommon.Map) {
	_ = "STUB: not implemented"
	return false, *new(pcommon.Map)
}

// CalculateDeltaDatapoints retrieves the SummaryDataPoint at the given index and perform calculation with sum and count while retain the quantile value.
func (dps summaryDataPointSlice) CalculateDeltaDatapoints(i int, instrumentationScopeName string, detailedMetrics bool, calculators *emfCalculators) ([]dataPoint, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// If a delta to the previous data point could not be computed use the current metric value instead

// Instead of sending metrics as a Statistical Set (contains min,max, count, sum), the emfexporter will enrich the
// values by sending each quantile values as a datapoint (from quantile 0 ... 1)

func (dps summaryDataPointSlice) IsStaleNaNInf(i int) (bool, pcommon.Map) {
	_ = "STUB: not implemented"
	return false, *new(pcommon.Map)
}

// createLabels converts OTel AttributesMap attributes to a map
// and optionally adds in the OTel instrumentation library name
func createLabels(attributes pcommon.Map, instrLibName string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Add OTel instrumentation lib name as an additional label if it is defined

// getDataPoints retrieves data points from OT Metric.
func getDataPoints(pmd pmetric.Metric, metadata cWMetricMetadata, logger *zap.Logger) dataPoints {
	_ = "STUB: not implemented"
	return *new(dataPoints)
}

//exhaustive:enforce

// For summaries coming from the prometheus receiver, the sum and count are cumulative, whereas for summaries
// coming from other sources, e.g. SDK, the sum and count are delta by being accumulated and reset periodically.
// In order to ensure metrics are sent as deltas, we check the receiver attribute (which can be injected by
// attribute processor) from resource metrics. If it exists, and is equal to prometheus, the sum and count will be
// converted.
// For more information: https://github.com/open-telemetry/opentelemetry-collector/blob/main/receiver/prometheusreceiver/DESIGN.md#summary
