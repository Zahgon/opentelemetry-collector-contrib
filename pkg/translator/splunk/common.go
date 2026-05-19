// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunk // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/splunk"

import (
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

// nanoToEpochSeconds converts a nanosecond timestamp to epoch seconds.
func nanoToEpochSeconds(ts pcommon.Timestamp) float64 { _ = "STUB: not implemented"; return 0 }

const (
	// hecEventMetricType is the type of HEC event. Set to metric, as per https://docs.splunk.com/Documentation/Splunk/8.0.3/Metrics/GetMetricsInOther.
	hecEventMetricType = "metric"

	// https://docs.splunk.com/Documentation/Splunk/9.2.1/Metrics/Overview#What_is_a_metric_data_point.3F
	// metric name can contain letters, numbers, underscore, dot or colon. cannot start with number or underscore, or contain metric_name
	metricNamePattern = `^metric_name:([A-Za-z.:][A-Za-z0-9_.:\\-]*)$`
)

var metricNameRegexp = regexp.MustCompile(metricNamePattern)

// HecToOtelAttrs defines the mapping of Splunk HEC metadata to attributes
type HecToOtelAttrs struct {
	// Source indicates the mapping of the source field to a specific unified model attribute.
	Source string `mapstructure:"source"`
	// SourceType indicates the mapping of the sourcetype field to a specific unified model attribute.
	SourceType string `mapstructure:"sourcetype"`
	// Index indicates the mapping of the index field to a specific unified model attribute.
	Index string `mapstructure:"index"`
	// Host indicates the mapping of the host field to a specific unified model attribute.
	Host string `mapstructure:"host"`
}

func DefaultHecToOtelAttrs() HecToOtelAttrs { _ = "STUB: not implemented"; return *new(HecToOtelAttrs) }

// OtelToHecFields defines the mapping of attributes to HEC fields
type OtelToHecFields struct {
	// SeverityText informs the exporter to map the severity text field to a specific HEC field.
	SeverityText string `mapstructure:"severity_text"`
	// SeverityNumber informs the exporter to map the severity number field to a specific HEC field.
	SeverityNumber string `mapstructure:"severity_number"`
}

func DefaultOtelToHecFields() OtelToHecFields {
	_ = "STUB: not implemented"
	return *new(OtelToHecFields)
}

// Event represents a metric in Splunk HEC format
type Event struct {
	// type of event: set to "metric" or nil if the event represents a metric, or is the payload of the event.
	Event any `json:"event"`
	// dimensions and metric data
	Fields map[string]any `json:"fields,omitempty"`
	// hostname
	Host string `json:"host"`
	// optional description of the source of the event; typically the app's name
	Source string `json:"source,omitempty"`
	// optional name of a Splunk parsing configuration; this is usually inferred by Splunk
	SourceType string `json:"sourcetype,omitempty"`
	// optional name of the Splunk index to store the event in; not required if the token has a default index set in Splunk
	Index string `json:"index,omitempty"`
	// optional epoch time - set to zero if the event timestamp is missing or unknown (will be added at indexing time)
	Time float64 `json:"time,omitempty"`
}

// IsMetric returns true if the Splunk event is a metric.
func (e *Event) IsMetric() bool { _ = "STUB: not implemented"; return false }

// checks if the field name matches the requirements for a metric datapoint field,
// and returns the metric name and a bool indicating whether the field is a metric.
func getMetricNameFromField(fieldName string) (string, bool) {
	_ = "STUB: not implemented"
	// only consider metric name if it fits regex criteria.
	// use matches[1] since first element contains entire string.
	// first subgroup will be the actual metric name.
	return "", false
}

// GetMetricValues extracts metric key value pairs from a Splunk HEC metric.
func (e *Event) GetMetricValues() map[string]any { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON unmarshals the JSON representation of an event
func (e *Event) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
