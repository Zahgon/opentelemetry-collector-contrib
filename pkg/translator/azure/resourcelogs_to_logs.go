// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azure // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure"

import (
	"errors"

	"github.com/goccy/go-json"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

const (
	// Constants for OpenTelemetry Specs
	scopeName = "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure"

	// Constants for Azure Log Records
	azureCategory          = "azure.category"
	azureCorrelationID     = "azure.correlation.id"
	azureDuration          = "azure.duration"
	azureIdentity          = "azure.identity"
	azureOperationName     = "azure.operation.name"
	azureOperationVersion  = "azure.operation.version"
	azureProperties        = "azure.properties"
	azureResourceID        = "azure.resource.id"
	azureResultType        = "azure.result.type"
	azureResultSignature   = "azure.result.signature"
	azureResultDescription = "azure.result.description"
	azureTenantID          = "azure.tenant.id"
)

var errMissingTimestamp = errors.New("missing timestamp")

// azureRecords represents an array of Azure log records
// as exported via an Azure Event Hub
type azureRecords struct {
	Records []azureLogRecord `json:"records"`
}

// azureLogRecord represents a single Azure log following
// the common schema:
// https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/resource-logs-schema
type azureLogRecord struct {
	Time              string       `json:"time"`
	Timestamp         string       `json:"timeStamp"`
	ResourceID        string       `json:"resourceId"`
	TenantID          *string      `json:"tenantId"`
	OperationName     string       `json:"operationName"`
	OperationVersion  *string      `json:"operationVersion"`
	Category          string       `json:"category"`
	ResultType        *string      `json:"resultType"`
	ResultSignature   *string      `json:"resultSignature"`
	ResultDescription *string      `json:"resultDescription"`
	DurationMs        *json.Number `json:"durationMs"`
	CallerIPAddress   *string      `json:"callerIpAddress"`
	CorrelationID     *string      `json:"correlationId"`
	Identity          *any         `json:"identity"`
	Level             any          `json:"Level"`
	Location          *string      `json:"location"`
	Properties        *any         `json:"properties"`
}

var _ plog.Unmarshaler = (*ResourceLogsUnmarshaler)(nil)

type ResourceLogsUnmarshaler struct {
	Version     string
	Logger      *zap.Logger
	TimeFormats []string
}

func (r ResourceLogsUnmarshaler) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func getTimestamp(record azureLogRecord, formats ...string) (pcommon.Timestamp, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp), nil
}

// asTimestamp will parse an ISO8601 string into an OpenTelemetry
// nanosecond timestamp. If the string cannot be parsed, it will
// return zero and the error.
func asTimestamp(s string, formats ...string) (pcommon.Timestamp, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp), nil
}

// Try parsing with provided formats first

// Fallback to ISO 8601 parsing if no format matches

// asSeverity converts the Azure log level to equivalent
// OpenTelemetry severity numbers. If the log level is not
// valid, then the 'Unspecified' value is returned.
func asSeverity(number any) plog.SeverityNumber {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber)
}

func extractRawAttributes(log azureLogRecord) map[string]any { _ = "STUB: not implemented"; return nil }

func setIf(attrs map[string]any, key string, value *string) { _ = "STUB: not implemented"; return }
