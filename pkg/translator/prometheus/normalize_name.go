// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheus // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// BuildCompliantName builds a Prometheus-compliant metric name for the specified metric
//
// Metric name is prefixed with specified namespace and underscore (if any).
// Namespace is not cleaned up. Make sure specified namespace follows Prometheus
// naming convention.
//
// See rules at https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels
// and https://prometheus.io/docs/practices/naming/#metric-and-label-naming
func BuildCompliantName(metric pmetric.Metric, namespace string, addMetricSuffixes bool) string {
	_ = "STUB: not implemented"
	return ""

	// Full normalization following standard Prometheus naming conventions
}

// Simple case (no full normalization, no units, etc.), we simply trim out forbidden chars

// Namespace?

// Metric name starts with a digit? Prefix it with an underscore

// Build a normalized name for the specified metric
func normalizeName(metric pmetric.Metric, namespace string) string {
	_ = "STUB: not implemented"
	// Split metric name in "tokens" (remove all non-alphanumeric)
	return ""
}

// Append unit if it exists

// Append _total for Counters

// Append _ratio for metrics with unit "1"
// Some Otel receivers improperly use unit "1" for counters of objects
// See https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aissue+some+metric+units+don%27t+follow+otel+semantic+conventions
// Until these issues have been fixed, we're appending `_ratio` for gauges ONLY
// Theoretically, counters could be ratios as well, but it's absurd (for mathematical reasons)

// Namespace

// Build the string from the tokens, separated with underscores

// Metric name cannot start with a digit, so prefix it with "_" in this case

// TrimPromSuffixes trims type and unit prometheus suffixes from a metric name.
// Following the [OpenTelemetry specs] for converting Prometheus Metric points to OTLP.
//
// [OpenTelemetry specs]: https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/data-model.md#metric-metadata
func TrimPromSuffixes(promName string, metricType pmetric.MetricType, unit string) string {
	_ = "STUB: not implemented"
	return ""
}

func removeTypeSuffixes(tokens []string, metricType pmetric.MetricType) []string {
	_ = "STUB: not implemented"
	return nil
}

// Only counters are expected to have a type suffix at this point.
// for other types, suffixes are removed during scrape.

func removeUnitSuffixes(nameTokens []string, unit string) []string {
	_ = "STUB: not implemented"
	return nil
}

func removeSuffix(tokens []string, suffix string) []string { _ = "STUB: not implemented"; return nil }

func RemovePromForbiddenRunes(s string) string { _ = "STUB: not implemented"; return "" }

// Remove the specified value from the slice
func removeItem(slice []string, value string) []string { _ = "STUB: not implemented"; return nil }
