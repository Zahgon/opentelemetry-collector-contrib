// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetricutiltest // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector/internal/pmetricutiltest"

import "go.opentelemetry.io/collector/pdata/pmetric"

// NewGauges returns a pmetric.Metrics with a uniform structure where resources, scopes, metrics,
// and datapoints are identical across all instances, except for one identifying field.
//
// Identifying fields:
// - Resources have an attribute called "resourceName" with a value of "resourceN".
// - Scopes have a name with a value of "scopeN".
// - Metrics have a name with a value of "metricN" and a single time series of data points.
// - DataPoints have an attribute "dpName" with a value of "dpN".
//
// Example: NewGauges("AB", "XYZ", "MN", "1234") returns:
//
//	resourceA, resourceB
//	    each with scopeX, scopeY, scopeZ
//	        each with metricM, metricN
//	            each with dp1, dp2, dp3, dp4
//
// Each byte in the input string is a unique ID for the corresponding element.
func NewGauges(resourceIDs, scopeIDs, metricIDs, dataPointIDs string) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func NewSums(resourceIDs, scopeIDs, metricIDs, dataPointIDs string, isMonotonic bool, aggregationTemporality pmetric.AggregationTemporality) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func NewHistograms(resourceIDs, scopeIDs, metricIDs, dataPointIDs string, aggregationTemporality pmetric.AggregationTemporality) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func NewExponentialHistograms(resourceIDs, scopeIDs, metricIDs, dataPointIDs string, aggregationTemporality pmetric.AggregationTemporality) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func NewSummaries(resourceIDs, scopeIDs, metricIDs, dataPointIDs string) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func NewMetricsFromOpts(resources ...pmetric.ResourceMetrics) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func Resource(id string, scopes ...pmetric.ScopeMetrics) pmetric.ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ResourceMetrics)
}

func Scope(id string, metrics ...pmetric.Metric) pmetric.ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ScopeMetrics)
}

func Gauge(id string, dps ...pmetric.NumberDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func Sum(id string, isMonotonic bool, aggregationTemporality pmetric.AggregationTemporality, dps ...pmetric.NumberDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func NumberDataPoint(id string) pmetric.NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPoint)
}

func Histogram(id string, aggregationTemporality pmetric.AggregationTemporality, dps ...pmetric.HistogramDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func HistogramDataPoint(id string) pmetric.HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.HistogramDataPoint)
}

func ExponentialHistogram(id string, aggregationTemporality pmetric.AggregationTemporality, dps ...pmetric.ExponentialHistogramDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func ExponentialHistogramDataPoint(id string) pmetric.ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.ExponentialHistogramDataPoint)
}

func Summary(id string, dps ...pmetric.SummaryDataPoint) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func SummaryDataPoint(id string) pmetric.SummaryDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.SummaryDataPoint)
}
