// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurefunctionsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

type Config struct {
	// HTTP defines the HTTP server settings for the Azure Functions invoke endpoints.
	HTTP *confighttp.ServerConfig `mapstructure:"http"`

	// Auth is the component.ID of the extension that provides Azure authentication
	Auth component.ID `mapstructure:"auth"`

	// Triggers holds configuration for Azure Functions triggers (e.g. Event Hub)
	Triggers *TriggersConfig `mapstructure:"triggers"`
}

// TriggersConfig groups all supported trigger types for this receiver.
type TriggersConfig struct {
	// EventHub configures the Event Hub trigger: log bindings and their encodings.
	EventHub *EventHubTriggerConfig `mapstructure:"event_hub"`

	_ struct{} // prevent unkeyed literal initialization
}

// EventHubTriggerConfig holds configuration for the Event Hub trigger.
type EventHubTriggerConfig struct {
	// Logs is the list of log bindings (e.g. name "logs" maps to path /logs). Each binding has its own encoding.
	Logs []LogsEncodingConfig `mapstructure:"logs"`
	// IncludeMetadata, when true, adds Azure Functions invoke metadata to resource attributes.
	IncludeMetadata bool `mapstructure:"include_metadata"`
}

// LogsEncodingConfig holds the binding name and encoding for a signal (e.g. one log binding).
// Name is the Azure Functions binding name and typically corresponds to the request path (e.g. /logs, /raw_logs).
type LogsEncodingConfig struct {
	Name     string       `mapstructure:"name"`
	Encoding component.ID `mapstructure:"encoding"`
}

// hasAnyBinding reports whether at least one trigger has at least one binding.
func (t *TriggersConfig) hasAnyBinding() bool { _ = "STUB: not implemented"; return false }

// Validate checks if the receiver configuration is valid.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
