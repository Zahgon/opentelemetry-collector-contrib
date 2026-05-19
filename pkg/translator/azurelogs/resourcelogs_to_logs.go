// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurelogs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azurelogs"

import (
	"encoding/json"
	"errors"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

const (
	// Constants for OpenTelemetry Specs
	scopeName = "otelcol/azureresourcelogs"

	attributeAzureCategory          = "azure.category"
	attributeAzureCorrelationID     = "azure.correlation_id"
	attributeAzureDuration          = "azure.duration"
	attributeAzureLocation          = "azure.location"
	attributeAzureOperationName     = "azure.operation.name"
	attributeAzureOperationVersion  = "azure.operation.version"
	attributeEventOriginal          = "event.original"
	attributeAzureResultType        = "azure.result.type"
	attributeAzureResultSignature   = "azure.result.signature"
	attributeAzureResultDescription = "azure.result.description"

	// Constants for Azure Log Record body fields
	azureCategory          = "category"
	azureCorrelationID     = "correlation.id"
	azureDuration          = "duration"
	azureIdentity          = "identity"
	azureOperationName     = "operation.name"
	azureOperationVersion  = "operation.version"
	azureProperties        = "properties"
	azureResultType        = "result.type"
	azureResultSignature   = "result.signature"
	azureResultDescription = "result.description"
	azureTenantID          = "tenant.id"

	// Constants for Identity > claims
	identityClaimIssuer    = "iss"
	identityClaimSubject   = "sub"
	identityClaimAudience  = "aud"
	identityClaimExpires   = "exp"
	identityClaimNotBefore = "nbf"
	identityClaimIssuedAt  = "iat"

	identityClaimScope                 = "http://schemas.microsoft.com/identity/claims/scope"
	identityClaimType                  = "idtyp"
	identityClaimApplicationID         = "appid"
	identityClaimAuthMethodsReferences = "http://schemas.microsoft.com/claims/authnmethodsreferences"
	identityClaimProvider              = "http://schemas.microsoft.com/identity/claims/identityprovider"
	identityClaimIdentifierObject      = "http://schemas.microsoft.com/identity/claims/objectidentifier"
	identityClaimIdentifierName        = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier"
	identityClaimEmailAddress          = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"
)

var errMissingTimestamp = errors.New("missing timestamp")

// as exported via an Azure Event Hub
type azureRecords struct {
	Records []azureLogRecord `json:"records"`
}

type evidence struct {
	Role                string `json:"role"`
	RoleAssignmentScope string `json:"roleAssignmentScope"`
	RoleAssignmentID    string `json:"roleAssignmentId"`
	RoleDefinitionID    string `json:"roleDefinitionId"`
	PrincipalID         string `json:"principalId"`
	PrincipalType       string `json:"principalType"`
}

type authorization struct {
	Scope    string    `json:"scope"`
	Action   string    `json:"action"`
	Evidence *evidence `json:"evidence"`
}

// identity describes the identity of the user or application that performed the operation
// described by the log event.
type identity struct {
	// Claims usually contains the JWT token used by Active Directory
	// to authenticate the user or application to perform this
	// operation in Resource Manager.
	Claims        map[string]string `json:"claims"`
	Authorization *authorization    `json:"authorization"`
}

// azureLogRecord represents a single Azure log following
// the common schema:
// https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/resource-logs-schema
type azureLogRecord struct {
	Time              string          `json:"time"`
	Timestamp         string          `json:"timeStamp"`
	ResourceID        string          `json:"resourceId"`
	TenantID          *string         `json:"tenantId"`
	OperationName     string          `json:"operationName"`
	OperationVersion  *string         `json:"operationVersion"`
	Category          string          `json:"category"`
	ResultType        *string         `json:"resultType"`
	ResultSignature   *string         `json:"resultSignature"`
	ResultDescription *string         `json:"resultDescription"`
	DurationMs        *json.Number    `json:"durationMs"`
	CallerIPAddress   *string         `json:"callerIpAddress"`
	CorrelationID     *string         `json:"correlationId"`
	Identity          json.RawMessage `json:"identity"`
	Level             *json.Number    `json:"Level"`
	Location          *string         `json:"location"`
	Properties        json.RawMessage `json:"properties"`
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

// Create rawRecordMap on first use

// TODO @constanca-m This will be removed once the categories
// are properly mapped to the semantic conventions in
// category_logs.go

// NOTE: event.name semantically belongs on each LogRecord (not the Resource).
// Legacy behavior incorrectly placed it on the Resource for all records.
// Use pkg.translator.azurelogs.EmitV1LogConventions to migrate to SetEventName().

func getTimestamp(record *azureLogRecord, formats ...string) (pcommon.Timestamp, error) {
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

// asTimeFromUnixTimestamp converts a Unix timestamp string to a time.Time.
func parseUnixTimestamp(unixStr string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// asSeverity converts the Azure log level to equivalent
// OpenTelemetry severity numbers. If the log level is not
// valid, then the 'Unspecified' value is returned.
func asSeverity(number json.Number) plog.SeverityNumber {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber)
}

func putStrPtr(field string, value *string, record plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

// addCommonSchema adds the common schema attributes to the log record.
func addCommonSchema(log *azureLogRecord, record plog.LogRecord) { _ = "STUB: not implemented"; return }

func extractRawAttributes(log *azureLogRecord, rawRecord json.RawMessage) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// The original log needs to be preserved for logs that don't have a properties field

// Format the JSON with proper indentation to match expected output

func copyPropertiesAndApplySemanticConventions(category string, properties []byte, attrs map[string]any) {
	_ = "STUB: not implemented"
	return
}

// TODO @constanca-m: This is a temporary workaround to
// this function. This will be removed once category_logs.log
// is implemented for all currently supported categories

// Try primitive value

// Parsing failed completely - just return without setting properties

func setIf(attrs map[string]any, key string, value *string) { _ = "STUB: not implemented"; return }

// addIdentityAttributes extracts identity details
//
// The `identity` field is part of the Top-level common schema for
// resource logs and it's also in use in the activity logs.
//
// We're applying the strategy to only pick the identity elements
// that we know are useful. This approach also minimizes the risk
// of accidentally including sensitive data.
func addIdentityAttributes(identityJSON json.RawMessage, record plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

// Authorization
// ------------------------------------------------------------

// Claims
// ------------------------------------------------------------

// Extract known claims details we want to include in the
// log record.
// Extract common claim fields
