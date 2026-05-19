// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	_ "embed"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

//go:embed sql/metrics_gauge_ddl.sql
var metricsGaugeDDL string

// dMetricGauge Gauge Metric to Doris
type dMetricGauge struct {
	*dMetric   `json:",inline"`
	Timestamp  string         `json:"timestamp"`
	Attributes map[string]any `json:"attributes"`
	StartTime  string         `json:"start_time"`
	Value      float64        `json:"value"`
	Exemplars  []*dExemplar   `json:"exemplars"`
}

type metricModelGauge struct {
	metricModelCommon[dMetricGauge]
}

func (*metricModelGauge) metricType() pmetric.MetricType {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricType)
}

func (*metricModelGauge) tableSuffix() string { _ = "STUB: not implemented"; return "" }

func (m *metricModelGauge) add(pm pmetric.Metric, dm *dMetric, e *metricsExporter) error {
	_ = "STUB: not implemented"
	return nil
}
