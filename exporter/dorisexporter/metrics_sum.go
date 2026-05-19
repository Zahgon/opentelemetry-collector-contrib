// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	_ "embed"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

//go:embed sql/metrics_sum_ddl.sql
var metricsSumDDL string

// dMetricSum Sum Metric to Doris
type dMetricSum struct {
	*dMetric               `json:",inline"`
	Timestamp              string         `json:"timestamp"`
	Attributes             map[string]any `json:"attributes"`
	StartTime              string         `json:"start_time"`
	Value                  float64        `json:"value"`
	Exemplars              []*dExemplar   `json:"exemplars"`
	AggregationTemporality string         `json:"aggregation_temporality"`
	IsMonotonic            bool           `json:"is_monotonic"`
}

type metricModelSum struct {
	metricModelCommon[dMetricSum]
}

func (*metricModelSum) metricType() pmetric.MetricType {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricType)
}

func (*metricModelSum) tableSuffix() string { _ = "STUB: not implemented"; return "" }

func (m *metricModelSum) add(pm pmetric.Metric, dm *dMetric, e *metricsExporter) error {
	_ = "STUB: not implemented"
	return nil
}
