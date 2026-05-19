// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetricutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector/internal/pmetricutil"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MoveResourcesIf calls f sequentially for each ResourceSpans present in the first pmetric.Metrics.
// If f returns true, the element is removed from the first pmetric.Metrics and added to the second pmetric.Metrics.
func MoveResourcesIf(from, to pmetric.Metrics, f func(pmetric.ResourceMetrics) bool) {
	_ = "STUB: not implemented"
	return
}

// CopyResourcesIf calls f sequentially for each ResourceSpans present in the first pmetric.Metrics.
// If f returns true, the element is copied from the first pmetric.Metrics to the second pmetric.Metrics.
func CopyResourcesIf(from, to pmetric.Metrics, f func(pmetric.ResourceMetrics) bool) {
	_ = "STUB: not implemented"
	return
}

// MoveMetricsWithContextIf calls f sequentially for each Metric present in the first pmetric.Metrics.
// If f returns true, the element is removed from the first pmetric.Metrics and added to the second pmetric.Metrics.
// Notably, the Resource and Scope associated with the Metric are created in the second pmetric.Metrics only once.
// Resources or Scopes are removed from the original if they become empty. All ordering is preserved.
func MoveMetricsWithContextIf(from, to pmetric.Metrics, f func(pmetric.ResourceMetrics, pmetric.ScopeMetrics, pmetric.Metric) bool) {
	_ = "STUB: not implemented"
	return
}

// CopyMetricsWithContextIf calls f sequentially for each Metric present in the first pmetric.Metrics.
// If f returns true, the element is copied from the first pmetric.Metrics to the second pmetric.Metrics.
// Notably, the Resource and Scope associated with the Metric are created in the second pmetric.Metrics only once.
func CopyMetricsWithContextIf(from, to pmetric.Metrics, f func(pmetric.ResourceMetrics, pmetric.ScopeMetrics, pmetric.Metric) bool) {
	_ = "STUB: not implemented"
	return
}

// MoveDataPointsWithContextIf calls f sequentially for each DataPoint present in the first pmetric.Metrics.
// If f returns true, the element is removed from the first pmetric.Metrics and added to the second pmetric.Metrics.
// Notably, the Resource, Scope, and Metric associated with the DataPoint are created in the second pmetric.Metrics only once.
// Resources, Scopes, or Metrics are removed from the original if they become empty. All ordering is preserved.
func MoveDataPointsWithContextIf(from, to pmetric.Metrics, f func(pmetric.ResourceMetrics, pmetric.ScopeMetrics, pmetric.Metric, any) bool) {
	_ = "STUB: not implemented"
	return
}

// TODO condense this code

// Do not remove unknown type.

// CopyDataPointsWithContextIf calls f sequentially for each DataPoint present in the first pmetric.Metrics.
// If f returns true, the element is copied from the first pmetric.Metrics to the second pmetric.Metrics.
// Notably, the Resource, Scope, and Metric associated with the DataPoint are created in the second pmetric.Metrics only once.
func CopyDataPointsWithContextIf(from, to pmetric.Metrics, f func(pmetric.ResourceMetrics, pmetric.ScopeMetrics, pmetric.Metric, any) bool) {
	_ = "STUB: not implemented"
	return
}

func copyResourceMetrics(from pmetric.ResourceMetrics, to pmetric.ResourceMetricsSlice) pmetric.ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ResourceMetrics)
}

func copyScopeMetrics(from pmetric.ScopeMetrics, to pmetric.ScopeMetricsSlice) pmetric.ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ScopeMetrics)
}

func copyMetricDescription(from pmetric.Metric, to pmetric.MetricSlice) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}
