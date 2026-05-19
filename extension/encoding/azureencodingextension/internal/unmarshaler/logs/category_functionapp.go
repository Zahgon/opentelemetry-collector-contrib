// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	// OpenTelemetry attribute name for Event ID
	attributeAzureEventID = "azure.event.id"

	// OpenTelemetry attribute name for Event Name
	attributeAzureEventName = "azure.event.name"
)

// Sometimes Azure sends invalid JSON as a string with a single-quoted keys/values
// for example: `"properties": "{'key':'value'}"`
// instead of `"properties": {"key":"value"}`
// So here a specific wrapper to make this JSON valid
type azureFunctionAppLogProperties struct {
	ActivityID           string      `json:"activityId"`
	AppName              string      `json:"appName"`
	Category             string      `json:"category"`
	EventID              json.Number `json:"eventId"` // int
	EventName            string      `json:"eventName"`
	ExceptionDetails     string      `json:"exceptionDetails"`
	ExceptionMessage     string      `json:"exceptionMessage"`
	ExceptionType        string      `json:"exceptionType"`
	FunctionInvocationID string      `json:"functionInvocationId"`
	FunctionName         string      `json:"functionName"`
	HostInstanceID       string      `json:"hostInstanceId"`
	HostVersion          string      `json:"hostVersion"`
	Level                string      `json:"level"`
	LevelID              json.Number `json:"levelId"` // int
	Message              string      `json:"message"`
	ProcessID            json.Number `json:"processId"` // int
	RoleInstanceID       string      `json:"roleInstance"`
}

func (p *azureFunctionAppLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove leading and trailing double quote

// Define an alias type to avoid infinite recursion

// Assign the unmarshaled fields from the alias to the original struct

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/functionapplogs
// There is no documentation about the structure of the logs, so we will
// implement this based on the fields from Azure Monitor table excluding generic fields
type azureFunctionAppLog struct {
	azureLogRecordBase

	Properties       json.RawMessage                `json:"properties"`
	parsedProperties *azureFunctionAppLogProperties `json:"-"`
	propertiesParsed bool                           `json:"-"`
}

// Override GetResource to add ServiceName and ServiceInstanceID from Properties
func (r *azureFunctionAppLog) GetResource() logsResourceAttributes {
	_ = "STUB: not implemented"
	// Try to parse Properties to get AppName and FunctionName
	return *new(logsResourceAttributes)
}

// We failed to parse "properties" - return basic Resource attributes only

// In general "appName" is naturally matched to "service.name" Resource Attribute,
// but in Azure multiple functions could be deployed into single FunctionApp
// So to make "service.name" correctly identifiable we will use the same approach
// as in SemConv "faas.name" - combine "appName" and "functionName"

// "RoleInstance" field typically matches "AppRoleInstance" in Azure Traces, so we'll assign it to the
// same attribute as "service.instance.id" as in Trace Unmarshaler for correlation purposes

func (r *azureFunctionAppLog) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	// Put some common attributes
	return nil
}

// Try to parse Properties field

// Properties field could not be parsed - put raw string to `azure.properties` attribute

// If we were able to parse Properties - put all known fields to attributes

// According to SemConv this attribute for Azure should be in form `<FUNCAPP>/<FUNC>`
// For clear func name we will use faas.invoked_name attribute

// "exceptionDetails" field typically contains a full stack trace of the exception

func (r *azureFunctionAppLog) getParsedProperties() *azureFunctionAppLogProperties {
	_ = "STUB: not implemented"
	return nil
}
