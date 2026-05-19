// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfxexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// removeAccessToken removes the SFX access token label if found in the give resource metric as a resource attribute
func removeAccessToken(dest pmetric.ResourceMetrics) { _ = "STUB: not implemented"; return }

// matchedHistogramResourceMetrics returns a map with the keys set to the index of resource Metrics containing
// Histogram metric type.
// The value is another map consisting of the ScopeMetric indices in the RM which contain Histogram metric type as keys
// and the value as an int slice with indices of the Histogram metric within the given scope.
// Example output {1: {1: [0,2]}}.
// The above output can be interpreted as Resource at index 1 contains Histogram metrics.
// Within this resource specifically the scope metric at index 1 contain Histograms.
// Lastly, the scope metric at index 1 has two Histogram type metric which can be found at index 0 and 2.
func matchedHistogramResourceMetrics(md pmetric.Metrics) (matchedRMIdx map[int]map[int][]int) {
	_ = "STUB: not implemented"
	return nil
}

// matchedHistogramScopeMetrics returns a map with keys equal to the ScopeMetric indices in the input resource metric
// which contain Histogram metric type.
// And the value is an int slice with indices of the Histogram metric within the keyed scope metric.
// Example output {1: [0,2]}.
// The above output can be interpreted as scope metrics at index 1 contains Histogram metrics.
// And that the scope metric at index 1 has two Histogram type metric which can be found at index 0 and 2.
func matchedHistogramScopeMetrics(rm pmetric.ResourceMetrics) (matchedSMIdx map[int][]int) {
	_ = "STUB: not implemented"
	return nil
}

// matchedHistogramMetrics returns an int slice with indices of metrics which are of Histogram type
// within the input scope metric.
// Example output [0,2].
// The above output can be interpreted as input scope metric has Histogram type metric at index 0 and 2.
func matchedHistogramMetrics(ilm pmetric.ScopeMetrics) (matchedMetricsIdx []int) {
	_ = "STUB: not implemented"
	return nil
}

// getHistograms returns new Metrics slice containing only Histogram metrics found in the input
// and the count of histogram metrics
// This function also adds the host ID attribute to the resource if it can be derived from the resource attributes
func getHistograms(md pmetric.Metrics) (pmetric.Metrics, int) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), 0
}

// Iterate over those ResourceMetrics which were found to contain histograms

// Copy resource metric properties to dest

// Remove Sfx access token

// Iterate over ScopeMetrics which were found to contain histograms

// Copy scope properties to dest

// Iterate over Metrics which contain histograms

// Copy metric properties to dest
