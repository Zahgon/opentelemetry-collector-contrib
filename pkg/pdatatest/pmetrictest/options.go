// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetrictest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"

import (
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// CompareMetricsOption can be used to mutate expected and/or actual metrics before comparing.
type CompareMetricsOption interface {
	applyOnMetrics(expected, actual pmetric.Metrics)
}

type compareMetricsOptionFunc func(expected, actual pmetric.Metrics)

func (f compareMetricsOptionFunc) applyOnMetrics(expected, actual pmetric.Metrics) {
	_ = "STUB: not implemented"
	return

	// IgnoreMetricValues is a CompareMetricsOption that clears all metric values.
}

func IgnoreMetricValues(metricNames ...string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskMetricValues(metrics pmetric.Metrics, metricNames ...string) {
	_ = "STUB: not implemented"
	return
}

// maskMetricSliceValues sets all data point values to zero.
func maskMetricSliceValues(metrics pmetric.MetricSlice, metricNames ...string) {
	_ = "STUB: not implemented"
	return
}

func getDataPointSlice(metric pmetric.Metric) pmetric.NumberDataPointSlice {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPointSlice)
}

//exhaustive:enforce

// maskDataPointSliceValues sets all data point values to zero.
func maskDataPointSliceValues(dataPoints pmetric.NumberDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

// maskHistogramDataPointSliceValues sets all data point values to zero.
func maskHistogramDataPointSliceValues(dataPoints pmetric.HistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

// IgnoreMetricFloatPrecision is a CompareMetricsOption that rounds away float precision discrepancies in metric values.
func IgnoreMetricFloatPrecision(precision int, metricNames ...string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func floatMetricValues(precision int, metrics pmetric.Metrics, metricNames ...string) {
	_ = "STUB: not implemented"
	return
}

// floatMetricSliceValues sets all data point values to zero.
func floatMetricSliceValues(precision int, metrics pmetric.MetricSlice, metricNames ...string) {
	_ = "STUB: not implemented"
	return
}

// maskDataPointSliceValues rounds all data point values at a given decimal.
func roundDataPointSliceValues(dataPoints pmetric.NumberDataPointSlice, precision int) {
	_ = "STUB: not implemented"
	return
}

// IgnoreExemplars is a CompareMetricsOption that clears exemplar fields on all data points.
func IgnoreExemplars() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskExemplars(metrics pmetric.Metrics) { _ = "STUB: not implemented"; return }

// IgnoreExemplarSlice is a CompareMetricsOption that clears exemplars slice on all data points.
func IgnoreExemplarSlice() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskExemplarSlice(metrics pmetric.Metrics) { _ = "STUB: not implemented"; return }

// IgnoreTimestamp is a CompareMetricsOption that clears Timestamp fields on all the data points.
func IgnoreTimestamp() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskTimestamp(metrics pmetric.Metrics, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

//exhaustive:enforce

// IgnoreStartTimestamp is a CompareMetricsOption that clears StartTimestamp fields on all the data points.
func IgnoreStartTimestamp() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskStartTimestamp(metrics pmetric.Metrics, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

//exhaustive:enforce

// IgnoreMetricAttributeValue is a CompareMetricsOption that clears value of the metric attribute.
func IgnoreMetricAttributeValue(attributeName string, metricNames ...string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

// IgnoreDatapointAttributesOrder is a CompareMetricsOption that ignores the order of datapoint attributes.
func IgnoreDatapointAttributesOrder() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func orderDatapointAttributes(metrics pmetric.Metrics) { _ = "STUB: not implemented"; return }

func maskMetricAttributeValue(metrics pmetric.Metrics, attributeName string, metricNames []string) {
	_ = "STUB: not implemented"
	return
}

// maskMetricSliceAttributeValues sets the value of the specified attribute to
// the zero value associated with the attribute data type.
// If metric names are specified, only the data points within those metrics will be masked.
// Otherwise, all data points with the attribute will be masked.
func maskMetricSliceAttributeValues(metrics pmetric.MetricSlice, attributeName string, metricNames []string) {
	_ = "STUB: not implemented"
	return
}

// If attribute values are ignored, some data points may become
// indistinguishable from each other, but sorting by value allows
// for a reasonably thorough comparison and a deterministic outcome.

// If attribute values are ignored, some data points may become
// indistinguishable from each other, but sorting by value allows
// for a reasonably thorough comparison and a deterministic outcome.

// maskDataPointSliceAttributeValues sets the value of the specified attribute to
// the zero value associated with the attribute data type.
func maskDataPointSliceAttributeValues(dataPoints pmetric.NumberDataPointSlice, attributeName string) {
	_ = "STUB: not implemented"
	return
}

// maskHistogramSliceAttributeValues sets the value of the specified attribute to
// the zero value associated with the attribute data type.
func maskHistogramSliceAttributeValues(dataPoints pmetric.HistogramDataPointSlice, attributeName string) {
	_ = "STUB: not implemented"
	return
}

// MatchMetricAttributeValue is a CompareMetricsOption that transforms a metric attribute value based on a regular expression.
func MatchMetricAttributeValue(attributeName, pattern string, metricNames ...string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func matchMetricAttributeValue(metrics pmetric.Metrics, attributeName string, re *regexp.Regexp, metricNames []string) {
	_ = "STUB: not implemented"
	return
}

func matchMetricSliceAttributeValues(metrics pmetric.MetricSlice, attributeName string, re *regexp.Regexp, metricNames []string) {
	_ = "STUB: not implemented"
	return
}

// If attribute values are ignored, some data points may become
// indistinguishable from each other, but sorting by value allows
// for a reasonably thorough comparison and a deterministic outcome.

// If attribute values are ignored, some data points may become
// indistinguishable from each other, but sorting by value allows
// for a reasonably thorough comparison and a deterministic outcome.

func matchDataPointSliceAttributeValues(dataPoints pmetric.NumberDataPointSlice, attributeName string, re *regexp.Regexp) {
	_ = "STUB: not implemented"
	return
}

func matchHistogramDataPointSliceAttributeValues(dataPoints pmetric.HistogramDataPointSlice, attributeName string, re *regexp.Regexp) {
	_ = "STUB: not implemented"
	return
}

// MatchResourceAttributeValue is a CompareMetricsOption that transforms a resource attribute value based on a regular expression.
func MatchResourceAttributeValue(attributeName, pattern string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func matchResourceAttributeValue(metrics pmetric.Metrics, attributeName string, re *regexp.Regexp) {
	_ = "STUB: not implemented"
	return
}

// IgnoreResourceAttributeValue is a CompareMetricsOption that removes a resource attribute
// from all resources.
func IgnoreResourceAttributeValue(attributeName string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskMetricsResourceAttributeValue(metrics pmetric.Metrics, attributeName string) {
	_ = "STUB: not implemented"
	return
}

// IgnoreResourceEntityRefs is a CompareMetricsOption that clears entity references
// on all resources.
func IgnoreResourceEntityRefs() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskMetricsResourceEntityRefs(metrics pmetric.Metrics) { _ = "STUB: not implemented"; return }

func ChangeResourceAttributeValue(attributeName string, changeFn func(string) string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func changeMetricsResourceAttributeValue(metrics pmetric.Metrics, attributeName string, changeFn func(string) string) {
	_ = "STUB: not implemented"
	return
}

// ChangeDatapointAttributeValue changes the metric datapoint value with the specified key
func ChangeDatapointAttributeValue(attributeName string, changeFn func(string) string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func changeMetricsDatapointAttributeValue(metrics pmetric.Metrics, attributeName string, changeFn func(string) string) {
	_ = "STUB: not implemented"
	return
}

// IgnoreSubsequentDataPoints is a CompareMetricsOption that ignores data points after the first.
func IgnoreSubsequentDataPoints(metricNames ...string) CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskSubsequentDataPoints(metrics pmetric.Metrics, metricNames []string) {
	_ = "STUB: not implemented"
	return
}

// IgnoreResourceMetricsOrder is a CompareMetricsOption that ignores the order of resource traces/metrics/logs.
func IgnoreResourceMetricsOrder() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func sortResourceMetricsSlice(rms pmetric.ResourceMetricsSlice) { _ = "STUB: not implemented"; return }

func IgnoreScopeVersion() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func maskScopeVersion(metrics pmetric.Metrics) { _ = "STUB: not implemented"; return }

// IgnoreScopeMetricsOrder is a CompareMetricsOption that ignores the order of instrumentation scope traces/metrics/logs.
func IgnoreScopeMetricsOrder() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func sortScopeMetricsSlices(ms pmetric.Metrics) { _ = "STUB: not implemented"; return }

// IgnoreMetricsOrder is a CompareMetricsOption that ignores the order of metrics.
func IgnoreMetricsOrder() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func sortMetricSlices(ms pmetric.Metrics) { _ = "STUB: not implemented"; return }

// IgnoreMetricDataPointsOrder is a CompareMetricsOption that ignores the order of metrics.
func IgnoreMetricDataPointsOrder() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func sortMetricDataPointSlices(ms pmetric.Metrics) { _ = "STUB: not implemented"; return }

//exhaustive:enforce

func sortNumberDataPointSlice(ndps pmetric.NumberDataPointSlice) { _ = "STUB: not implemented"; return }

func sortHistogramDataPointSlice(hdps pmetric.HistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func sortExponentialHistogramDataPointSlice(hdps pmetric.ExponentialHistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func sortSummaryDataPointSlice(sds pmetric.SummaryDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

// IgnoreSummaryDataPointValueAtQuantileSliceOrder is a CompareMetricsOption that ignores the order of summary data point quantile slice.
func IgnoreSummaryDataPointValueAtQuantileSliceOrder() CompareMetricsOption {
	_ = "STUB: not implemented"
	return *new(CompareMetricsOption)
}

func sortSummaryDataPointValueAtQuantileSlices(ms pmetric.Metrics) {
	_ = "STUB: not implemented"
	return
}
