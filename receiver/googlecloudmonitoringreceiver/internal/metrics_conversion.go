// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudmonitoringreceiver/internal"

import (
	"regexp"

	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/api/distribution"
	"google.golang.org/protobuf/types/known/anypb"
)

type MetricsBuilder struct {
	logger *zap.Logger
}

const (
	spanContextTypeURL   = "type.googleapis.com/google.monitoring.v3.SpanContext"
	droppedLabelsTypeURL = "type.googleapis.com/google.monitoring.v3.DroppedLabels"
)

var spanNameRegex = regexp.MustCompile("projects/[^/]*/traces/(?P<traceId>[[:alnum:]]*)/spans/(?P<spanId>[[:alnum:]]*)")

func NewMetricsBuilder(logger *zap.Logger) *MetricsBuilder { _ = "STUB: not implemented"; return nil }

func (mb *MetricsBuilder) ConvertGaugeToMetrics(ts *monitoringpb.TimeSeries, m pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// Directly check and set the StartTimestamp if valid

// Check if EndTime is set and valid

func (mb *MetricsBuilder) ConvertSumToMetrics(ts *monitoringpb.TimeSeries, m pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// Directly check and set the StartTimestamp if valid

// Check if EndTime is set and valid

func (mb *MetricsBuilder) ConvertDeltaToMetrics(ts *monitoringpb.TimeSeries, m pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// Directly check and set the StartTimestamp if valid

// Check if EndTime is set and valid

// ConvertDistributionToMetrics converts from Google cloud monitoring distributions to OpenTelemetry histograms.
// See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/TypedValue#Distribution for information on distribution metrics.
func (mb *MetricsBuilder) ConvertDistributionToMetrics(ts *monitoringpb.TimeSeries, m pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// Note: Google cloud monitoring distributions use inclusive lower bound and exclusive upper bound (see
// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/TypedValue#bucketoptions), while OpenTelemetry histograms use
// exclusive lower bounds and inclusive upper bounds. For the conversion, we are deliberately ignoring this discrepancy in
// accordance with https://opentelemetry.io/docs/specs/otel/metrics/data-model/#histogram-bucket-inclusivity, quote:
// > Importers and exporters working with OpenTelemetry Metrics data are meant to disregard this specification when
// > translating to and from histogram formats that use inclusive lower bounds and exclusive upper bounds.

// fall back to the current time stamp if no timestamp is available on the source data point

// The number of elements in bucket_counts array must be by one greater than the number of elements in explicit_bounds array.

// The exception to this rule is when the length of bucket_counts is 0, then the length of explicit_bounds must also be 0.
// Google says this case should never happen, but you never know...

// A Histogram without buckets conveys a population in terms of only the sum and count, and may be interpreted as a
// histogram with single bucket covering (-Inf, +Inf).

// If present, `bucket_counts` should contain N values, where N is the number
// of buckets specified in `bucket_options`. If you supply fewer than N
// values, the remaining values are assumed to be 0.
// (see https://pkg.go.dev/google.golang.org/genproto/googleapis/api/distribution#Distribution)

// The source data type is int64, the target type is uint64, so we need to handle negative counts in some way.
// We normalize negative values to zero, so all other counts are at the correct position in the target data point.
// (Obviously, bucket counts should never be negative in the source data anyway, so this branch should never be
// executed.)

func convertLabelsToMetricAttributes(sourceLabels map[string]string) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func (mb *MetricsBuilder) convertDistributionDataPointExplicitBuckets(
	metricType string,
	numberOfSourceBucketCounts int,
	buckets *distribution.Distribution_BucketOptions_Explicit,
	targetDataPoint *pmetric.HistogramDataPoint,
) {
	_ = "STUB: not implemented"
	return
}

// Note: There is also an implicit overflow bucket with boundary +Inf.

func (mb *MetricsBuilder) convertDistributionDataPointLinearBuckets(
	metricType string,
	buckets *distribution.Distribution_BucketOptions_Linear,
	targetDataPoint *pmetric.HistogramDataPoint,
) {
	_ = "STUB: not implemented"
	return
}

// See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/TypedValue#linear:
// There are numFiniteBuckets + 2 (= N) buckets. Bucket i has the following boundaries:
// Lower bound (1 <= i < N): offset + (width * (i - 1)).
// Upper bound (0 <= i < N-1): offset + (width * i).

// bucket 0: [-infinity, offset) and
// buckets 1 - N-1: [offset + (width * (i - 1)), offset + (width * i))

// bucket N: implicit overflow bucket [offset + (width * (N - 1)), +infinity)

func (mb *MetricsBuilder) convertDistributionDataPointExponentialBuckets(
	metricType string,
	buckets *distribution.Distribution_BucketOptions_Exponential,
	targetDataPoint *pmetric.HistogramDataPoint,
) {
	_ = "STUB: not implemented"
	// Note: This method converts a distribution with exponential buckets to an OpenTelemetry histogram with explicit bounds.
	// An obvious alternative would be to convert it to an exponential histogram. However, this cannot be done without loss of
	// precision. An OpenTelemetry exponential histograms sets the boundary for bucket i at (2^(2^(-scale)))^i, with scale being
	// an integer. Google Cloud Monitoring exponential buckets use a floating point scale and growth factor and set the boundary
	// at scale * (growthFactor ^ i). Depending on the chosen scale and growth factor, even the best approximation of the
	// Google Cloud Monitoring bucket boundaries might deviate from the actual boundaries significantly. Therefore, we instead
	// convert the exponential distribution to a plain OpenTelemetry histogram with explicit bounds. This choice trades payload
	// size for precision (that is, we avoid loss of precision/information by accepting a larger payload size).
	return
}

// See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/TypedValue#exponential:
// There are numFiniteBuckets + 2 (= N) buckets. Bucket i has the following boundaries:
// Lower bound (1 <= i < N): scale * (growthFactor ^ (i - 1)).
// Upper bound (0 <= i < N-1): scale * (growthFactor ^ i).

// bucket 0: [-infinity, offset) and
// buckets 1 - N-1: [offset + (width * (i - 1)), offset + (width * i))

// bucket N: implicit overflow bucket [offset + (width * (N - 1)), +infinity)

func (mb *MetricsBuilder) convertExemplars(distributionValue *distribution.Distribution, targetDataPoint *pmetric.HistogramDataPoint) {
	_ = "STUB: not implemented"
	return
}

// skip attachments with nil or empty value

func (mb *MetricsBuilder) convertDistributionExemplarAttachmentSpanContext(sourceValueBytes []byte, targetExemplarDataPoint *pmetric.Exemplar) {
	_ = "STUB: not implemented"
	return
}

func (mb *MetricsBuilder) convertDistributionExemplarAttachmentDroppedLabels(sourceValue *anypb.Any, targetExemplarDataPoint *pmetric.Exemplar) {
	_ = "STUB: not implemented"
	return
}
