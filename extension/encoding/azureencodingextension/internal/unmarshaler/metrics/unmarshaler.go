// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/metrics"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var timeGrains = map[string]int64{
	"PT1M":  60,
	"PT5M":  300,
	"PT15M": 900,
	"PT30M": 1800,
	"PT1H":  3600,
	"PT6H":  21600,
	"PT12H": 43200,
	"P1D":   86400,
}

// azureMetricRecord represents a single Azure Metric following
// the common schema does not exist (yet):
type azureMetricRecord struct {
	Time       string  `json:"time"`
	ResourceID string  `json:"resourceId"`
	MetricName string  `json:"metricName"`
	TimeGrain  string  `json:"timeGrain"`
	Total      float64 `json:"total"`
	Count      float64 `json:"count"`
	Minimum    float64 `json:"minimum"`
	Maximum    float64 `json:"maximum"`
	Average    float64 `json:"average"`
	// Fields below are available only via DCR, not via Diagnostic Settings
	Unit       string         `json:"unit"`
	Dimensions map[string]any `json:"dimension"`
}

type ResourceMetricsUnmarshaler struct {
	buildInfo    component.BuildInfo
	logger       *zap.Logger
	timeFormat   []string
	aggregations []MetricAggregation
}

// UnmarshalMetrics takes a byte array containing a JSON-encoded
// payload with Azure metric records and transforms it into
// an OpenTelemetry pmetric.Metrics object. The data in the Azure
// metric record appears as fields and attributes in the
// OpenTelemetry representation
func (r ResourceMetricsUnmarshaler) UnmarshalMetrics(buf []byte) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// ND JSON is a specific case...
// We will use bufio.Scanner trick to read it line by line
// as unmarshal each line as a Metric Record

// Both formats are valid JSON and can be parsed directly
// `gojson.Path.Extract` is a bit faster and use ~25% less bytes per operation
// comparing to unmarshaling to intermediate structure (e.g. using `var recordsHolder []json.RawMessage`)

// This will allow us to parse Azure Metric Records in both formats:
// 1) As exported to Azure Event Hub, e.g. `{"records": [ {...}, {...} ]}`
// 2) As exported to Azure Blob Storage, e.g. `[ {...}, {...} ]`

// This should never happen, but still...

// This should never happen, but still...

// This happens on empty input

// Do not add empty ResourceMetrics
// This can happen if there are no metrics for a given ResourceID
// or timeGrain/timestamp is invalid

func (r ResourceMetricsUnmarshaler) unmarshalRecord(allResourceScopeMetrics map[string]pmetric.ScopeMetrics, record []byte) {
	_ = "STUB: not implemented"
	return
}

// Grouping set of metrics by Azure ResourceID's

// Only for records exported via DCRs

func NewAzureResourceMetricsUnmarshaler(buildInfo component.BuildInfo, logger *zap.Logger, cfg MetricsConfig) ResourceMetricsUnmarshaler {
	_ = "STUB: not implemented"
	return *new(ResourceMetricsUnmarshaler)
}
