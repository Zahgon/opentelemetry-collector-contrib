// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudlogentryencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension"

import (
	"time"

	gojson "github.com/goccy/go-json"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	gcpProjectField        = "gcp.project"
	gcpOrganizationField   = "gcp.organization"
	gcpBillingAccountField = "gcp.billing_account"
	gcpFolderField         = "gcp.folder"
	gcpResourceTypeField   = "gcp.resource_type"

	gcpOperationIDField       = "gcp.operation.id"
	gcpOperationProducerField = "gcp.operation.producer"
	gcpOperationFirstField    = "gcp.operation.first"
	gcpOperationLast          = "gcp.operation.last"

	gcpCacheLookupField                   = "gcp.cache.lookup"
	gcpCacheHitField                      = "gcp.cache.hit"
	gcpCacheValidatedWithOriginSeverField = "gcp.cache.validated_with_origin_server"
	gcpCacheFillBytes                     = "gcp.cache.fill_bytes"

	refererHeaderField         = "http.request.header.referer"
	requestServerDurationField = "http.request.server.duration"

	gcpSplitUIDField   = "gcp.split.uid"
	gcpSplitIndexField = "gcp.split.index"
	gcpSplitTotalField = "gcp.split.total"

	gcpErrorGroupField = "gcp.error_group"

	gcpAppHubPrefix                       = "gcp.apphub"
	gcpAppHubDestinationPrefix            = "gcp.apphub_destination"
	gcpAppHubApplicationContainerField    = "application.container"
	gcpAppHubApplicationLocationField     = "application.location"
	gcpAppHubApplicationIDField           = "application.id"
	gcpAppHubServiceIDField               = "service.id"
	gcpAppHubServiceEnvironmentTypeField  = "service.environment_type"
	gcpAppHubServiceCriticalityTypeField  = "service.criticality_type"
	gcpAppHubWorkloadIDField              = "workload.id"
	gcpAppHubWorkloadEnvironmentTypeField = "workload.environment_type"
	gcpAppHubWorkloadCriticalityTypeField = "workload.criticality_type"
)

// getEncodingFormat maps GCP log types to encoding format values
func getEncodingFormat(logType string) string { _ = "STUB: not implemented"; return "" }

// See: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
type logEntry struct {
	ProtoPayload gojson.RawMessage `json:"protoPayload"`
	TextPayload  string            `json:"textPayload"`
	JSONPayload  gojson.RawMessage `json:"jsonPayload"`

	ReceiveTimestamp *time.Time `json:"receiveTimestamp"`
	Timestamp        *time.Time `json:"timestamp"`

	InsertID     string            `json:"insertId"`
	LogName      string            `json:"logName"`
	Severity     string            `json:"severity"`
	Trace        string            `json:"trace"`
	SpanID       string            `json:"spanId"`
	TraceSampled *bool             `json:"traceSampled"`
	Labels       map[string]string `json:"labels"`

	HTTPRequest *httpRequest `json:"httpRequest"`

	Resource *struct {
		Type   string            `json:"type"`
		Labels map[string]string `json:"labels"`
	} `json:"resource"`

	Operation *operation `json:"operation"`

	SourceLocation *sourceLocation `json:"sourceLocation"`

	Split *split `json:"split"`

	ErrorGroups []errorGroup `json:"errorGroups"`

	AppHub *appHub `json:"apphub"`

	AppHubDestination *appHub `json:"apphubDestination"`
}

type errorGroup struct {
	ID string `json:"id"`
}

type split struct {
	UID         string `json:"uid"`
	Index       *int64 `json:"index"`
	TotalSplits *int64 `json:"totalSplits"`
}

type sourceLocation struct {
	File     string `json:"file"`
	Line     string `json:"line"`
	Function string `json:"function"`
}
type operation struct {
	ID       string `json:"id"`
	Producer string `json:"producer"`
	First    *bool  `json:"first"`
	Last     *bool  `json:"last"`
}

type httpRequest struct {
	RequestMethod                  string `json:"requestMethod"`
	RequestURL                     string `json:"requestURL"`
	RequestSize                    string `json:"requestSize"`
	Status                         *int64 `json:"status"`
	ResponseSize                   string `json:"responseSize"`
	UserAgent                      string `json:"userAgent"`
	RemoteIP                       string `json:"remoteIP"`
	ServerIP                       string `json:"serverIP"`
	Referer                        string `json:"referer"`
	Latency                        string `json:"latency"`
	CacheLookup                    *bool  `json:"cacheLookup"`
	CacheHit                       *bool  `json:"cacheHit"`
	CacheValidatedWithOriginServer *bool  `json:"cacheValidatedWithOriginServer"`
	CacheFillBytes                 string `json:"cacheFillBytes"`
	Protocol                       string `json:"protocol"`
}

