// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver"

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// filterDimensions transforms a list of azure dimensions into a list of string, taking in account the DimensionConfig
// given by the user.
func filterDimensions(dimensions []*armmonitor.LocalizableString, cfg DimensionsConfig, resourceType, metricName string) []string {
	_ = "STUB: not implemented"
	// Only skip if explicitly disabled. Enabled by default.
	return nil
}

// If dimensions are overridden for that resource type and metric name, we take it

// Otherwise we get all dimensions

// serializeDimensions build a comma separated string from trimmed, sorted dimensions list.
// It is designed to be used as a key in scraper maps.
func serializeDimensions(dimensions []string) string { _ = "STUB: not implemented"; return "" }

// buildDimensionsFilter takes a serialized dimensions input to build an Azure Request filter that will allow us to
// receive metrics values split by these dimensions.
func buildDimensionsFilter(dimensionsStr string) *string { _ = "STUB: not implemented"; return nil }
