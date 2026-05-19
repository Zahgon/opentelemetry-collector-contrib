// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetricassert // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetricassert"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// AssertMetrics compares actual against the assertion file at expectedPath.
//
// The comparison is order-insensitive across resources, scopes, metrics, and
// datapoints. Resource attributes, scope identity, metric metadata, and
// datapoint attribute permutations must match exactly. Values, timestamps,
// and exemplars are ignored.
func AssertMetrics(expectedPath string, actual pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func compareDocuments(expected, actual *document) error { _ = "STUB: not implemented"; return nil }

func compareResource(expected, actual resourceAssertion) error {
	_ = "STUB: not implemented"
	return nil
}

func indexScopes(ss []scopeAssertion) map[string]scopeAssertion {
	_ = "STUB: not implemented"
	return nil
}

func compareScope(expected, actual scopeAssertion) error { _ = "STUB: not implemented"; return nil }

func indexMetrics(ms []metricAssertion) map[string]metricAssertion {
	_ = "STUB: not implemented"
	return nil
}

func compareMetric(expected, actual metricAssertion) error { _ = "STUB: not implemented"; return nil }

func compareDatapoints(expected, actual []datapointAssertion) error {
	_ = "STUB: not implemented"
	return nil
}

// findMatchingAttributes returns the first unmatched index whose attributes
// satisfy the expected attribute map, or -1 if none do.
func findMatchingAttributes(expected map[string]any, matched []bool, n int, attrsAt func(int) map[string]any) int {
	_ = "STUB: not implemented"
	return 0
}

func compareAttributes(expected, actual map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func boolPtrEqual(a, b *bool) bool { _ = "STUB: not implemented"; return false }

func boolPtrString(p *bool) string { _ = "STUB: not implemented"; return "" }
