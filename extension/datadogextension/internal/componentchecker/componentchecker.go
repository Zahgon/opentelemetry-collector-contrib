// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// package componentchecker will define the functions and types necessary to parse component status and config components
package componentchecker // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/componentchecker"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/service"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/payload"
)

const (
	receiverKind  = "receiver"
	processorKind = "processor"
	exporterKind  = "exporter"
	extensionKind = "extension"
	connectorKind = "connector"
	providerKind  = "provider"
	converterKind = "converter"
	pipelinesKind = "pipelines"
)

func isComponentConfigured(typ component.Type, c map[component.ID]component.Config) bool {
	_ = "STUB: not implemented"
	// Check if at least one of the component type is configured in the config map
	// Note, there could be more than one component of the same type, this function
	// only guarantees that at least one component of the type is configured
	// in the config map.
	return false
}

// DataToFlattenedJSONString is a helper function to ensure payload strings are
// properly formatted for JSON parsing. This is necessary due to escaped newline
// characters and whitespace causing failure on parsing when loading into
// the underlying data platform.
func DataToFlattenedJSONString(data any) string { _ = "STUB: not implemented"; return "" }

// PopulateFullComponentsJSON creates a ModuleInfoJSON struct with all components from ModuleInfos
func PopulateFullComponentsJSON(moduleInfo service.ModuleInfos, c *confmap.Conf) (*payload.ModuleInfoJSON, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: add Providers and Converters after upstream change accepted to add these to moduleinfos

// PopulateActiveComponents gets a list of active components in the collector service
// configuration and returns a list of ServiceComponent structs for inclusion in a fleet payload
func PopulateActiveComponents(logger *zap.Logger, c *confmap.Conf, moduleInfoJSON *payload.ModuleInfoJSON) (*[]payload.ServiceComponent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process extensions

// This extension may be a custom component from the datadog-agent repository (i.e. ddflareextension)

// TODO: Add component status parsing, potentially via pkg/status

// Define a struct to generalize processing of pipeline components

// Define the pipeline components to process

// Connectors can act as both receivers and exporters, so we need to handle them separately

// Process pipelines

// This check exists solely for Connectors; since getIDs returns all receivers and exporters,
// we only want to match the components that are actually connectors.

// TODO: Add component status parsing, potentially via pkg/status
