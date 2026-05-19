// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	_ "embed"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

//go:embed sql/metrics_exponential_histogram_ddl.sql
var metricsExponentialHistogramDDL string

// dMetricExponentialHistogram Exponential Histogram Metric to Doris
type dMetricExponentialHistogram struct {
	*dMetric               `json:",inline"`
	Timestamp              string         `json:"timestamp"`
	Attributes             map[string]any `json:"attributes"`
	StartTime              string         `json:"start_time"`
	Count                  int64          `json:"count"`
	Sum                    float64        `json:"sum"`
	Scale                  int32          `json:"scale"`
	ZeroCount              int64          `json:"zero_count"`
	PositiveOffset         int32          `json:"positive_offset"`
	PositiveBucketCounts   []int64        `json:"positive_bucket_counts"`
	NegativeOffset         int32          `json:"negative_offset"`
	NegativeBucketCounts   []int64        `json:"negative_bucket_counts"`
	Exemplars              []*dExemplar   `json:"exemplars"`
	Min                    float64        `json:"min"`
	Max                    float64        `json:"max"`
	ZeroThreshold          float64        `json:"zero_threshold"`
	AggregationTemporality string         `json:"aggregation_temporality"`
}

type metricModelExponentialHistogram struct {
	metricModelCommon[dMetricExponentialHistogram]
}

func (*metricModelExponentialHistogram) metricType() pmetric.MetricType {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricType)
}

func (*metricModelExponentialHistogram) tableSuffix() string { _ = "STUB: not implemented"; return "" }

func (m *metricModelExponentialHistogram) add(pm pmetric.Metric, dm *dMetric, e *metricsExporter) error {
	_ = "STUB: not implemented"
	return nil
}
