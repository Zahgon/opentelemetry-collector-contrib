// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/traces"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appavailabilityresults
type azureAppAvailabilityResults struct {
	azureTracesRecordBase

	// AppAvailabilityResults only fields
	Location string `json:"Location"` // The location from where the test ran
	Message  string `json:"Message"`  // Application-defined message
}

// GetResource returns resource attributes for the parsed Trace Record,
// adding cloud.region from `Location` field
func (r *azureAppAvailabilityResults) GetResource() traceResourceAttributes {
	_ = "STUB: not implemented"
	return *new(traceResourceAttributes)
}

// GetSpanKind determines the SpanKind
func (*azureAppAvailabilityResults) GetSpanKind() ptrace.SpanKind {
	_ = "STUB: not implemented"
	// As we unsure about AvailabilityResults SpanKind - set it to Internal
	return *new(ptrace.SpanKind)
}

// GetSpanStatus returns Span Status Code and optional Status Message,
// based on `Success` field
func (r *azureAppAvailabilityResults) GetSpanStatus() (ptrace.StatusCode, string) {
	_ = "STUB: not implemented"
	// According to Azure Docs if `Success` is false - the operation failed,
	// so we'll mark Span Status as Error for such cases according to OpenTelemetry Specs
	return *new(ptrace.StatusCode), ""
}

// In all other cases - return Unset Status Code as recommended by OpenTelemetry Specs
