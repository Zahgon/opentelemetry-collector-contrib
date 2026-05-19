// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetricassert // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetricassert"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// normalize produces the identity-only document form of m.
//
// Normalization merges compatible resources (by resource attributes), scopes
// (by name+version) and metrics (by name) so that batch boundaries do not
// influence the assertion. Datapoints are folded into a set keyed by their
// attribute values; duplicate logical MTS entries collapse to one.
func normalize(m pmetric.Metrics) *document { _ = "STUB: not implemented"; return nil }

func buildMetricAssertion(metric pmetric.Metric) metricAssertion {
	_ = "STUB: not implemented"
	return *new(metricAssertion)
}

func metricTypeString(t pmetric.MetricType) string { _ = "STUB: not implemented"; return "" }

func temporalityString(t pmetric.AggregationTemporality) string {
	_ = "STUB: not implemented"
	return ""
}

func extractDatapointAttributes(metric pmetric.Metric) []pcommon.Map {
	_ = "STUB: not implemented"
	return nil
}

func attrMapToRaw(m pcommon.Map) map[string]any { _ = "STUB: not implemented"; return nil }

// canonKey produces a stable string key for a map-like structure. It is used
// for map lookup and for deterministic sort order in the emitted document.
func canonKey(v any) string { _ = "STUB: not implemented"; return "" }

func sortedAny(v any) any { _ = "STUB: not implemented"; return *new(any) }
