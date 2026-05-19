// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/traces"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	// OpenTelemetry attribute name for Azure detailed information about the dependency call,
	// from `Data` field in Azure Trace Record
	attributeAzureDependencyData = "azure.dependency.data"

	// OpenTelemetry attribute name for Azure dependency type, such as "HTTP" or "SQL",
	// from `DependencyType` field in Azure Trace Record
	attributeAzureDependencyType = "azure.dependency.type"

	// OpenTelemetry attribute name for Azure target of a dependency call, such as a Web or a SQL server name,
	// from `Target` field in Azure Trace Record
	attributeAzureDependencyTarget = "azure.dependency.target"
)

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appdependencies
type azureAppDependencies struct {
	azureTracesRecordBase

	// AppDependencies only fields
	ResultCode     *json.Number `json:"ResultCode"`     // Result code returned by or to the application
	Data           string       `json:"Data"`           // Detailed information about the dependency call, such as a full URI or a SQL statement
	DependencyType string       `json:"DependencyType"` // Dependency type, such as HTTP or SQL
	Target         string       `json:"Target"`         // Target of a dependency call, such as a Web or a SQL server name
}

// GetSpanKind determines the SpanKind based on the Azure Dependency Type
func (r *azureAppDependencies) GetSpanKind() ptrace.SpanKind {
	_ = "STUB: not implemented"
	// `DependencyType` == "Queue Message | <name of messaging system>" => Producer Span Kind
	return *new(ptrace.SpanKind)
}

// `Data` field is set - we can assume that it's "Client" SpanKind,
// as according to Azure docs - `Data` field contains URI/DB Statement

// By default - returns "Internal" Span Kind as of OpenTelemetry SemConv Spec

// PutCommonAttributes puts already parsed common attributes into provided Attributes Map/Body
func (r *azureAppDependencies) PutCommonAttributes(attrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// AppDependencies doesn't have "URL" field, it is using "Data" field instead

// Otherwise, we will save Data field as is
