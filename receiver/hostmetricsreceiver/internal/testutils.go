// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal"

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func AssertDescriptorEqual(t *testing.T, expected, actual pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

func AssertSumMetricHasAttributeValue(t *testing.T, metric pmetric.Metric, index int, labelName string, expectedVal pcommon.Value) {
	_ = "STUB: not implemented"
	return
}

func AssertSumMetricHasAttribute(t *testing.T, metric pmetric.Metric, index int, labelName string) {
	_ = "STUB: not implemented"
	return
}

func AssertSumMetricStartTimeEquals(t *testing.T, metric pmetric.Metric, startTime pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func AssertGaugeMetricHasAttributeValue(t *testing.T, metric pmetric.Metric, index int, labelName string, expectedVal pcommon.Value) {
	_ = "STUB: not implemented"
	return
}

func AssertGaugeMetricHasAttribute(t *testing.T, metric pmetric.Metric, index int, labelName string) {
	_ = "STUB: not implemented"
	return
}

func AssertGaugeMetricStartTimeEquals(t *testing.T, metric pmetric.Metric, startTime pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func AssertSameTimeStampForAllMetrics(t *testing.T, metrics pmetric.MetricSlice) {
	_ = "STUB: not implemented"
	return
}

func AssertSameTimeStampForMetrics(t *testing.T, metrics pmetric.MetricSlice, startIdx, endIdx int) {
	_ = "STUB: not implemented"
	return
}

// AssertExpectedMetrics checks that metrics contains expected metrics and no other metrics.
// It also checks that each expected metric has at least one data point.
func AssertExpectedMetrics(t *testing.T, expectedMetrics map[string]bool, actualMetrics pmetric.Metrics) {
	_ = "STUB: not implemented"
	return
}

// Check that we have some metrics and data points

// Check that each expected metric has at least one data point

// Check that no unexpected metrics were collected
