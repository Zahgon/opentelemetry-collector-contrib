// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tinybirdexporter/internal"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type baseMetricSignal struct {
	ResourceSchemaURL  string            `json:"resource_schema_url"`
	ResourceAttributes map[string]string `json:"resource_attributes"`
	ServiceName        string            `json:"service_name"`
	StartTimestamp     string            `json:"start_timestamp"`
	Timestamp          string            `json:"timestamp"`
	Flags              uint32            `json:"flags"`
	MetricName         string            `json:"metric_name"`
	MetricDescription  string            `json:"metric_description"`
	MetricUnit         string            `json:"metric_unit"`
	MetricAttributes   map[string]string `json:"metric_attributes"`
	ScopeName          string            `json:"scope_name"`
	ScopeVersion       string            `json:"scope_version"`
	ScopeSchemaURL     string            `json:"scope_schema_url"`
	ScopeAttributes    map[string]string `json:"scope_attributes"`
	exemplars
}

type genericDataPoint interface {
	Exemplars() pmetric.ExemplarSlice
	Attributes() pcommon.Map
	StartTimestamp() pcommon.Timestamp
	Timestamp() pcommon.Timestamp
	Flags() pmetric.DataPointFlags
}

/*
Auxiliary method to populate data from a data point representation. This is needed
to be able to lazy load all the dependant fields inside baseMetricSignal which depend
on datapoint data.

This method must be called after baseMetricSignal initialization for each data point.
*/
func loadDataPoint[T genericDataPoint](metric *baseMetricSignal, dp T) {
	_ = "STUB: not implemented"
	return
}

type exemplars struct {
	ExemplarsFilteredAttributes []map[string]string `json:"exemplars_filtered_attributes"`
	ExemplarsTimestamp          []string            `json:"exemplars_timestamp"`
	ExemplarsValue              []float64           `json:"exemplars_value"`
	ExemplarsSpanID             []string            `json:"exemplars_span_id"`
	ExemplarsTraceID            []string            `json:"exemplars_trace_id"`
}

type sumMetricSignal struct {
	baseMetricSignal
	Value                  float64 `json:"value"`
	AggregationTemporality int32   `json:"aggregation_temporality"`
	IsMonotonic            bool    `json:"is_monotonic"`
}

type gaugeMetricSignal struct {
	baseMetricSignal
	Value float64 `json:"value"`
}

type histogramMetricSignal struct {
	baseMetricSignal
	Count                  uint64    `json:"count"`
	Sum                    float64   `json:"sum"`
	BucketCounts           []uint64  `json:"bucket_counts"`
	ExplicitBounds         []float64 `json:"explicit_bounds"`
	Min                    *float64  `json:"min,omitempty"`
	Max                    *float64  `json:"max,omitempty"`
	AggregationTemporality int32     `json:"aggregation_temporality"`
}

type exponentialHistogramMetricSignal struct {
	baseMetricSignal
	Count                  uint64   `json:"count"`
	Sum                    float64  `json:"sum"`
	Scale                  int32    `json:"scale"`
	ZeroCount              uint64   `json:"zero_count"`
	PositiveOffset         int32    `json:"positive_offset"`
	PositiveBucketCounts   []uint64 `json:"positive_bucket_counts"`
	NegativeOffset         int32    `json:"negative_offset"`
	NegativeBucketCounts   []uint64 `json:"negative_bucket_counts"`
	Min                    *float64 `json:"min,omitempty"`
	Max                    *float64 `json:"max,omitempty"`
	AggregationTemporality int32    `json:"aggregation_temporality"`
}

func convertExemplars(exem pmetric.ExemplarSlice) exemplars {
	_ = "STUB: not implemented"
	return *new(exemplars)
}

// Value is unset, use 0.0 as default

func covertValue(dp pmetric.NumberDataPoint) float64 { _ = "STUB: not implemented"; return 0 }

func ConvertMetrics(md pmetric.Metrics, sumEncoder, gaugeEncoder, histogramEncoder, exponentialHistogramEncoder Encoder) error {
	_ = "STUB: not implemented"
	return nil
}
