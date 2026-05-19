// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

type MetricType string

const (
	MetricTypeGauge                MetricType = "Gauge"
	MetricTypeSum                  MetricType = "Sum"
	MetricTypeHistogram            MetricType = "Histogram"
	MetricTypeExponentialHistogram MetricType = "ExponentialHistogram"
)

// String is used both by fmt.Print and by Cobra in help text
func (e *MetricType) String() string {
	_ = "STUB: not implemented"

	// Set must have pointer receiver so it doesn't change the value of a copy
	return ""
}

func (e *MetricType) Set(v string) error { _ = "STUB: not implemented"; return nil }

// Type is only used in help text
func (*MetricType) Type() string { _ = "STUB: not implemented"; return "" }
