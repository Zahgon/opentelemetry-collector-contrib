// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datapoints // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/datapoints"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
)

type Histogram struct {
	pmetric.HistogramDataPoint
	elasticsearch.MappingHintGetter
	metric pmetric.Metric
}

func NewHistogram(metric pmetric.Metric, dp pmetric.HistogramDataPoint) Histogram {
	_ = "STUB: not implemented"
	return *new(Histogram)
}

func (dp Histogram) Value() (pcommon.Value, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), nil
}

func (dp Histogram) DynamicTemplate(_ pmetric.Metric, mode DynamicTemplateMode) string {
	_ = "STUB: not implemented"
	return ""
}

func (dp Histogram) DocCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (dp Histogram) Metric() pmetric.Metric { _ = "STUB: not implemented"; return *new(pmetric.Metric) }

func histogramToValue(dp pmetric.HistogramDataPoint, metric pmetric.Metric, raw bool) (pcommon.Value, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), nil
}

// It is possible for explicit bounds to be nil. In this case create
// a bucket using the count and sum which are required to be present.
// See https://opentelemetry.io/docs/specs/otel/metrics/data-model/#histogram

// In raw mode, the overflow bucket would have the same value
// as the last real bucket. Merge the overflow count into the
// last real bucket to avoid duplicate values, which violates
// ES histogram's strictly increasing values requirement.

// midpointBucketValue returns the midpoint centroid for bucket at index i.
// Caller must ensure len(bucketCounts) == len(explicitBounds) + 1.
func midpointBucketValue(explicitBounds pcommon.Float64Slice, i int) float64 {
	_ = "STUB: not implemented"

	// (-infinity, explicit_bounds[i]]
	return 0
}

// (explicit_bounds[i-1], +infinity)

// [explicit_bounds[i-1], explicit_bounds[i])

// rawBucketValue returns the explicit bound for bucket at index i without
// any midpoint approximation.
func rawBucketValue(explicitBounds pcommon.Float64Slice, i int) float64 {
	_ = "STUB: not implemented"
	return 0
}
