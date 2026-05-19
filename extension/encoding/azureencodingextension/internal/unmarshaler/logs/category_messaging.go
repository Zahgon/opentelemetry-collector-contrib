// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// Non-SemConv attributes that are used for common Azure Messaging Log Record fields
const (
	// OpenTelemetry attribute name for Azure Scale Unit name
	attributeAzureMSScaleUnit = "azure.autoscale.unit"
)

// Non-SemConv attributes specific for each Azure Messaging Log Record types
const (
	// OpenTelemetry attribute name for Task Name
	attributeAzureMSTaskName = "azure.messaging.task.name"

	// OpenTelemetry attribute name for Azure Messaging Error Count
	attributeAzureMSErrorCount = "azure.messaging.error.count"

	// OpenTelemetry attribute name for Partition ID
	attributeMessagingPartitionID = "messaging.destination.partition.id"

	// OpenTelemetry attribute name for Azure Messaging Auth Type (Microsoft Entra ID or SAS Policy)
	attributeAzureAuthType = "azure.auth.type"

	// OpenTelemetry attribute name for Azure Messaging Auth ID (Microsoft Entra application ID or SAS policy name)
	attributeAzureAuthID = "azure.auth.id"

	// OpenTelemetry attribute name for Messaging Message Count
	// Total number of operations performed during the aggregated period of 1 minute
	attributeMessagingMessageCount = "messaging.message.count"

	// OpenTelemetry attribute name for the caller of operation (the Azure portal or management client)
	attributeClientType = "client.type"

	// OpenTelemetry attribute name for the reason why the action was done
	attributeSecurityEvaluationReason = "security_rule.evaluation.reason"

	// OpenTelemetry attribute name for the number of times taken by security rule
	attributeSecurityEvaluationCount = "security_rule.evaluation.count"
)

// azureMSCommon it's common struct for all Azure Messaging Audit logs,
// like ServiceBus, EventHub, etc. (AZMS*******Logs)
type azureMSCommon struct {
	EventTimestamp  string `json:"eventTimestamp"`
	EventTimeString string `json:"EventTimeString"`
	Environment     string `json:"Environment"`
	Region          string `json:"Region"`
	ScaleUnit       string `json:"ScaleUnit"`
	ActivityID      string `json:"ActivityId"`
	ActivityName    string `json:"ActivityName"`
	SubscriptionID  string `json:"SubscriptionId"`
	ResourceID      string `json:"ResourceId"`
	NamespaceName   string `json:"NamespaceName"`
	EntityType      string `json:"EntityType"`
	EntityName      string `json:"EntityName"`
}

func (r *azureMSCommon) GetResource() logsResourceAttributes {
	_ = "STUB: not implemented"
	return *new(logsResourceAttributes)
}

func (r *azureMSCommon) GetTimestamp(formats ...string) (pcommon.Timestamp, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp), nil
}

func (*azureMSCommon) GetLevel() (plog.SeverityNumber, string, bool) {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber), "", false
}

func (r *azureMSCommon) PutCommonAttributes(attrs pcommon.Map, _ pcommon.Value) {
	_ = "STUB: not implemented"
	return
}

// EntityType is actually the messaging system name,
// so we'll try to map it to SemConv "messaging.system" attribute

// If EntityType is not set or empty - we'll use ResourceID to detect messaging system

func (*azureMSCommon) PutProperties(_ pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	// By default - no "properties", so nothing to do here
	return nil
}

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/service-bus-messaging/monitor-service-bus-reference.md#diagnostic-error-logs
// and https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/azmsdiagnosticerrorlogs
// Available for microsoft.servicebus/namespaces and microsoft.eventhub/namespaces
type azureMSDiagnosticErrorLog struct {
	azureMSCommon

	TaskName string `json:"TaskName"`

	OperationResult string      `json:"OperationResult"`
	ErrorMessage    string      `json:"ErrorMessage"`
	ErrorCount      json.Number `json:"ErrorCount"` // int
}

