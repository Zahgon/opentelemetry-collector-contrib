// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package resourcetotelemetry // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// Settings defines configuration for converting resource attributes to telemetry attributes.
// When used, it must be embedded in the exporter configuration:
//
//	type Config struct {
//	  // ...
//	  resourcetotelemetry.Settings `mapstructure:"resource_to_telemetry_conversion"`
//	}
type Settings struct {
	// Enabled indicates whether to convert resource attributes to telemetry attributes. Default is `false`.
	Enabled bool `mapstructure:"enabled"`
	// ExcludeServiceAttributes indicates whether to exclude `service.name`, `service.instance.id` and `service.namespace`
	// resource attributes from being converted to metric attributes. Default is `false`.
	// When set to `true`, these attributes will not be added to metric labels since they are
	// already mapped to Prometheus `job` and `instance` labels respectively.
	ExcludeServiceAttributes bool `mapstructure:"exclude_service_attributes"`
}

type wrapperMetricsExporter struct {
	exporter.Metrics
	excludeServiceAttributes bool
}

func (wme *wrapperMetricsExporter) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (*wrapperMetricsExporter) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	// Always return true since this wrapper modifies data inplace.
	return *new(consumer.Capabilities)
}

// WrapMetricsExporter wraps a given exporter.Metrics and based on the given settings
// converts incoming resource attributes to metrics attributes.
func WrapMetricsExporter(set Settings, exporter exporter.Metrics) exporter.Metrics {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics)
}

func (wme *wrapperMetricsExporter) convertToMetricsAttributes(md pmetric.Metrics) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// Filter resource attributes if excludeServiceAttributes is enabled

// filterServiceAttributes returns a new Map without service.name and service.instance.id attributes.
func filterServiceAttributes(attrs pcommon.Map) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func shouldSkipResourceAttributeKey(k string) bool { _ = "STUB: not implemented"; return false }

// addAttributesToMetric adds additional labels to the given metric
func addAttributesToMetric(metric pmetric.Metric, labelMap pcommon.Map) {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return
}

func addAttributesToNumberDataPoints(ps pmetric.NumberDataPointSlice, newAttributeMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func addAttributesToHistogramDataPoints(ps pmetric.HistogramDataPointSlice, newAttributeMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func addAttributesToSummaryDataPoints(ps pmetric.SummaryDataPointSlice, newAttributeMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func addAttributesToExponentialHistogramDataPoints(ps pmetric.ExponentialHistogramDataPointSlice, newAttributeMap pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func joinAttributeMaps(from, to pcommon.Map) { _ = "STUB: not implemented"; return }
