// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"go.opentelemetry.io/collector/pdata/pcommon" // Applies resource attributes values to data properties
)

const (
	instrumentationLibraryName    string = "instrumentationlibrary.name"
	instrumentationLibraryVersion string = "instrumentationlibrary.version"
)

// Applies resource attributes values to data properties
func applyResourcesToDataProperties(dataProperties map[string]string, resourceAttributes pcommon.Map) {
	_ = "STUB: not implemented"
	// Copy all the resource labels into the base data properties. Resource values are always strings
	return
}

// Sets important ai.cloud.* tags on the envelope
func applyCloudTagsToEnvelope(envelope *contracts.Envelope, resourceAttributes pcommon.Map) {
	_ = "STUB: not implemented"
	// Extract key service.* labels from the Resource labels and construct CloudRole and CloudRoleInstance envelope tags
	// https://github.com/open-telemetry/opentelemetry-specification/tree/main/specification/resource/semantic_conventions
	return
}

// Sets ai.application.* tags on the envelope
func applyApplicationTagsToEnvelope(envelope *contracts.Envelope, resourceAttributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// Sets ai.device.* tags on the envelope
func applyDeviceTagsToEnvelope(envelope *contracts.Envelope, resourceAttributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// Applies internal sdk version tag on the envelope
func applyInternalSdkVersionTagToEnvelope(envelope *contracts.Envelope) {
	_ = "STUB: not implemented"
	return
}

// Applies instrumentation values to data properties
func applyInstrumentationScopeValueToDataProperties(dataProperties map[string]string, instrumentationScope pcommon.InstrumentationScope) {
	_ = "STUB: not implemented"
	// Copy the instrumentation properties
	return
}

// Applies attributes as Application Insights properties
func setAttributesAsProperties(attributeMap pcommon.Map, properties map[string]string) {
	_ = "STUB: not implemented"
	return
}
