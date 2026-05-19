// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricsaslogsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/metricsaslogsconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

const (
	attrMetricName                   = "metric.name"
	attrMetricType                   = "metric.type"
	attrMetricDescription            = "metric.description"
	attrMetricUnit                   = "metric.unit"
	attrMetricIsMonotonic            = "metric.is_monotonic"
	attrMetricAggregationTemporality = "metric.aggregation_temporality"

	attrGaugeValue = "gauge.value"
	attrSumValue   = "sum.value"

	attrHistogramCount          = "histogram.count"
	attrHistogramSum            = "histogram.sum"
	attrHistogramMin            = "histogram.min"
	attrHistogramMax            = "histogram.max"
	attrHistogramBucketCounts   = "histogram.bucket_counts"
	attrHistogramExplicitBounds = "histogram.explicit_bounds"

	attrExponentialHistogramCount     = "exponential_histogram.count"
	attrExponentialHistogramSum       = "exponential_histogram.sum"
	attrExponentialHistogramScale     = "exponential_histogram.scale"
	attrExponentialHistogramZeroCount = "exponential_histogram.zero_count"
	attrExponentialHistogramMin       = "exponential_histogram.min"
	attrExponentialHistogramMax       = "exponential_histogram.max"

	attrSummaryCount          = "summary.count"
	attrSummarySum            = "summary.sum"
	attrSummaryQuantileValues = "summary.quantile_values"
	attrQuantile              = "quantile"
	attrValue                 = "value"
)

type metricsAsLogs struct {
	logsConsumer consumer.Logs
	config       *Config
	logger       *zap.Logger
	component.StartFunc
	component.ShutdownFunc
}

func (*metricsAsLogs) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (m *metricsAsLogs) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *metricsAsLogs) processResourceMetrics(resourceMetric pmetric.ResourceMetrics, logs plog.Logs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) processScopeMetrics(scopeMetric pmetric.ScopeMetrics, resourceLogs plog.ResourceLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) processMetric(metric pmetric.Metric, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) processGaugeDataPointsWithMetric(metric pmetric.Metric, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) processSumDataPointsWithMetric(metric pmetric.Metric, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) processHistogramDataPointsWithMetric(metric pmetric.Metric, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) processExponentialHistogramDataPointsWithMetric(metric pmetric.Metric, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) processSummaryDataPointsWithMetric(metric pmetric.Metric, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) convertGaugeDataPointToLogRecordWithMetric(metric pmetric.Metric, dataPoint pmetric.NumberDataPoint, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) convertSumDataPointToLogRecordWithMetric(metric pmetric.Metric, sum pmetric.Sum, dataPoint pmetric.NumberDataPoint, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) convertHistogramDataPointToLogRecordWithMetric(metric pmetric.Metric, histogram pmetric.Histogram, dataPoint pmetric.HistogramDataPoint, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) convertExponentialHistogramDataPointToLogRecordWithMetric(metric pmetric.Metric, expHistogram pmetric.ExponentialHistogram, dataPoint pmetric.ExponentialHistogramDataPoint, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (m *metricsAsLogs) convertSummaryDataPointToLogRecordWithMetric(metric pmetric.Metric, dataPoint pmetric.SummaryDataPoint, scopeLogs *plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return
}

func (*metricsAsLogs) addAggregationTemporalityAttribute(logRecord plog.LogRecord, temporality pmetric.AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}

func (*metricsAsLogs) addCommonMetricAttributes(logRecord plog.LogRecord, metric pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

func (*metricsAsLogs) addNumberDataPointAttributes(logRecord plog.LogRecord, dataPoint pmetric.NumberDataPoint, valueAttr string) {
	_ = "STUB: not implemented"
	return
}

func (*metricsAsLogs) addHistogramDataPointAttributes(logRecord plog.LogRecord, dataPoint pmetric.HistogramDataPoint) {
	_ = "STUB: not implemented"
	return
}

func (*metricsAsLogs) addExponentialHistogramDataPointAttributes(logRecord plog.LogRecord, dataPoint pmetric.ExponentialHistogramDataPoint) {
	_ = "STUB: not implemented"
	return
}

func (*metricsAsLogs) addSummaryDataPointAttributes(logRecord plog.LogRecord, dataPoint pmetric.SummaryDataPoint) {
	_ = "STUB: not implemented"
	return
}

func (*metricsAsLogs) setLogRecordFromDataPoint(logRecord plog.LogRecord, metricName, metricType string, attributes pcommon.Map, timestamp, startTimestamp pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// Copy datapoint attributes first, before adding metric-specific attributes
