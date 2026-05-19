// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operationsmanagement // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter/internal/operationsmanagement"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

// BMCHelixOMMetric represents the structure of the payload that will be sent to BMC Helix Operations Management
type BMCHelixOMMetric struct {
	Labels  map[string]string  `json:"labels"`
	Samples []BMCHelixOMSample `json:"samples"`
}

// BMCHelixOMSample represents the individual sample for a metric
type BMCHelixOMSample struct {
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

// MetricsProducer is responsible for converting OpenTelemetry metrics into BMC Helix Operations Management metrics
type MetricsProducer struct {
	logger                     *zap.Logger
	previousCounters           map[string]BMCHelixOMSample
	enrichMetricWithAttributes bool
}

// NewMetricsProducer creates a new MetricsProducer
func NewMetricsProducer(logger *zap.Logger, enrichMetricWithAttributes bool) *MetricsProducer {
	_ = "STUB: not implemented"
	return nil
}

// coreAttributes are label keys that should be ignored when building metric name suffixes.
var coreAttributes = map[string]struct{}{
	"source":                 {},
	"unit":                   {},
	"hostType":               {},
	"isDeviceMappingEnabled": {},
	"metricName":             {},
	"hostname":               {},
	"entityTypeId":           {},
	"entityName":             {},
	"instanceName":           {},
	"entityId":               {},
}

const rateMetricFlag = "bmchelix.requiresRateMetric"

// ProduceHelixPayload takes the OpenTelemetry metrics and converts them into the BMC Helix Operations Management metric format
func (mp *MetricsProducer) ProduceHelixPayload(metrics pmetric.Metrics) ([]BMCHelixOMMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Iterate through each pmetric.ResourceMetrics instance

// Extract resource-level attributes (e.g., "host.name", "service.instance.id")

// Iterate through each pmetric.ScopeMetrics within the pmetric.ResourceMetrics instance

// Iterate through each individual pmetric.Metric instance

// Create the payload for each metric

// Grow the helixMetrics slice for the new metrics

// Loop through the newly created metrics and append them to the helixMetrics slice
// while also creating parent entities for container metrics

// appends the metric to the helixMetrics slice and creates a parent entity if it doesn't exist
func appendMetricWithParentEntity(helixMetrics []BMCHelixOMMetric, helixMetric BMCHelixOMMetric, containerParentEntities map[string]BMCHelixOMMetric) []BMCHelixOMMetric {
	_ = "STUB: not implemented"
	// Extract parent entity information
	return nil
}

// Create a parent entity if not already created

// Represents the parent entity itself

// Parent entities don't have samples

// Add parent reference to the child metric

// createHelixMetrics converts each OpenTelemetry datapoint into an individual BMCHelixOMMetric
func (mp *MetricsProducer) createHelixMetrics(metric pmetric.Metric, resourceAttrs map[string]string) ([]BMCHelixOMMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the metric is a counter, add a flag to compute the rate metric later

// This will create a new enriched metric with distinguishing attributes appended to the metric name
// for consistent identification in BMC Helix Operations Management

// Propagate rate flag if needed (Sum monotonic)

// Remove entityId from the original metric since a derived metric is being created.
// Without entityId, the original metric is forwarded to victoriametrics only (not BHOM),
// while the enriched metric with its unique name is reported in BHOM.

// Remove rate flag from the original metric to prevent rate computation with an
// empty entityId (":"+metricName key collision)

// This will create a new enriched metric with distinguishing attributes appended to the metric name
// for consistent identification in BMC Helix Operations Management

// Remove entityId from the original metric since a derived metric is being created.
// Without entityId, the original metric is forwarded to victoriametrics only (not BHOM),
// while the enriched metric with its unique name is reported in BHOM.

// Add percentage variants for ratio metrics (unit "1")

// Compute rate metrics for counter metrics that require it
// This will add a new metric with the same labels but with ".rate" suffix in the metric name
// and the value being the rate of change per second

// addRateVariants checks each metric for the 'bmchelix.requiresRateMetric' label
// and computes the rate metric from the counter metric if required.
func (mp *MetricsProducer) addRateVariants(helixMetrics []BMCHelixOMMetric) []BMCHelixOMMetric {
	_ = "STUB: not implemented"
	return nil
}

// Compute the rate metric from the counter metric

// Add the rate metric to the helixMetrics slice

// Remove the 'bmchelix.requiresRateMetric' label

// createSingleDatapointMetric creates a single BMCHelixOMMetric from a single OpenTelemetry datapoint
func (mp *MetricsProducer) createSingleDatapointMetric(dp pmetric.NumberDataPoint, metric pmetric.Metric, resourceAttrs map[string]string) (*BMCHelixOMMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add resource attributes with normalization

// Set the metric unit

// Set the host type

// Indicates the monitor in the hierarchy that is mapped to the device

// Update the metric name for the BMC Helix Operations Management payload

// Update the entity information

// Update the entity information for the BMC Helix Operations Management payload
func (*MetricsProducer) updateEntityInformation(labels map[string]string, metricName string, resourceAttrs map[string]string, dpAttributes map[string]any) error {
	_ = "STUB: not implemented"
	// Try to get the hostname from resource attributes first
	return nil
}

// Fallback to metric attributes if not found or empty in resource attributes

// Add the hostname as a label (required for BMC Helix Operations Management payload)

// Convert metricAttrs from map[string]any to map[string]string for compatibility

// Add the resource attributes to the metric attributes

// entityTypeId is required for the BMC Helix Operations Management payload

// entityName is required for the BMC Helix Operations Management payload

// Set the entityTypeId, entityId, instanceName and entityName in labels
// Use NormalizeEntityValue to ensure ":" is not present as it is the separator in entityId

// newSample creates a new BMCHelixOMSample from the OpenTelemetry data point
func newSample(dp pmetric.NumberDataPoint) BMCHelixOMSample {
	_ = "STUB: not implemented"
	return *new(BMCHelixOMSample)
}

// convert int to float for consistency

// extractResourceAttributes extracts the resource attributes from OpenTelemetry resource data
func extractResourceAttributes(resource pcommon.Resource) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// createEnrichedMetricWithDpAttributes creates a copy of the metric with non-core datapoint attribute
// values appended to the metric name as a dot-separated suffix. Attribute keys are sorted
// alphabetically and the resulting metric name is normalized (e.g., special chars replaced with "_").
// Returns nil if no non-core attributes exist. The original metric is not modified.
func createEnrichedMetricWithDpAttributes(metric *BMCHelixOMMetric, dpAttrs map[string]any) *BMCHelixOMMetric {
	_ = "STUB: not implemented"
	// Collect dp attribute keys excluding core ones (sorted using insertSorted)
	return nil
}

// Values-only suffix in sorted key order

// If the enriched metric name is empty, skip this enriched metric

// Binary-inserted sorted slice
func insertSorted(keys []string, key string) []string { _ = "STUB: not implemented"; return nil }

// find insertion index
// grow the slice by 1
// shift right to make room
// insert the new key

// addPercentageVariants adds percentage variants of metrics that are ratios (unit "1")
// This is done to ensure that the BMC Helix Operations Management payload contains both the original
// ratio metric and its percentage variant, which is often useful for visualization and analysis.
func addPercentageVariants(metrics []BMCHelixOMMetric) []BMCHelixOMMetric {
	_ = "STUB: not implemented"
	return nil
}

// Not a ratio

// Clone the original

// Rename metricName

// Convert sample value

// toPercentMetricName converts a metric name to its percentage variant
func toPercentMetricName(originalName string) string { _ = "STUB: not implemented"; return "" }

// already transformed

// computeRateMetricFromCounter computes a rate metric from a counter metric
func (mp *MetricsProducer) computeRateMetricFromCounter(metric BMCHelixOMMetric) *BMCHelixOMMetric {
	_ = "STUB: not implemented"
	return nil
}

// not enough data

// Avoid negative rates

// ms to sec

// Clone labels

// Modify metric name and unit for rate
