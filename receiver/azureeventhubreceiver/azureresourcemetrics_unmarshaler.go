// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

const azureResourceID = "azure.resource.id"

type azureResourceMetricsUnmarshaler struct {
	buildInfo   component.BuildInfo
	logger      *zap.Logger
	TimeFormat  []string
	Aggregation string
}

type azureResourceMetricsConfiger interface {
	GetLogger() *zap.Logger
	GetBuildVersion() string
	GetTimeFormat() []string
	GetAggregation() string
}

// azureMetricRecords represents an array of Azure metric records
// as exported via an Azure Event Hub
type azureMetricRecords struct {
	Records []azureGenericMetricRecord `json:"records"`
}

type azureMetricAppender interface {
	AppendMetrics(azureResourceMetricsConfiger, *pmetric.Metrics) error
}

type azureGenericMetricRecord struct {
	Record azureMetricAppender
}

func (r *azureGenericMetricRecord) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// azureMetricRecord represents a single Azure Platform Metric record following
// the common schema does not exist (yet):
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/platform/stream-monitoring-data-event-hubs#data-formats
type azureResourceMetricRecord struct {
	Time       string `json:"time"`
	ResourceID string `json:"resourceId"`
	MetricName string `json:"metricName"`
	TimeGrain  string `json:"timeGrain"`

	Total   float64 `json:"total"`
	Count   float64 `json:"count"`
	Minimum float64 `json:"minimum"`
	Maximum float64 `json:"maximum"`
	Average float64 `json:"average"`
}

func (r *azureResourceMetricRecord) AppendMetrics(c azureResourceMetricsConfiger, md *pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// azureMetricRecord represents a single Azure Application Metric record
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appmetrics
type azureAppMetricRecord struct {
	Time string `json:"time"`

	ResourceID      string `json:"resourceId"`
	AppRoleInstance string `json:"AppRoleInstance"`
	AppRoleName     string `json:"AppRoleName"`
	AppVersion      string `json:"AppVersion"`
	SDKVersion      string `json:"SDKVersion"`

	ClientCountryOrRegion string `json:"ClientCountryOrRegion"`
	ClientOS              string `json:"ClientOS"`

	Properties map[string]string `json:"Properties"`

	MetricName string  `json:"Name"`
	Total      float64 `json:"Sum"`
	Minimum    float64 `json:"Min"`
	Maximum    float64 `json:"Max"`
	Count      float64 `json:"ItemCount"`
}

func (r *azureAppMetricRecord) AppendMetrics(c azureResourceMetricsConfiger, md *pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func newAzureResourceMetricsUnmarshaler(buildInfo component.BuildInfo, logger *zap.Logger, cfg *Config) eventMetricsUnmarshaler {
	_ = "STUB: not implemented"
	return *new(eventMetricsUnmarshaler)
}

// UnmarshalMetrics takes a byte array containing a JSON-encoded
// payload with Azure metric records and transforms it into
// an OpenTelemetry pmetric.Metrics object. The data in the Azure
// metric record appears as fields and attributes in the
// OpenTelemetry representation;
func (r *azureResourceMetricsUnmarshaler) UnmarshalMetrics(event *azureEvent) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (r *azureResourceMetricsUnmarshaler) GetLogger() *zap.Logger {
	_ = "STUB: not implemented"
	return nil
}

func (r *azureResourceMetricsUnmarshaler) GetBuildVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *azureResourceMetricsUnmarshaler) GetTimeFormat() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *azureResourceMetricsUnmarshaler) GetAggregation() string {
	_ = "STUB: not implemented"
	return ""

	// asTimestamp will parse an ISO8601 string into an OpenTelemetry
	// nanosecond timestamp. If the string cannot be parsed, it will
	// return zero and the error.
}

func asTimestamp(s string, formats []string) (pcommon.Timestamp, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp), nil
}

// Try parsing with provided formats first

// Fallback to ISO 8601 parsing if no format matches
