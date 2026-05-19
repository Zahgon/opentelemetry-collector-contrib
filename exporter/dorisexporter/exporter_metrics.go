// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"context"
	_ "embed" // for SQL file embedding

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var ddls = []string{
	metricsGaugeDDL,
	metricsSumDDL,
	metricsHistogramDDL,
	metricsExponentialHistogramDDL,
	metricsSummaryDDL,
}

//go:embed sql/metrics_view.sql
var metricsView string

type metricsExporter struct {
	*commonExporter
}

func newMetricsExporter(logger *zap.Logger, cfg *Config, set component.TelemetrySettings) *metricsExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *metricsExporter) initMetricMap(ms pmetric.Metrics) map[pmetric.MetricType]metricModel {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) pushMetricData(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) pushMetricDataParallel(ctx context.Context, metricMap map[pmetric.MetricType]metricModel) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) pushMetricDataInternal(ctx context.Context, metrics metricModel) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricsExporter) getNumberDataPointValue(dp pmetric.NumberDataPoint) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *metricsExporter) getExemplarValue(ep pmetric.Exemplar) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *metricsExporter) generateMetricLabel(m metricModel) string {
	_ = "STUB: not implemented"
	return ""
}
