// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsecscontainermetrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver/internal/awsecscontainermetrics"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func convertToOTLPMetrics(prefix string, m ECSMetrics, r pcommon.Resource, timestamp pcommon.Timestamp) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// Ephemeral storage metrics are only available at the task level.
// They represent the shared ephemeral storage for the entire Fargate task.

func convertStoppedContainerDataToOTMetrics(prefix string, containerResource pcommon.Resource, timestamp pcommon.Timestamp, duration float64) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func appendIntGauge(metricName, unit string, value int64, ts pcommon.Timestamp, ilm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func appendIntSum(metricName, unit string, value int64, ts pcommon.Timestamp, ilm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func appendDoubleGauge(metricName, unit string, value float64, ts pcommon.Timestamp, ilm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func appendIntDataPoint(dataPoints pmetric.NumberDataPointSlice, value int64, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func appendMetric(ilm pmetric.ScopeMetrics, name, unit string) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}
