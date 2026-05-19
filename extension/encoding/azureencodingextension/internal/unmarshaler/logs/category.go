// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"encoding/json"
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// List of supported Azure Resource Log Categories
const (
	categoryApplicationGatewayAccessLog        = "ApplicationGatewayAccessLog"
	categoryApplicationGatewayPerformanceLog   = "ApplicationGatewayPerformanceLog"
	categoryApplicationGatewayFirewallLog      = "ApplicationGatewayFirewallLog"
	categoryAppServiceAppLogs                  = "AppServiceAppLogs"
	categoryAppServiceAuditLogs                = "AppServiceAuditLogs"
	categoryAppServiceAuthenticationLogs       = "AppServiceAuthenticationLogs"
	categoryAppServiceConsoleLogs              = "AppServiceConsoleLogs"
	categoryAppServiceFileAuditLogs            = "AppServiceFileAuditLogs"
	categoryAppServiceHTTPLogs                 = "AppServiceHTTPLogs"
	categoryAppServiceIPSecAuditLogs           = "AppServiceIPSecAuditLogs"
	categoryAppServicePlatformLogs             = "AppServicePlatformLogs"
	categoryAzureCdnAccessLog                  = "AzureCdnAccessLog"
	categoryAzureMSApplicationMetricsLog       = "ApplicationMetricsLogs"
	categoryAzureMSDiagnosticErrorLog          = "DiagnosticErrorLogs"
	categoryAzureMSOperationalLog              = "OperationalLogs"
	categoryAzureMSRuntimeAuditLog             = "RuntimeAuditLogs"
	categoryAzureMSVNetAndIPFilteringLog       = "VNetAndIPFilteringLogs"
	categoryDataFactoryActivityRuns            = "ActivityRuns"
	categoryDataFactoryPipelineRuns            = "PipelineRuns"
	categoryDataFactoryTriggerRuns             = "TriggerRuns"
	categoryFrontDoorAccessLog                 = "FrontDoorAccessLog"
	categoryFrontDoorHealthProbeLog            = "FrontDoorHealthProbeLog"
	categoryFrontdoorWebApplicationFirewallLog = "FrontDoorWebApplicationFirewallLog"
	categoryFunctionAppLogs                    = "FunctionAppLogs"
	categoryStorageRead                        = "StorageRead"
	categoryStorageWrite                       = "StorageWrite"
	categoryStorageDelete                      = "StorageDelete"
	categoryAdministrative                     = "Administrative"
	categoryAlert                              = "Alert"
	categoryAutoscale                          = "Autoscale"
	categorySecurity                           = "Security"
	categoryPolicy                             = "Policy"
	categoryRecommendation                     = "Recommendation"
	categoryResourceHealth                     = "ResourceHealth"
	categoryServiceHealth                      = "ServiceHealth"
)

// Non-SemConv attributes that are used for common Azure Log Record fields
const (
	// OpenTelemetry attribute name for Azure Correlation ID,
	// from `correlationId` field in Azure Log Record
	attributeAzureCorrelationID = "azure.correlation_id"

	// OpenTelemetry attribute name for Azure Operation Version,
	// from `operationVersion` field in Azure Log Record
	attributeAzureOperationVersion = "azure.operation.version"

	// OpenTelemetry attribute name for generic Azure Operation Duration,
	// from `durationMs` field in Azure Log Record
	attributeAzureOperationDuration = "azure.operation.duration"

	// OpenTelemetry attribute name for Azure Log Record properties,
	// from `properties` field in Azure Log Record
	// Used when we cannot map parse "properties" field or
	// cannot map parsed "properties" to attributes directly
	attributesAzureProperties = "azure.properties"

	// OpenTelemetry attribute name for Azure Result Type,
	// from `resultType` field in Azure Log Record
	attributeAzureResultType = "azure.result.type"

	// OpenTelemetry attribute name for Azure Result Signature,
	// from `resultSignature` field in Azure Log Record
	attributeAzureResultSignature = "azure.result.signature"

	// OpenTelemetry attribute name for Azure Result Description,
	// from `resultDescription` field in Azure Log Record
	attributesAzureResultDescription = "azure.result.description"
)

// Common Non-SemConv attributes that are used in "properties" fields across multiple
// Azure Log Record Categories
const (
	// OpenTelemetry attribute name for "Host" HTTP Header value
	attributeHTTPHeaderHost = "http.request.header.host"

	// OpenTelemetry attribute name for the WAF action taken on the request
	attributeSecurityRuleActionKey = "security_rule.action"

	// OpenTelemetry attribute name for the operations mode of the WAF policy
	attributeSecurityRuleRulesetModeKey = "security_rule.ruleset.mode"

	// OpenTelemetry attribute name for Error Code
	attributeErrorCode = "error.code"

	// OpenTelemetry attribute name for Azure HTTP Request Duration,
	// from `durationMs` field in Azure Log Record
	attributeAzureRequestDuration = "azure.request.duration"
)

