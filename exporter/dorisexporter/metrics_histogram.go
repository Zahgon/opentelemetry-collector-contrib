// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	_ "embed"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

//go:embed sql/metrics_histogram_ddl.sql
var metricsHistogramDDL string

// dMetricHistogram Histogram Metric to Doris
type dMetricHistogram struct {
	*dMetric               `json:",inline"`
	Timestamp              string         `json:"timestamp"`
	Attributes             map[string]any `json:"attributes"`
	StartTime              string         `json:"start_time"`
	Count                  int64          `json:"count"`
	Sum                    float64        `json:"sum"`
	BucketCounts           []int64        `json:"bucket_counts"`
	ExplicitBounds         []float64      `json:"explicit_bounds"`
	Exemplars              []*dExemplar   `json:"exemplars"`
	Min                    float64        `json:"min"`
	Max                    float64        `json:"max"`
	AggregationTemporality string         `json:"aggregation_temporality"`
}

type metricModelHistogram struct {
	metricModelCommon[dMetricHistogram]
}

func (*metricModelHistogram) metricType() pmetric.MetricType {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricType)
}

func (*metricModelHistogram) tableSuffix() string { _ = "STUB: not implemented"; return "" }

func (m *metricModelHistogram) add(pm pmetric.Metric, dm *dMetric, e *metricsExporter) error {
	_ = "STUB: not implemented"
	return nil
}
