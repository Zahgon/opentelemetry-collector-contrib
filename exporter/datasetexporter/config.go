// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasetexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"

import (
	"time"

	datasetConfig "github.com/scalyr/dataset-go/pkg/config"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	exportSeparatorDefault     = "."
	exportDistinguishingSuffix = "_"
)

// exportSettings configures separator and distinguishing suffixes for all exported fields
type exportSettings struct {
	// ExportSeparator is separator used when flattening exported attributes
	// Default value: .
	ExportSeparator string `mapstructure:"export_separator"`

	// ExportDistinguishingSuffix is suffix used to be appended to the end of attribute name in case of collision
	// Default value: _
	ExportDistinguishingSuffix string `mapstructure:"export_distinguishing_suffix"`
}

// newDefaultExportSettings returns the default settings for exportSettings.
func newDefaultExportSettings() exportSettings {
	_ = "STUB: not implemented"
	return *new(exportSettings)
}

type TracesSettings struct {
	// exportSettings configures separator and distinguishing suffixes for all exported fields
	exportSettings `mapstructure:",squash"`
}

// newDefaultTracesSettings returns the default settings for TracesSettings.
func newDefaultTracesSettings() TracesSettings {
	_ = "STUB: not implemented"
	return *new(TracesSettings)
}

const (
	logsExportResourceInfoDefault                  = false
	logsExportResourcePrefixDefault                = "resource.attributes."
	logsExportScopeInfoDefault                     = true
	logsExportScopePrefixDefault                   = "scope.attributes."
	logsDecomposeComplexMessageFieldDefault        = false
	logsDecomposedComplexMessageFieldPrefixDefault = "body.map."
)

type LogsSettings struct {
	// ExportResourceInfo is optional flag to signal that the resource info is being exported to DataSet while exporting Logs.
	// This is especially useful when reducing DataSet billable log volume.
	// Default value: false
	ExportResourceInfo bool `mapstructure:"export_resource_info_on_event"`

	// ExportResourcePrefix is prefix for the resource attributes when they are exported (see ExportResourceInfo).
	// Default value: resource.attributes.
	ExportResourcePrefix string `mapstructure:"export_resource_prefix"`

	// ExportScopeInfo is an optional flag that signals if scope info should be exported (when available) with each event. If scope
	// information is not utilized, it makes sense to disable exporting it since it will result in increased billable log volume.
	// Default value: true
	ExportScopeInfo bool `mapstructure:"export_scope_info_on_event"`

	// ExportScopePrefix is prefix for the scope attributes when they are exported (see ExportScopeInfo).
	// Default value: scope.attributes.
	ExportScopePrefix string `mapstructure:"export_scope_prefix"`

	// DecomposeComplexMessageField is an optional flag to signal that message / body of complex types (e.g. a map) should be
	// decomposed / deconstructed into multiple fields. This is usually done outside of the main DataSet integration on the
	// client side (e.g. as part of the attribute processor or similar) or on the server side (DataSet server side JSON parser
	// for message field) and that's why this functionality is disabled by default.
	DecomposeComplexMessageField bool `mapstructure:"decompose_complex_message_field"`

	// DecomposedComplexMessagePrefix is prefix for the decomposed complex message (see DecomposeComplexMessageField).
	// Default value: body.map.
	DecomposedComplexMessagePrefix string `mapstructure:"decomposed_complex_message_prefix"`

	// exportSettings configures separator and distinguishing suffixes for all exported fields
	exportSettings `mapstructure:",squash"`
}

// newDefaultLogsSettings returns the default settings for LogsSettings.
func newDefaultLogsSettings() LogsSettings { _ = "STUB: not implemented"; return *new(LogsSettings) }

const (
	bufferMaxLifetime          = 5 * time.Second
	bufferPurgeOlderThan       = 30 * time.Second
	bufferRetryInitialInterval = 5 * time.Second
	bufferRetryMaxInterval     = 30 * time.Second
	bufferRetryMaxElapsedTime  = 300 * time.Second
	bufferRetryShutdownTimeout = 30 * time.Second
	bufferMaxParallelOutgoing  = 100
)

type BufferSettings struct {
	MaxLifetime          time.Duration `mapstructure:"max_lifetime"`
	PurgeOlderThan       time.Duration `mapstructure:"purge_older_than"`
	GroupBy              []string      `mapstructure:"group_by"`
	RetryInitialInterval time.Duration `mapstructure:"retry_initial_interval"`
	RetryMaxInterval     time.Duration `mapstructure:"retry_max_interval"`
	RetryMaxElapsedTime  time.Duration `mapstructure:"retry_max_elapsed_time"`
	RetryShutdownTimeout time.Duration `mapstructure:"retry_shutdown_timeout"`
	MaxParallelOutgoing  int           `mapstructure:"max_parallel_outgoing"`
}

// newDefaultBufferSettings returns the default settings for BufferSettings.
func newDefaultBufferSettings() BufferSettings {
	_ = "STUB: not implemented"
	return *new(BufferSettings)
}

type ServerHostSettings struct {
	UseHostName bool   `mapstructure:"use_hostname"`
	ServerHost  string `mapstructure:"server_host"`
}

// newDefaultBufferSettings returns the default settings for BufferSettings.
func newDefaultServerHostSettings() ServerHostSettings {
	_ = "STUB: not implemented"
	return *new(ServerHostSettings)
}

const debugDefault = false

type Config struct {
	DatasetURL                string              `mapstructure:"dataset_url"`
	APIKey                    configopaque.String `mapstructure:"api_key"`
	Debug                     bool                `mapstructure:"debug"`
	BufferSettings            `mapstructure:"buffer"`
	TracesSettings            `mapstructure:"traces"`
	LogsSettings              `mapstructure:"logs"`
	ServerHostSettings        `mapstructure:"server_host"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	TimeoutSettings           exporterhelper.TimeoutConfig                             `mapstructure:"timeout"`
}

func (c *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// Validate checks if all required fields in Config are set and have valid values.
// If any of the required fields are missing or have invalid values, it returns an error.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// String returns a string representation of the Config object.
// It includes all the fields and their values in the format "field_name: field_value".
func (c *Config) String() string { _ = "STUB: not implemented"; return "" }

func (c *Config) convert() *exporterConfig { _ = "STUB: not implemented"; return nil }

type exporterConfig struct {
	datasetConfig      *datasetConfig.DataSetConfig
	tracesSettings     TracesSettings
	logsSettings       LogsSettings
	serverHostSettings ServerHostSettings
}
