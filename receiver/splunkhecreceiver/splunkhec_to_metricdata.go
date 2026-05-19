// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	translator "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/splunk"
)

// splunkHecToMetricsData converts Splunk HEC metric points to
// pmetric.Metrics. Returning the converted data and the number of
// dropped time series.
func splunkHecToMetricsData(logger *zap.Logger, events []*translator.Event, resourceCustomizer func(pcommon.Resource), config *Config) (pmetric.Metrics, int) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), 0
}

// drop this point as we do not know how to extract a value from it

func convertString(logger *zap.Logger, numDroppedTimeSeries *int, metrics pmetric.MetricSlice, metricName string, pointTimestamp pcommon.Timestamp, s string, attributes pcommon.Map) {
	_ = "STUB: not implemented"
	// best effort, cast to string and turn into a number
	return
}

func addIntGauge(metrics pmetric.MetricSlice, metricName string, value int64, ts pcommon.Timestamp, attributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func addDoubleGauge(metrics pmetric.MetricSlice, metricName string, value float64, ts pcommon.Timestamp, attributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

/*
	Splunk HEC timestamps can be in nanoseconds, microseconds, milliseconds, and seconds epoch.
	Example:
		- 1234567890
		- 1234567890123
		- 1234567890123456
		- 1234567890123456789

	The format can also be <second>.<sub-second>.
	Example:
		- 1234567890.000
		- 1234567890.123
		- 1234567890.123456
		- 1234567890.123456789
*/

func convertTimestamp(t float64) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

// nano

// micro

// milli

// second

// Extract dimensions from the Splunk event fields to populate metric data point attributes.
func buildAttributes(dimensions map[string]any) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// TODO: Log or metric for this odd ball?
