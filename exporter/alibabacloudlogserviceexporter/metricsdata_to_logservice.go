// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

import (
	"strings"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

const (
	metricNameKey = "__name__"
	labelsKey     = "__labels__"
	timeNanoKey   = "__time_nano__"
	valueKey      = "__value__"
	// same with : https://github.com/prometheus/common/blob/b5fe7d854c42dc7842e48d1ca58f60feae09d77b/expfmt/text_create.go#L445
	infinityBoundValue = "+Inf"
	bucketLabelKey     = "le"
	summaryLabelKey    = "quantile"
)

type keyValue struct {
	Key   string
	Value string
}

type keyValues struct {
	keyValues []keyValue
}

func (kv *keyValues) Len() int      { _ = "STUB: not implemented"; return 0 }
func (kv *keyValues) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (kv *keyValues) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (kv *keyValues) Sort() { _ = "STUB: not implemented"; return }

func (kv *keyValues) Replace(key, value string) { _ = "STUB: not implemented"; return }

func (kv *keyValues) Append(key, value string) { _ = "STUB: not implemented"; return }

func (kv *keyValues) Clone() keyValues { _ = "STUB: not implemented"; return *new(keyValues) }

func (kv *keyValues) String() string { _ = "STUB: not implemented"; return "" }

func (kv *keyValues) labelToStringBuilder(sb *strings.Builder) { _ = "STUB: not implemented"; return }

func formatMetricName(name string) string { _ = "STUB: not implemented"; return "" }

func newMetricLogFromRaw(
	name string,
	labels keyValues,
	nsec int64,
	value float64,
) *sls.Log {
	_ = "STUB: not implemented"
	return nil
}

func resourceToMetricLabels(labels *keyValues, resource pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

func numberMetricsToLogs(name string, data pmetric.NumberDataPointSlice, defaultLabels keyValues) (logs []*sls.Log) {
	_ = "STUB: not implemented"
	return nil
}

func doubleHistogramMetricsToLogs(name string, data pmetric.HistogramDataPointSlice, defaultLabels keyValues) (logs []*sls.Log) {
	_ = "STUB: not implemented"
	return nil
}

func doubleSummaryMetricsToLogs(name string, data pmetric.SummaryDataPointSlice, defaultLabels keyValues) (logs []*sls.Log) {
	_ = "STUB: not implemented"
	return nil
}

// Adding the "quantile" dimension.

func metricDataToLogServiceData(md pmetric.Metric, defaultLabels keyValues) (logs []*sls.Log) {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return nil
}

func metricsDataToLogServiceData(
	_ *zap.Logger,
	md pmetric.Metrics,
) (logs []*sls.Log) {
	_ = "STUB: not implemented"
	return nil
}

// ignore insMetrics.Scope()
