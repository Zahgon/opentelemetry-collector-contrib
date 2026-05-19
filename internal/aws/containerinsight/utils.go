// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package containerinsight // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/containerinsight"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

// SumFields takes an array of type map[string]any and do
// the summation on the values corresponding to the same keys.
// It is assumed that the underlying type of any to be float64.
func SumFields(fields []map[string]any) map[string]float64 { _ = "STUB: not implemented"; return nil }

// Use the first element as the base

// IsNode checks if a type belongs to node level metrics (for EKS)
func IsNode(mType string) bool { _ = "STUB: not implemented"; return false }

// IsInstance checks if a type belongs to instance level metrics (for ECS)
func IsInstance(mType string) bool { _ = "STUB: not implemented"; return false }

// IsContainer checks if a type belongs to container level metrics
func IsContainer(mType string) bool { _ = "STUB: not implemented"; return false }

// IsPod checks if a type belongs to container level metrics
func IsPod(mType string) bool { _ = "STUB: not implemented"; return false }

func getPrefixByMetricType(mType string) string { _ = "STUB: not implemented"; return "" }

// MetricName returns the metric name based on metric type and measurement name
// For example, a type "node" and a measurement "cpu_utilization" gives "node_cpu_utilization"
func MetricName(mType, measurement string) string { _ = "STUB: not implemented"; return "" }

// RemovePrefix removes the prefix (e.g. "node_", "pod_") from the metric name
func RemovePrefix(mType, metricName string) string { _ = "STUB: not implemented"; return "" }

// GetUnitForMetric returns unit for a given metric
func GetUnitForMetric(metric string) string { _ = "STUB: not implemented"; return "" }

// ConvertToOTLPMetrics converts a field containing metric values and tags containing the relevant labels to OTLP metrics.
// For legacy reasons, the timestamp is stored in the tags map with the key "Timestamp", but, unlike other tags,
// it is not added as a resource attribute to avoid high-cardinality metrics.
func ConvertToOTLPMetrics(fields map[string]any, tags map[string]string, logger *zap.Logger) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// Do not add Timestamp as a resource attribute to avoid high-cardinality.

func intGauge(ilm pmetric.ScopeMetrics, metricName, unit string, value int64, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func doubleGauge(ilm pmetric.ScopeMetrics, metricName, unit string, value float64, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func initMetric(ilm pmetric.ScopeMetrics, name, unit string) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}
