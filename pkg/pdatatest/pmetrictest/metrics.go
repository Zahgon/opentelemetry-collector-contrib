// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetrictest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func CompareMetrics(expected, actual pmetric.Metrics, options ...CompareMetricsOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching resources so that each can only be matched once

func CompareResourceMetrics(expected, actual pmetric.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching resources so that each can only be matched once

// CompareScopeMetrics compares each part of two given ScopeMetrics and returns
// an error if they don't match. The error describes what didn't match. The
// expected and actual values are clones before options are applied.
func CompareScopeMetrics(expected, actual pmetric.ScopeMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching records so that each record can only be matched once

func CompareMetric(expected, actual pmetric.Metric) error { _ = "STUB: not implemented"; return nil }

//exhaustive:enforce

// compareNumberDataPointSlices compares each part of two given NumberDataPointSlices and returns
// an error if they don't match. The error describes what didn't match.
func compareNumberDataPointSlices(expected, actual pmetric.NumberDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching data points so that each point can only be matched once

// CompareNumberDataPoint compares each part of two given NumberDataPoints and returns
// an error if they don't match. The error describes what didn't match.
func CompareNumberDataPoint(expected, actual pmetric.NumberDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}

// compareExemplarSlice compares each part of two given ExemplarSlice and returns
// an error if they don't match. The error describes what didn't match.
func compareExemplarSlice(expected, actual pmetric.ExemplarSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching exemplars so that each exemplar can only be matched once

// CompareExemplar compares each part of two given pmetric.Exemplar and returns
// an error if they don't match. The error describes what didn't match.
func CompareExemplar(expected, actual pmetric.Exemplar) error {
	_ = "STUB: not implemented"
	return nil
}

// compareHistogramDataPointSlices compares each part of two given HistogramDataPointSlices and returns
// an error if they don't match. The error describes what didn't match.
func compareHistogramDataPointSlices(expected, actual pmetric.HistogramDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching data points so that each point can only be matched once

// CompareHistogramDataPoints compares each part of two given HistogramDataPoints and returns
// an error if they don't match. The error describes what didn't match.
func CompareHistogramDataPoints(expected, actual pmetric.HistogramDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}

// compareExponentialHistogramDataPointSlice compares each part of two given ExponentialHistogramDataPointSlices and
// returns an error if they don't match. The error describes what didn't match.
func compareExponentialHistogramDataPointSlice(expected, actual pmetric.ExponentialHistogramDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching data points so that each point can only be matched once

// CompareExponentialHistogramDataPoint compares each part of two given ExponentialHistogramDataPoints and returns
// an error if they don't match. The error describes what didn't match.
func CompareExponentialHistogramDataPoint(expected, actual pmetric.ExponentialHistogramDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}

// compareSummaryDataPointSlices compares each part of two given SummaryDataPoint slices and returns
// an error if they don't match. The error describes what didn't match.
func compareSummaryDataPointSlices(expected, actual pmetric.SummaryDataPointSlice) error {
	_ = "STUB: not implemented"
	return nil
}

// CompareSummaryDataPoint compares each part of two given SummaryDataPoint and returns
// an error if they don't match. The error describes what didn't match.
func CompareSummaryDataPoint(expected, actual pmetric.SummaryDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}
