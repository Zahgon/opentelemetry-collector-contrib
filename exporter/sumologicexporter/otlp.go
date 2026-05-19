// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// decomposeHistograms decomposes any histograms present in the metric data into individual Sums and Gauges
// This is a noop if no Histograms are present, but otherwise makes a copy of the whole structure
// This exists because Sumo doesn't support OTLP histograms yet, and has the same semantics as the conversion to Prometheus format in prometheus_formatter.go
func decomposeHistograms(md pmetric.Metrics) pmetric.Metrics {
	_ = "STUB: not implemented"
	// short circuit and do nothing if no Histograms are present
	return *new(pmetric.Metrics)
}

// decomposeHistogram decomposes a single Histogram metric into individual metrics for count, bucket and sum
// non-Histograms give an empty slice as output
func decomposeHistogram(metric pmetric.Metric) pmetric.MetricSlice {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricSlice)
}

func getHistogramBucketsMetric(metric pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// need to add one more bucket at +Inf

func getHistogramSumMetric(metric pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

func getHistogramCountMetric(metric pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// decomposeSummaries decomposes any summaries present in the metric data into individual Gauges and Sums
// This is a noop if no Summaries are present, but otherwise makes a copy of the whole structure
// This exists because Sumo doesn't support OTLP summaries with proper quantile dimensions
func decomposeSummaries(md pmetric.Metrics) pmetric.Metrics {
	_ = "STUB: not implemented"
	// short circuit and do nothing if no Summaries are present
	return *new(pmetric.Metrics)
}

// decomposeSummary decomposes a single Summary metric into individual metrics for quantiles, count, and sum
// non-Summaries give an empty slice as output
func decomposeSummary(metric pmetric.Metric) pmetric.MetricSlice {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricSlice)
}

// getSummaryQuantilesMetric extracts quantile values from a summary as individual gauge datapoints
// Each quantile becomes a separate datapoint with a "quantile" attribute
func getSummaryQuantilesMetric(metric pmetric.Metric) pmetric.MetricSlice {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricSlice)
}

// Create one gauge metric for all quantiles (multiple datapoints with quantile attribute)

// Copy original attributes

// Add quantile as an attribute

// Set timestamps and value

func getSummaryCountMetric(metric pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// count is unitless

// Set sum properties for counter-like behavior

func getSummarySumMetric(metric pmetric.Metric) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// Set sum properties for counter-like behavior
