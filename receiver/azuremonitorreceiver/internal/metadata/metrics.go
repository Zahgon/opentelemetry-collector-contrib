// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver/internal/metadata"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

const metricsPrefix = "azure_"

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config          MetricsBuilderConfig // config of the metrics builder.
	startTime       pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity int                  // maximum observed number of metrics per resource.
	metricsBuffer   pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo       component.BuildInfo  // contains version information
	metrics         map[string]*metricAzureAbstract
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	_ = "STUB: not implemented"

	// WithStartTime sets startTime on the metrics builder.
	return
}

func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	_ = "STUB: not implemented"
	return *new(MetricBuilderOption)
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	_ = "STUB: not implemented"
	return nil
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted metrics.
func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	_ = "STUB: not implemented"
	return
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	_ = "STUB: not implemented"

	// WithResource sets the provided resource on the emitted ResourceMetrics.
	// It's recommended to use ResourceBuilder to create the resource.
	return
}

func WithResource(res pcommon.Resource) ResourceMetricsOption {
	_ = "STUB: not implemented"
	return *new(ResourceMetricsOption)
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	_ = "STUB: not implemented"
	return
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) { _ = "STUB: not implemented"; return }

type metricAzureAbstract struct {
	data     pmetric.Metric // data buffer for generated metric.
	capacity int            // max observed number of data points added to the metric.
}

func (m *metricAzureAbstract) updateCapacity() { _ = "STUB: not implemented"; return }

func (m *metricAzureAbstract) init(name, unit string) {
	m.data.SetName(name)
	m.data.SetUnit(unit)
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (mb *MetricsBuilder) getMetric(resourceMetricID string) (*metricAzureAbstract, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (mb *MetricsBuilder) addMetric(resourceMetricID, logicalMetricID, unit string) (*metricAzureAbstract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mb *MetricsBuilder) AddDataPoint(
	resourceID,
	metric,
	aggregation,
	unit string,
	attributes map[string]*string,
	ts pcommon.Timestamp,
	val float64,
) {
	_ = "STUB: not implemented"
	return
}

func getLogicalMetricID(metric, aggregation string) string { _ = "STUB: not implemented"; return "" }

func getLogicalResourceMetricID(resourceID, logicalMetricID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (mb *MetricsBuilder) EmitAllMetrics(ils pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}
