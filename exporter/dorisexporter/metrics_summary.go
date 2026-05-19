// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	_ "embed"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

//go:embed sql/metrics_summary_ddl.sql
var metricsSummaryDDL string

// dMetricSummary Summary metric model to Doris
type dMetricSummary struct {
	*dMetric       `json:",inline"`
	Timestamp      string            `json:"timestamp"`
	Attributes     map[string]any    `json:"attributes"`
	StartTime      string            `json:"start_time"`
	Count          int64             `json:"count"`
	Sum            float64           `json:"sum"`
	QuantileValues []*dQuantileValue `json:"quantile_values"`
}

// dQuantileValue Quantile Value to Doris
type dQuantileValue struct {
	Quantile float64 `json:"quantile"`
	Value    float64 `json:"value"`
}

type metricModelSummary struct {
	metricModelCommon[dMetricSummary]
}

func (*metricModelSummary) metricType() pmetric.MetricType {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricType)
}

func (*metricModelSummary) tableSuffix() string { _ = "STUB: not implemented"; return "" }

func (m *metricModelSummary) add(pm pmetric.Metric, dm *dMetric, e *metricsExporter) error {
	_ = "STUB: not implemented"
	return nil
}
