// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package collectdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type collectDRecord struct {
	Time           *float64       `json:"time"`
	Host           *string        `json:"host"`
	Interval       *float64       `json:"interval"`
	Plugin         *string        `json:"plugin"`
	PluginInstance *string        `json:"plugin_instance"`
	TypeS          *string        `json:"type"`
	TypeInstance   *string        `json:"type_instance"`
	Message        *string        `json:"message"`
	Meta           map[string]any `json:"meta"`
	Severity       *string        `json:"severity"`
	Dstypes        []*string      `json:"dstypes"`
	Dsnames        []*string      `json:"dsnames"`
	Values         []*json.Number `json:"values"`
}

type createMetricInfo struct {
	DsType *string
	Val    *json.Number
	Name   string
}

func (cdr *collectDRecord) isEvent() bool { _ = "STUB: not implemented"; return false }

func (cdr *collectDRecord) protoTime() pcommon.Timestamp {
	_ = "STUB: not implemented"
	// Return 1970-01-01 00:00:00 +0000 UTC.
	return *new(pcommon.Timestamp)
}

func (cdr *collectDRecord) startTimestamp(metricType string) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (cdr *collectDRecord) appendToMetrics(logger *zap.Logger, scopeMetrics pmetric.ScopeMetrics, defaultLabels map[string]string) error {
	_ = "STUB: not implemented"
	// Ignore if record is an event instead of data point
	return nil
}

// Create new metric, get labels, then setting attribute and metric info
func (cdr *collectDRecord) newMetric(createMetric createMetricInfo, labels map[string]string) (pmetric.Metric, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric), nil
}

func setAttributes(labels map[string]string) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// Set new metric info with name, datapoint, time, attributes
func (cdr *collectDRecord) setMetric(createMetric createMetricInfo, atr pcommon.Map) (pmetric.Metric, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric), nil
}

// check type to decide metric type and return data point
func setDataPoint(typ string, metric pmetric.Metric) pmetric.NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.NumberDataPoint)
}

// getReasonableMetricName creates metrics names by joining them (if non empty) type.typeinstance
// if there are more than one dsname append .dsname for the particular uint. if there's only one it
// becomes a dimension.
func (cdr *collectDRecord) getReasonableMetricName(index int, attrs map[string]string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// pointTypeInstance extracts information from the TypeInstance field and appends to the metric name when possible.
func (cdr *collectDRecord) pointTypeInstance(attrs map[string]string, parts []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func isNilOrEmpty(str *string) bool { _ = "STUB: not implemented"; return false }

func addIfNotNullOrEmpty(m map[string]string, key string, val *string) {
	_ = "STUB: not implemented"
	return
}

func parseAndAddLabels(labels map[string]string, pluginInstance, host *string) {
	_ = "STUB: not implemented"
	return
}

func parseNameForLabels(labels map[string]string, key string, val *string) {
	_ = "STUB: not implemented"
	return
}
