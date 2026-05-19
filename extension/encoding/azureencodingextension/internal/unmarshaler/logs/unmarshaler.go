// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/constants"
)

// Commonly used non-SemConv attributes
const (
	// Predefined value for `cloudevents.event_source` resource attribute
	attributeCloudEventSourceValue = "azure.resource.log"
	// OpenTelemetry resource attribute name for Azure Tenant ID
	attributeAzureTenantID = "azure.tenant.id"
)

// logsResourceAttributes is a helper struct to hold resource attributes for specific log records
// Each Category Parser decides which Resource Attributes to populate
type logsResourceAttributes struct {
	ResourceID        string
	TenantID          string
	SubscriptionID    string
	Location          string
	SeviceNamespace   string
	ServiceName       string
	ServiceInstanceID string
	Environment       string
}

// scopeKey keys ScopeLogs by resource attributes and encoding.format so each scope has a single format.
type scopeKey struct {
	Resource logsResourceAttributes
	Format   constants.Format
}

// categoryHolder is a small helper struct to get `category`/`type` fields
// from each Log Record
type categoryHolder struct {
	Category string `json:"category"`
	Type     string `json:"type"` // Used in AppInsights logs instead of `category`
}

type ResourceLogsUnmarshaler struct {
	buildInfo         component.BuildInfo
	logger            *zap.Logger
	timeFormat        []string
	includeCategories map[string]bool
	excludeCategories map[string]bool
	hasIncludes       bool
}

func (r ResourceLogsUnmarshaler) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// ND JSON is a specific case...
// We will use bufio.Scanner trick to read it line by line
// as unmarshal each line as a Log Record

// Both formats are valid JSON and can be parsed directly
// `gojson.Path.Extract` is a bit faster and use ~25% less bytes per operation
// comparing to unmarshaling to intermediate structure (e.g. using `var recordsHolder []json.RawMessage`)

// This will allow us to parse Azure Log Records in both formats:
// 1) As exported to Azure Event Hub, e.g. `{"records": [ {...}, {...} ]}`
// 2) As exported to Azure Blob Storage, e.g. `[ {...}, {...} ]`

// This should never happen, but still...

// This should never happen, but still...

// This happens on empty input

// Set SemConv attributes

// Resource attributes parsed by Category Parser

// In Azure - Subscription is the closes analog to the Account,
// so we'll transform SubscriptionID into `cloud.account.id`

func (r ResourceLogsUnmarshaler) unmarshalRecord(allResourceScopeLogs map[scopeKey]plog.ScopeLogs, record []byte) {
	_ = "STUB: not implemented"
	// Despite of the fact that official Azure documentation states that exists common Logs schema
	// (see https://learn.microsoft.com/en-us/azure/azure-monitor/platform/resource-logs-schema),
	// in reality - it's not true, some Resources exposing Logs in totally different formats.
	// For example, Azure Service Bus logs doesn't conform schema above
	// (see examples here - https://github.com/noakup/AzMonLogsAgent/blob/main/NGSchema/AzureServiceBus/SampleInputRecords/ServiceBusOperationLogSample.json)
	// The only 2 fields that SHOULD be present in most Log schemas - `category` and `resourceId`
	// So, proper way to correctly decode incoming Log Record - first get value from `category` field
	// and Unmarshal record into category-specific struct.
	// That's actually double-unmarshaling, but there is no other way to parse variety of Azure Logs schemas
	return
}

// We couldn't do any SemConv conversion as it's an unknown Log Schema for us,
// because it doesn't have a "category" field which we rely on
// So we will save incoming Log Record as a JSON string into Body just
// not to loose data

// Filter out categories based on provided configuration

// Let's parse it

// Get timestamp for any of the possible fields

// Do not set Log Severity if it's not provided in the Log Record
// to avoid confusion with actual SeverityNumberUnspecified value

// Put Log Category anyway

// Parse Common Attributes + Properties (if applicable)

// getScopeLog gets or creates ScopeLogs keyed by resource attributes and encoding.format (one format per scope).
func (r ResourceLogsUnmarshaler) getScopeLog(allResourceScopeLogs map[scopeKey]plog.ScopeLogs, rs logsResourceAttributes, format constants.Format) plog.ScopeLogs {
	_ = "STUB: not implemented"
	return *new(plog.ScopeLogs)
}

// storeRawLog stores incoming Azure Resource Log Record as a string into log.Body
// It's used we couldn't do any SemConv conversion, for example:
// * In case when there is no "category" field
// * In case when JSON unmarshaling failed
// Stored record than can be used for debugging and fixing purposes
func (r ResourceLogsUnmarshaler) storeRawLog(allResourceScopeLogs map[scopeKey]plog.ScopeLogs, record []byte) {
	_ = "STUB: not implemented"
	// We couldn't do any SemConv conversion as it's an unknown Log Schema for us,
	// because it doesn't have a "category" field which we rely on
	// So we will save incoming Log Record as a JSON string into Body just
	// not to loose data
	return
}

// We couldn't get timestamp from incoming Record, so to keep the Log
// we will set timestamp to current time

// Set unspecified log level

// Put record to Body as-is

func NewAzureResourceLogsUnmarshaler(buildInfo component.BuildInfo, logger *zap.Logger, cfg LogsConfig) ResourceLogsUnmarshaler {
	_ = "STUB: not implemented"
	return *new(ResourceLogsUnmarshaler)
}
