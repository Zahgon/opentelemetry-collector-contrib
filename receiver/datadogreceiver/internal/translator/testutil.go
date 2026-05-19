// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const aggregationTemporality = pmetric.AggregationTemporalityDelta

func createMetricsTranslator() *MetricsTranslator { _ = "STUB: not implemented"; return nil }

func requireResourceAttributes(t *testing.T, attrs, expectedAttrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

//nolint:unparam
func requireScopeMetrics(t *testing.T, result pmetric.Metrics, expectedScopeMetricsLen, expectedMetricsLen int) {
	_ = "STUB: not implemented"
	return
}

func requireScope(t *testing.T, result pmetric.Metrics, expectedAttrs pcommon.Map, expectedVersion string) {
	_ = "STUB: not implemented"
	return
}

func requireMetricAndDataPointCounts(t *testing.T, result pmetric.Metrics, expectedMetricCount, expectedDpCount int) {
	_ = "STUB: not implemented"
	return
}

func requireSum(t *testing.T, metric pmetric.Metric, expectedName string, expectedDpsLen int) {
	_ = "STUB: not implemented"
	return
}

func requireGauge(t *testing.T, metric pmetric.Metric, expectedName string, expectedDpsLen int) {
	_ = "STUB: not implemented"
	return
}

func requireDp(t *testing.T, dp pmetric.NumberDataPoint, expectedAttrs pcommon.Map, expectedTime int64, expectedValue float64) {
	_ = "STUB: not implemented"
	return
}

func totalHistBucketCounts(hist pmetric.ExponentialHistogramDataPoint) uint64 {
	_ = "STUB: not implemented"
	return 0
}
