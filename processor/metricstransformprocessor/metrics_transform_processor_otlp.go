// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/aggregateutil"
)

// extractAndRemoveMatchedMetrics extracts matched metrics from ms metric slice and returns a new slice.
// Extracted metrics can have reduced number of data point if not all of them match the filter.
// All matched metrics, including metrics with only a subset of matched data points,
// are removed from the original ms metric slice.
func extractAndRemoveMatchedMetrics(dest pmetric.MetricSlice, f internalFilter, ms pmetric.MetricSlice) {
	_ = "STUB: not implemented"
	return
}

// matchMetrics returns a slice of metrics matching the filter f. Original metrics slice is not affected.
func matchMetrics(f internalFilter, metrics pmetric.MetricSlice) []pmetric.Metric {
	_ = "STUB: not implemented"
	return nil
}

// extractMatchedMetric returns a metric matching the filter.
// If provided metric matches the filter with all its data points, the original metric returned as is.
// If only part of data points match the filter, a new metric is returned with data points matching the filter.
// Otherwise, an invalid metric is returned.
func (f internalFilterStrict) extractMatchedMetric(metric pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func (f internalFilterStrict) matchMetric(metric pmetric.Metric) bool {
	_ = "STUB: not implemented"
	return false
}

func (internalFilterStrict) submatches(pmetric.Metric) []int { _ = "STUB: not implemented"; return nil }

func (internalFilterStrict) expand(string, string) string { _ = "STUB: not implemented"; return "" }

func (f internalFilterStrict) matchAttrs(attrs pcommon.Map) bool {
	_ = "STUB: not implemented"
	return false
}

// extractMatchedMetric returns a metric matching the filter.
// If provided metric matches the filter with all its data points, the original metric returned as is.
// If only part of data points match the filter, a new metric is returned with data points matching the filter.
// Otherwise, an invalid metric is returned.
func (f internalFilterRegexp) extractMatchedMetric(metric pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func (f internalFilterRegexp) matchMetric(metric pmetric.Metric) bool {
	_ = "STUB: not implemented"
	return false
}

func (f internalFilterRegexp) submatches(metric pmetric.Metric) []int {
	_ = "STUB: not implemented"
	return nil
}

func (f internalFilterRegexp) expand(metricTemplate, metricName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (f internalFilterRegexp) matchAttrs(attrs pcommon.Map) bool {
	_ = "STUB: not implemented"
	return false
}

// matchAnyDps checks whether any metric data points match the filter, returns true if metric has no data points.
func matchAnyDps(metric pmetric.Metric, f internalFilter) bool {
	_ = "STUB: not implemented"
	return false
}

// matchAllDps checks whether all metric data points match the filter, returns true if metric has no data points.
func matchAllDps(metric pmetric.Metric, f internalFilter) bool {
	_ = "STUB: not implemented"
	return false
}

// matchDps returns a slice of bool values representing data points matches following by the total number of matched
// data points.
// For example, for a metric with 3 data points where only first and third match the filter, the output will be:
// ([]bool{true, false, true}, 2).
func matchDps(metric pmetric.Metric, f internalFilter) (matchedDps []bool, matchedDpsCount int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// extractMetricWithMatchingAttrs returns a metric with data points matching attrMatchers.
// New metric is returned if part of data points match the filter,
// original metric returned if all data points match the filter,
// and invalid metric returned if no data points match the filter.
func extractMetricWithMatchingAttrs(metric pmetric.Metric, f internalFilter) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

//exhaustive:enforce

func matchAttrs(attrMatchers map[string]StringMatcher, attrs pcommon.Map) bool {
	_ = "STUB: not implemented"
	return false
}

// attribute values doesn't match, drop datapoint

// if a label-key is not found then return nil only if the given label-value is non-empty. If a given label-value is empty
// and the key is not found then move forward. In this approach we can make sure certain key is not present which is a valid use case.

func (mtp *metricsTransformProcessor) processMetrics(_ context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// TODO: report via trace / metric instead

// Save len, so we don't iterate over the newly generated metrics that are appended at the end.

// Drop the metric if all the data points were dropped after transformations.

func initResourceMetrics(dest pmetric.ResourceMetrics, resource pcommon.Resource, scope pcommon.InstrumentationScope, transform internalTransform) {
	_ = "STUB: not implemented"
	return
}

// canBeCombined returns true if all the provided metrics share the same type, unit, and labels
func canBeCombined(metrics []pmetric.Metric) error { _ = "STUB: not implemented"; return nil }

func metricAttributeKeys(metric pmetric.Metric) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// combine combines the metrics based on the supplied filter.
// canBeCombined must be called before.
func combine(transform internalTransform, metrics pmetric.MetricSlice) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// create combined metric with relevant name & descriptor

// append attribute keys based on the transform filter's named capturing groups

// if the subexpression is not named, use regexp notation, e.g. $1

// append attr values based on regex submatches

// groupMetrics groups all the provided timeseries that will be aggregated together based on all the label values.
// Returns a map of grouped timeseries and the corresponding selected labels
// canBeCombined must be called before.
func groupMetrics(metrics pmetric.MetricSlice, aggType aggregateutil.AggregationType, to pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

func copyMetricDetails(from, to pmetric.Metric) { _ = "STUB: not implemented"; return }

//exhaustive:enforce

// rangeDataPointAttributes calls f sequentially on attributes of every metric data point.
// The iteration terminates if f returns false.
func rangeDataPointAttributes(metric pmetric.Metric, f func(pcommon.Map) bool) {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return
}

func countDataPoints(metric pmetric.Metric) int {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return 0
}

// transformMetric updates the metric content based on operations indicated in transform and returns a flag
// specifying whether the metric is valid after applying the translations,
// e.g. false is returned if all the data points were removed after applying the translations.
func transformMetric(metric pmetric.Metric, transform internalTransform) bool {
	_ = "STUB: not implemented"
	return false
}

// Consider metric invalid if all its data points were removed after applying the operations.
