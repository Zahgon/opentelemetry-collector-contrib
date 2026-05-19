// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricsgenerationprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

func getNameToMetricMap(rm pmetric.ResourceMetrics) map[string]pmetric.Metric {
	_ = "STUB: not implemented"
	return nil
}

// getMetricValue returns the value of the first data point from the given metric.
func getMetricValue(metric pmetric.Metric) float64 { _ = "STUB: not implemented"; return 0 }

// generateCalculatedMetrics creates a new metric based on the given rule and adds it to the scope metric.
// The value for newly calculated metrics is always a floating point number.
// Note: This method assumes the matchAttributes feature flag is enabled.
func generateCalculatedMetrics(rm pmetric.ResourceMetrics, metric2 pmetric.Metric, rule internalRule, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Calculates a new metric based on the calculation-type rule specified. New data points will be generated for each
// calculation of the input metrics where overlapping attributes have matching values.
func generateMetricFromMatchingAttributes(metric1, metric2 pmetric.Metric, rule internalRule, logger *zap.Logger) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// Setup to metric and get metric1 data points

// Get metric2 data points

// Always return true to ensure iteration over all attributes

func dataPointValue(dp pmetric.NumberDataPoint) float64 { _ = "STUB: not implemented"; return 0 }

func dataPointAttributesMatch(dp1, dp2 pmetric.NumberDataPoint) bool {
	_ = "STUB: not implemented"
	return false
}

// generateScalarMetrics creates a new metric based on a scalar type rule and adds it to the scope metric.
// The value for newly calculated metrics is always a floating point number.
func generateScalarMetrics(rm pmetric.ResourceMetrics, operand2 float64, rule internalRule, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

func generateMetricFromOperand(from pmetric.Metric, operand2 float64, operation string, logger *zap.Logger) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// Only add a new data point if it was a valid operation

// Append the new metric to the scope metrics. This will only append the new metric if it
// has data points.
func appendNewMetric(ilm pmetric.ScopeMetrics, newMetric pmetric.Metric, name, unit string) {
	_ = "STUB: not implemented"
	return
}

// Only create a new metric if valid data points were calculated successfully

func calculateValue(operand1, operand2 float64, operation, metricName string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
