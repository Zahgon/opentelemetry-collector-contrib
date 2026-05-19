// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package goldendataset // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/goldendataset"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// Simple utilities for generating metrics for testing

// MetricsCfg holds parameters for generating dummy metrics for testing. Set values on this struct to generate
// metrics with the corresponding number/type of attributes and pass into MetricsFromCfg to generate metrics.
type MetricsCfg struct {
	// The type of metric to generate
	MetricDescriptorType pmetric.MetricType
	// MetricValueType is the type of the numeric value: int or double.
	MetricValueType pmetric.NumberDataPointValueType
	// If MetricDescriptorType is one of the Sum, this describes if the sum is monotonic or not.
	IsMonotonicSum bool
	// A prefix for every metric name
	MetricNamePrefix string
	// The number of instrumentation library metrics per resource
	NumILMPerResource int
	// The size of the MetricSlice and number of Metrics
	NumMetricsPerILM int
	// The number of labels on the LabelsMap associated with each point
	NumPtLabels int
	// The number of points to generate per Metric
	NumPtsPerMetric int
	// The number of Attributes to insert into each Resource's AttributesMap
	NumResourceAttrs int
	// The number of ResourceMetrics for the single MetricData generated
	NumResourceMetrics int
	// The base value for each point
	PtVal int
	// The start time for each point
	StartTime uint64
	// The duration of the steps between each generated point starting at StartTime
	StepSize uint64
}

// DefaultCfg produces a MetricsCfg with default values. These should be good enough to produce sane
// (but boring) metrics, and can be used as a starting point for making alterations.
func DefaultCfg() MetricsCfg { _ = "STUB: not implemented"; return *new(MetricsCfg) }

// MetricsFromCfg produces pmetric.Metrics with the passed-in config.
func MetricsFromCfg(cfg MetricsCfg) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

type metricGenerator struct {
	metricID int
}

func newMetricGenerator() metricGenerator { _ = "STUB: not implemented"; return *new(metricGenerator) }

func (g *metricGenerator) genMetricFromCfg(cfg MetricsCfg) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func (g *metricGenerator) populateIlm(cfg MetricsCfg, rm pmetric.ResourceMetrics) {
	_ = "STUB: not implemented"
	return
}

func (g *metricGenerator) populateMetrics(cfg MetricsCfg, ilm pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

//exhaustive:enforce

func (g *metricGenerator) populateMetricDesc(cfg MetricsCfg, metric pmetric.Metric) {
	_ = "STUB: not implemented"
	return
}

func populateNumberPoints(cfg MetricsCfg, pts pmetric.NumberDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func populateDoubleHistogram(cfg MetricsCfg, dh pmetric.Histogram) {
	_ = "STUB: not implemented"
	return
}

func setDoubleHistogramBounds(hdp pmetric.HistogramDataPoint, bounds ...float64) {
	_ = "STUB: not implemented"
	return
}

func addDoubleHistogramVal(hdp pmetric.HistogramDataPoint, val float64) {
	_ = "STUB: not implemented"
	return
}

// TODO: HasSum, Min, HasMin, Max, HasMax are not covered in tests.

func populatePtAttributes(cfg MetricsCfg, lm pcommon.Map) { _ = "STUB: not implemented"; return }

func getTimestamp(startTime, stepSize uint64, i int) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func populateExpoHistogram(cfg MetricsCfg, dh pmetric.ExponentialHistogram) {
	_ = "STUB: not implemented"
	return
}
