// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal"

import (
	"errors"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/labels"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const (
	metricsSuffixCount  = "_count"
	metricsSuffixBucket = "_bucket"
	metricsSuffixSum    = "_sum"
	metricSuffixTotal   = "_total"
	metricSuffixInfo    = "_info"
	metricSuffixCreated = "_created"
	startTimeMetricName = "process_start_time_seconds"
	scrapeUpMetricName  = "up"

	transport  = "http"
	dataformat = "prometheus"
)

var (
	trimmableSuffixes     = []string{metricsSuffixBucket, metricsSuffixCount, metricsSuffixSum, metricSuffixTotal, metricSuffixInfo, metricSuffixCreated}
	errNoDataToBuild      = errors.New("there's no data to build")
	errNoBoundaryLabel    = errors.New("given metricType has no 'le' or 'quantile' label")
	errEmptyQuantileLabel = errors.New("'quantile' label on summary metric is missing or empty")
	errEmptyLeLabel       = errors.New("'le' label on histogram metric is missing or empty")
	errMetricNameNotFound = errors.New("metricName not found from labels")
	errTransactionAborted = errors.New("transaction aborted")
	errNoJobInstance      = errors.New("job or instance cannot be found from labels")

	notUsefulLabelsOther = sortString([]string{
		model.MetricNameLabel, model.InstanceLabel, model.SchemeLabel,
		model.MetricsPathLabel, model.JobLabel,
	})
	notUsefulLabelsHistogram = sortString(append(notUsefulLabelsOther, model.BucketLabel))
	notUsefulLabelsSummary   = sortString(append(notUsefulLabelsOther, model.QuantileLabel))
)

func sortString(strs []string) []string { _ = "STUB: not implemented"; return nil }

func getSortedNotUsefulLabels(mType pmetric.MetricType) []string {
	_ = "STUB: not implemented"
	return nil
}

func getSortedNotUsefulLabelsForSeries(mType pmetric.MetricType, ls labels.Labels) []string {
	_ = "STUB: not implemented"
	return nil
}

func timestampFromFloat64(ts float64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func timestampFromMs(timeAtMs int64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func getBoundary(metricType pmetric.MetricType, labels labels.Labels) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// convToMetricType returns the data type and if it is monotonic
func convToMetricType(metricType model.MetricType, exponentialHistogram bool) (pmetric.MetricType, bool) {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricType), false
}

// always use float64, as it's the internal data type used in prometheus

// model.MetricTypeUnknown is converted to gauge by default to prevent Prometheus untyped metrics from being dropped

// dropping support for gaugehistogram for now until we have an official spec of its implementation
// a draft can be found in: https://docs.google.com/document/d/1KwV0mAXwwbvvifBvDKH_LU1YjyXE_wxCkHNoCGq1GX0/edit#heading=h.1cvzqd4ksd23
// case model.MetricTypeGaugeHistogram:
//	return <pdata gauge histogram type>

// including: model.MetricTypeGaugeHistogram

func normalizeMetricName(name string) string { _ = "STUB: not implemented"; return "" }
