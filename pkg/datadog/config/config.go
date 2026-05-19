// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"

import (
	"errors"
	"regexp"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap"
)

var (
	// ErrUnsetAPIKey is returned when the API key is not set.
	ErrUnsetAPIKey = errors.New("api.key is not set")
	// ErrNoMetadata is returned when only_metadata is enabled but host metadata is disabled or hostname_source is not first_resource.
	ErrNoMetadata = errors.New("only_metadata can't be enabled when host_metadata::enabled = false or host_metadata::hostname_source != first_resource")
	// ErrInvalidHostname is returned when the hostname is invalid.
	ErrEmptyEndpoint = errors.New("endpoint cannot be empty")
	// NonHexRegex is a regex of characters that are always invalid in a Datadog API key
	NonHexRegex = regexp.MustCompile(NonHexChars)
)

const (
	// DefaultSite is the default site of the Datadog intake to send data to
	DefaultSite = "datadoghq.com"
	// NonHexChars is a regex of characters that are always invalid in a Datadog API key
	NonHexChars = "[^0-9a-fA-F]"
)

// APIConfig defines the API configuration options
type APIConfig struct {
	// Key is the Datadog API key to associate your Agent's data with your organization.
	// Create a new API key here: https://app.datadoghq.com/account/settings
	Key configopaque.String `mapstructure:"key"`

	// Site is the site of the Datadog intake to send data to.
	// The default value is "datadoghq.com".
	Site string `mapstructure:"site"`

	// FailOnInvalidKey states whether to exit at startup on invalid API key.
	// The default value is false.
	FailOnInvalidKey bool `mapstructure:"fail_on_invalid_key"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// TagsConfig defines the tag-related configuration
// It is embedded in the configuration
type TagsConfig struct {
	// Hostname is the fallback hostname used for payloads without hostname-identifying attributes.
	// This option will NOT change the hostname applied to your metrics, traces and logs if they already have hostname-identifying attributes.
	// If unset, the hostname will be determined automatically. See https://docs.datadoghq.com/opentelemetry/schema_semantics/hostname/?tab=datadogexporter#fallback-hostname-logic for details.
	//
	// Prefer using the `datadog.host.name` resource attribute over using this setting.
	// See https://docs.datadoghq.com/opentelemetry/schema_semantics/hostname/?tab=datadogexporter#general-hostname-semantic-conventions for details.
	Hostname string `mapstructure:"hostname"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// Config defines configuration for the Datadog exporter.
type Config struct {
	confighttp.ClientConfig   `mapstructure:",squash"`                                 // squash ensures fields are correctly decoded in embedded struct.
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`

	TagsConfig `mapstructure:",squash"`

	// API defines the Datadog API configuration.
	API APIConfig `mapstructure:"api"`

	// Metrics defines the Metrics exporter specific configuration
	Metrics MetricsConfig `mapstructure:"metrics"`

	// Traces defines the Traces exporter specific configuration
	Traces TracesExporterConfig `mapstructure:"traces"`

	// Logs defines the Logs exporter specific configuration
	Logs LogsConfig `mapstructure:"logs"`

	// HostMetadata defines the host metadata specific configuration
	HostMetadata HostMetadataConfig `mapstructure:"host_metadata"`

	// HostnameDetectionTimeout defines the timeout for hostname detection.
	// This is necessary for initializing datadog exporter
	// On K8s, it must be set to less than `failureThreshold * periodSeconds` due to
	// initialization blocking health_check liveness probes on startup.
	// If set to zero duration, there will be no timeout applied.
	// Default is 25 seconds.
	HostnameDetectionTimeout time.Duration `mapstructure:"hostname_detection_timeout"`

	// OnlyMetadata defines whether to only send metadata
	// This is useful for agent-collector setups, so that
	// metadata about a host is sent to the backend even
	// when telemetry data is reported via a different host.
	//
	// This flag is incompatible with disabling host metadata,
	// `use_resource_metadata`, or `host_metadata::hostname_source != first_resource`
	OnlyMetadata bool `mapstructure:"only_metadata"`

	OrchestratorExplorer OrchestratorExplorerConfig `mapstructure:"orchestrator_explorer"`

	// Non-fatal warnings found during configuration loading.
	warnings []error
}

// LogWarnings logs warning messages that were generated on unmarshaling.
func (c *Config) LogWarnings(logger *zap.Logger) { _ = "STUB: not implemented"; return }

// AddWarning adds a warning message to the configuration.
// This allows external modules to add warnings that will be logged later.
func (c *Config) AddWarning(warning error) { _ = "STUB: not implemented"; return }

// AddWarningf adds a formatted warning message to the configuration.
// This allows external modules to add formatted warnings that will be logged later.
func (c *Config) AddWarningf(format string, args ...any) { _ = "STUB: not implemented"; return }

// GetWarnings returns a copy of all warnings stored in the configuration.
// This allows external modules to retrieve and process warnings as needed.
func (c *Config) GetWarnings() []error {
	_ = "STUB: not implemented"
	// Return a copy to prevent external modification of the internal slice
	return nil
}

var _ component.Config = (*Config)(nil)

// Validate the configuration for errors. This is required by component.Config.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// StaticAPIKey Check checks if api::key is either empty or contains invalid (non-hex) characters
// It does not validate online; this is handled on startup.
//
// Deprecated: [v0.136.0] Do not use, will be removed on the next minor version
func StaticAPIKeyCheck(key string) error { _ = "STUB: not implemented"; return nil }

func validateClientConfig(cfg confighttp.ClientConfig) error { _ = "STUB: not implemented"; return nil }

var _ error = (*renameError)(nil)

// renameError is an error related to a renamed setting.
type renameError struct {
	// oldName of the configuration option.
	oldName string
	// newName of the configuration option.
	newName string
	// issueNumber on opentelemetry-collector-contrib for tracking
	issueNumber uint
}

// List of settings that have been removed, but for which we keep a custom error.
var removedSettings = []renameError{
	{
		oldName:     "metrics::send_monotonic_counter",
		newName:     "metrics::sums::cumulative_monotonic_mode",
		issueNumber: 8489,
	},
	{
		oldName:     "tags",
		newName:     "host_metadata::tags",
		issueNumber: 9099,
	},
	{
		oldName:     "send_metadata",
		newName:     "host_metadata::enabled",
		issueNumber: 9099,
	},
	{
		oldName:     "use_resource_metadata",
		newName:     "host_metadata::hostname_source",
		issueNumber: 9099,
	},
	{
		oldName:     "metrics::report_quantiles",
		newName:     "metrics::summaries::mode",
		issueNumber: 8845,
	},
	{
		oldName:     "metrics::instrumentation_library_metadata_as_tags",
		newName:     "metrics::instrumentation_scope_as_tags",
		issueNumber: 11135,
	},
}

// Error implements the error interface.
func (e renameError) Error() string { _ = "STUB: not implemented"; return "" }

func handleRemovedSettings(configMap *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

var _ confmap.Unmarshaler = (*Config)(nil)

// Unmarshal a configuration map into the configuration struct.
func (c *Config) Unmarshal(configMap *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// Add deprecation warnings for deprecated settings.

// If an endpoint is not explicitly set, override it based on the site.

// Return an error if an endpoint is explicitly set to ""

func defaultClientConfig() confighttp.ClientConfig {
	_ = "STUB: not implemented"
	return *new(confighttp.ClientConfig)
}

// CreateDefaultConfig creates the default exporter configuration
func CreateDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// set to 25 to prevent 30-second pod restart on K8s as reported in issue #40372 and #40373

// CheckAndCastConfig checks a component.Config type and casts it to the Datadog Config struct.
func CheckAndCastConfig(c component.Config) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