var errNoTimestamp = errors.New("no valid time fields are set on Log record ('time' or 'timestamp')")

// azureLogRecord is a common interface for all category-specific structures
type azureLogRecord interface {
	GetResource() logsResourceAttributes
	GetTimestamp(formats ...string) (pcommon.Timestamp, error)
	GetLevel() (plog.SeverityNumber, string, bool)
	PutCommonAttributes(attrs pcommon.Map, body pcommon.Value)
	PutProperties(attrs pcommon.Map, body pcommon.Value) error
}

// azureLogRecordBase represents a single Azure log following the common schema:
// https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/resource-logs-schema
// This schema are applicable to most Resource Logs and
// can be extended with additional fields for specific Log Categories
type azureLogRecordBase struct {
	Time              string       `json:"time"`      // most Categories use this field for timestamp
	TimeStamp         string       `json:"timestamp"` // some Categories use this field for timestamp
	ResourceID        string       `json:"resourceId"`
	TenantID          string       `json:"tenantId"`
	OperationName     string       `json:"operationName"`
	OperationVersion  *string      `json:"operationVersion"`
	ResultType        *string      `json:"resultType"`
	ResultSignature   *string      `json:"resultSignature"`
	ResultDescription *string      `json:"resultDescription"`
	DurationMs        *json.Number `json:"durationMs"` // int
	CallerIPAddress   *string      `json:"callerIpAddress"`
	CorrelationID     *string      `json:"correlationId"`
	Level             *json.Number `json:"level"`
	Location          string       `json:"location"`
}

// GetResource returns resource attributes for the parsed Log Record
// As for now it includes ResourceID, TenantID and Location
func (r *azureLogRecordBase) GetResource() logsResourceAttributes {
	_ = "STUB: not implemented"
	return *new(logsResourceAttributes)
}

// GetTimestamp tries to parse timestamp from either `time` or `timestamp` fields
// using provided list of time formats.
// If both fields are empty (undefined), or parsing failed - return an error
func (r *azureLogRecordBase) GetTimestamp(formats ...string) (pcommon.Timestamp, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp), nil
}

// GetLevel tries to convert the Log Level into OpenTelemetry SeverityNumber
// If level is not set - return SeverityNumberUnspecified and flag that level is not set
// If level is set, but invalid - return SeverityNumberUnspecified and flag that level is set
func (r *azureLogRecordBase) GetLevel() (plog.SeverityNumber, string, bool) {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber), "", false
}

// Saving original log.Level text,
// not the internal OpenTelemetry SeverityNumber -> SeverityText mapping

// PutCommonAttributes puts already parsed common attributes into provided Attributes Map/Body
func (r *azureLogRecordBase) PutCommonAttributes(attrs pcommon.Map, _ pcommon.Value) {
	_ = "STUB: not implemented"
	// Common fields for all Azure Resource Log Categories should be
	// placed as attributes, no matter if we can map the category or not
	return
}

// Identity is NOT processed here. Each category-specific struct is
// responsible for calling the appropriate identity parser in its own
// PutCommonAttributes override, because the identity field has different
// structures across Azure log categories (Activity, Storage, etc.).

// PutProperties puts already attributes from "properties" field into provided Attributes Map/Body
// MUST be implemented by each specific logCategory structure if "properties" field is expected there
func (*azureLogRecordBase) PutProperties(_ pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	// By default - no "properties", so nothing to do here
	return nil
}

// azureLogRecordGeneric represents a single Azure log following the common schema,
// but has unknown for us Category.
// In this case we couldn't correctly map properties to attributes and simply copy them
// as-is to the attributes.
// Identity is not handled for unknown categories - each known category handles
// its own identity structure with a typed struct.
type azureLogRecordGeneric struct {
	azureLogRecordBase

	Properties json.RawMessage `json:"properties"`
}

func (r *azureLogRecordGeneric) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// Nothing to parse

// We expect "properties" to be a correct JSON object in most cases,
// so we'll try to parse it as JSON here
// If parsing will fail - we will put value of "properties" field
// into `azure.properties` Attribute and return parse error to caller

// Put everything into attributes

// Keep all other fields as-is

// processLogRecord tries to parse incoming record based of provided logCategory
func processLogRecord(logCategory string, record []byte) (azureLogRecord, error) {
	_ = "STUB: not implemented"
	return *new(azureLogRecord), nil
}

// StorageRead, StorageWrite, StorageDelete share the same properties,
// called StorageBlobLogs, see https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/storagebloblogs

// Unfortunately, "goccy/go-json" has a bug with case-insensitive key matching
// for nested structures, so we have to use jsoniter here
// see https://github.com/goccy/go-json/issues/470