type appHub struct {
	Application *struct {
		Container string `json:"container"`
		Location  string `json:"location"`
		ID        string `json:"id"`
	} `json:"application"`
	Service *struct {
		ID              string `json:"id"`
		EnvironmentType string `json:"environmentType"`
		CriticalityType string `json:"criticalityType"`
	} `json:"service"`
	Workload *struct {
		ID              string `json:"id"`
		EnvironmentType string `json:"environmentType"`
		CriticalityType string `json:"criticalityType"`
	} `json:"workload"`
}

// handleHTTPRequestField will place the HTTP attributes in the log record
func handleHTTPRequestField(attributes pcommon.Map, req *httpRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// handleOperationField will place the operation attributes in the log record
func handleOperationField(attributes pcommon.Map, op *operation) { _ = "STUB: not implemented"; return }

// handleSourceLocationField will place the source location attributes in the log record
func handleSourceLocationField(attributes pcommon.Map, sourceLoc *sourceLocation) error {
	_ = "STUB: not implemented"
	return nil
}

// handleSplitField will place the split attributes in the log record
func handleSplitField(attributes pcommon.Map, s *split) { _ = "STUB: not implemented"; return }

// handleErrorGroupField will place all ids of the error group in a new log record attribute
func handleErrorGroupField(attributes pcommon.Map, errGroup []errorGroup) {
	_ = "STUB: not implemented"
	return
}

func handleAppHubField(attributes pcommon.Map, appHub *appHub, prefix string) {
	_ = "STUB: not implemented"
	return
}

// getTraceID will parse the given trace and return the decoding id
func getTraceID(trace string) ([16]byte, error) {
	_ = "STUB: not implemented"
	// Format: projects/my-gcp-project/traces/4ebc71f1def9274798cac4e8960d0095
	return nil, nil
}

// getTraceID will return the decoded span id
func getSpanID(spanIDStr string) ([8]byte, error) {
	_ = "STUB: not implemented"
	// TODO cloud Run sends invalid span id's, make sure we're not crashing,
	// see https://issuetracker.google.com/issues/338634230?pli=1
	return nil, nil
}

// getSeverityNumber will map the severity to the plog.SeverityNumber
func getSeverityNumber(severity string) plog.SeverityNumber {
	_ = "STUB: not implemented"
	// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
	return *new(plog.SeverityNumber)
}

func setBodyFromJSON(logRecord plog.LogRecord, value gojson.RawMessage) error {
	_ = "STUB: not implemented"
	// {json,proto,text}_payload -> Body
	return nil
}

// Note: json.Unmarshal will turn a bare string into a
// go string, so this call will correctly set the body
// to a string Value.

func setBodyFromText(logRecord plog.LogRecord, value string) { _ = "STUB: not implemented"; return }

func handleTextPayloadField(logRecord plog.LogRecord, value string) {
	_ = "STUB: not implemented"
	return
}

func handleJSONPayloadField(logRecord plog.LogRecord, value gojson.RawMessage, config Config) error {
	_ = "STUB: not implemented"
	return nil
}

func handleProtoPayloadField(logRecord plog.LogRecord, value gojson.RawMessage, config Config) error {
	_ = "STUB: not implemented"
	return nil
}

// handleLogNameField parses a GCP logName string and extracts resource identifiers and log type.
// The logName must follow one of these formats:
//   - "projects/[PROJECT_ID]/logs/[LOG_ID]"
//   - "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
//   - "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
//   - "folders/[FOLDER_ID]/logs/[LOG_ID]"
//
// The function populates the provided resourceAttr map with the extracted ID and log type
// under the appropriate attribute keys. Returns the log type and any error encountered.
func handleLogNameField(logName string, resourceAttr pcommon.Map) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func handlePayload(encodingFormat string, log logEntry, logRecord plog.LogRecord, scope pcommon.InstrumentationScope, cfg Config) error {
	_ = "STUB: not implemented"
	return nil
}

// Add encoding.format to scope attributes for audit logs

// Add encoding.format to scope attributes for VPC flow logs

// Add encoding.format to scope attributes for Load balancer logs

// TODO Add support for more log types

// Fall through to default payload handling for non-armor load balancer logs
// TODO Add support for more log types

// if the log type was not recognized, add the payload to the log record body

// handleLogEntryFields will place each entry of logEntry as either an attribute of the log,
// or as part of the log body, in case of payload.
func handleLogEntryFields(resourceAttributes pcommon.Map, scopeLogs plog.ScopeLogs, log logEntry, cfg Config) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle log name, get type and encoding format