func (*azureMSDiagnosticErrorLog) GetLevel() (plog.SeverityNumber, string, bool) {
	_ = "STUB: not implemented"
	// Diagnostic Error logs are always Error level
	return *new(plog.SeverityNumber), "", false
}

func (r *azureMSDiagnosticErrorLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes

type azureMSApplicationMetricsLogProperties struct {
	ApplicationGroupName string `json:"ApplicationGroupName"`
}

func (p *azureMSApplicationMetricsLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// This properties is actually an escaped JSON string,
// so we need to unescape it first

// Define an alias type to avoid infinite recursion

// Assign the unmarshaled fields from the alias to the original struct

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/azmsapplicationmetriclogs
// Available for microsoft.servicebus/namespaces and microsoft.eventhub/namespaces
type azureMSApplicationMetricsLog struct {
	azureMSCommon

	ChildEntityType string                                 `json:"ChildEntityType"`
	ChildEntityName string                                 `json:"ChildEntityName"`
	PartitionID     string                                 `json:"PartitionId"`
	Outcome         string                                 `json:"Outcome"`
	Protocol        string                                 `json:"Protocol"`
	AuthType        string                                 `json:"AuthType"`
	AuthID          string                                 `json:"AuthId"`
	NetworkType     string                                 `json:"NetworkType"`
	ClientIP        string                                 `json:"ClientIp"`
	Count           json.Number                            `json:"Count"` // int
	Properties      azureMSApplicationMetricsLogProperties `json:"Properties"`
}

func (r *azureMSApplicationMetricsLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes
// We will skip "ChildEntityType" and "ChildEntityName" for now,
// as they are not documented and available sample data doesn't provide meaningful values

func (*azureMSApplicationMetricsLog) PutProperties(_ pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	// We will skip "ApplicationGroupName" for now,
	// as they it not documented and available sample data doesn't provide meaningful values
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/azmsoperationallogs
// or https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/service-bus-messaging/monitor-service-bus-reference.md#operational-logs
// Available for microsoft.servicebus/namespaces and microsoft.eventhub/namespaces
type azureMSOperationalLogProperties struct {
	SubscriptionID string `json:"SubscriptionId"`
	Namespace      string `json:"Namespace"`
	ViaURL         string `json:"Via"`
	TrackingID     string `json:"TrackingId"`
	ErrorCode      string `json:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage"`
}

func (p *azureMSOperationalLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// This properties is actually an escaped JSON string,
// so we need to unescape it first

// Define an alias type to avoid infinite recursion

// Assign the unmarshaled fields from the alias to the original struct

type azureMSOperationalLog struct {
	azureMSCommon

	EventName string `json:"EventName"`
	Status    string `json:"Status"`
	Caller    string `json:"Caller"`

	Properties azureMSOperationalLogProperties `json:"EventProperties"`
}

func (r *azureMSOperationalLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes

func (r *azureMSOperationalLog) PutProperties(attrs pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	// SubscriptionId and Namespace are already in top-level attributes, so skip them here
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/azmsruntimeauditlogs
// or https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/service-bus-messaging/monitor-service-bus-reference.md#runtime-audit-logs
// Available for microsoft.servicebus/namespaces and microsoft.eventhub/namespaces
type azureMSRuntimeAuditLog struct {
	azureMSCommon

	TaskName    string      `json:"TaskName"`
	Status      string      `json:"Status"`
	Protocol    string      `json:"Protocol"`
	AuthType    string      `json:"AuthType"`
	AuthID      string      `json:"AuthId"`
	NetworkType string      `json:"NetworkType"`
	ClientIP    string      `json:"ClientIp"`
	Count       json.Number `json:"Count"`      // int
	Properties  string      `json:"Properties"` // unknown structure, save as is
}

func (r *azureMSRuntimeAuditLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes

// Put unparsed properties to log.Body as common approach

type azureMSVNetAndIPFilteringLog struct {
	azureMSCommon

	EventName string      `json:"EventName"`
	IPAddress string      `json:"ipAddress"`
	Action    string      `json:"action"`
	Reason    string      `json:"reason"`
	Count     json.Number `json:"count"` // int
}

func (r *azureMSVNetAndIPFilteringLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes
