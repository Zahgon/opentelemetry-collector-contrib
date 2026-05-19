// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/traces"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

// categoryHolder is a small helper struct to get Category (`type`) fields
// from each Trace Record
type categoryHolder struct {
	Category string `json:"type"` // Used in AppInsights instead of `category` field
}

// traceResourceAttributes is a helper struct to hold resource attributes
type traceResourceAttributes struct {
	ResourceID      string
	ServiceName     string
	ServiceInstance string
	ServiceVersion  string
	SDKVersion      string
	Location        string
}

type ResourceTracesUnmarshaler struct {
	buildInfo  component.BuildInfo
	logger     *zap.Logger
	timeFormat []string
}

func (r ResourceTracesUnmarshaler) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
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

// 6 pre-defined attributes + 2 for SDK name/version

// SDK Name and Version parsing

// Unparsable SDK string - save as is

func (r ResourceTracesUnmarshaler) unmarshalRecord(allResourceScopeSpans map[traceResourceAttributes]ptrace.ScopeSpans, record []byte) {
	_ = "STUB: not implemented"
	// That's actually double-unmarshaling, but there is no other way to parse variety of Azure Resource schemas
	return
}

// We couldn't do any SemConv conversion if Schema Category is not defined

// Unsupported Category Type - skip processing with Warning

// Let's parse it

// Get timestamp for start time (and end time calculation)

// Sometimes, ResourceID is not enough to uniquely identify a resource in Azure cloud
// For example, multiple Azure Functions deployed into single Azure Functions App,
// will be sharing the same ResourceID,
// but actually it's different Resources with own ServiceName, ServiceInstance and ServiceVersion.

// Grouping set of traces by Resource Attributes

// Populate Span data

// Populate Span Attributes

func NewAzureResourceTracesUnmarshaler(buildInfo component.BuildInfo, logger *zap.Logger, cfg TracesConfig) ResourceTracesUnmarshaler {
	_ = "STUB: not implemented"
	return *new(ResourceTracesUnmarshaler)
}
