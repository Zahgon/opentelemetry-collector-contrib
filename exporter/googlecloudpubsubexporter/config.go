// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"

import (
	"regexp"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

var topicMatcher = regexp.MustCompile(`^projects/[a-z][a-z0-9\-]*/topics/`)

type Config struct {
	// Timeout for all API calls. If not set, defaults to 12 seconds.
	TimeoutSettings           exporterhelper.TimeoutConfig                             `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	// Google Cloud Project ID where the Pubsub client will connect to
	ProjectID string `mapstructure:"project"`
	// User agent that will be used by the Pubsub client to connect to the service
	UserAgent string `mapstructure:"user_agent"`
	// Override of the Pubsub Endpoint, leave empty for the default endpoint
	Endpoint string `mapstructure:"endpoint"`
	// Only has effect if Endpoint is not ""
	Insecure bool `mapstructure:"insecure"`

	// The fully qualified resource name of the Pubsub topic
	Topic string `mapstructure:"topic"`
	// Compression of the payload (only gzip or is supported, no compression is the default)
	Compression string `mapstructure:"compression"`
	// Watermark defines the watermark (the ce-time attribute on the message) behavior
	Watermark WatermarkConfig `mapstructure:"watermark"`
	// Ordering configures the ordering keys
	Ordering OrderingConfig `mapstructure:"ordering"`
	// LogsSignalConfig allows for custom log configuration
	LogsSignalConfig SignalConfig `mapstructure:"logs"`
	// MetricsSignalConfig allows for custom log configuration
	MetricsSignalConfig SignalConfig `mapstructure:"metrics"`
	// TracesSignalConfig allows for custom log configuration
	TracesSignalConfig SignalConfig `mapstructure:"traces"`
}

// WatermarkConfig customizes the behavior of the watermark
type WatermarkConfig struct {
	// Behavior of the watermark. Currently, only of the message (none, earliest and current, current being the default)
	// will set the timestamp on pubsub based on timestamps of the events inside the message
	Behavior string `mapstructure:"behavior"`
	// Indication on how much the timestamp can drift from the current time, the timestamp will be capped to the allowed
	// maximum. A duration of 0 is the same as maximum duration
	AllowedDrift time.Duration `mapstructure:"allowed_drift"`
}

// OrderingConfig customizes the behavior of the ordering
type OrderingConfig struct {
	// Enabled indicates if ordering is enabled
	Enabled bool `mapstructure:"enabled"`
	// FromResourceAttribute is a resource attribute that will be used as the ordering key.
	FromResourceAttribute string `mapstructure:"from_resource_attribute"`
	// RemoveResourceAttribute indicates if the ordering key should be removed from the resource attributes.
	RemoveResourceAttribute bool `mapstructure:"remove_resource_attribute"`
}

// SignalConfig holds signal-specific configuration for the Kafka exporter.
type SignalConfig struct {
	// Encoding is a custom encoding for the marshaling the data onto the message
	Encoding component.ID `mapstructure:"encoding"`
	// Attributes are custom Pub/Sub message attributes
	Attributes map[string]string `mapstructure:"attributes"`
}

func (config *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (config *WatermarkConfig) validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *OrderingConfig) validate() error { _ = "STUB: not implemented"; return nil }

func (config *Config) parseCompression() (compression, error) {
	_ = "STUB: not implemented"
	return *new(compression), nil
}

func (config *WatermarkConfig) parseWatermarkBehavior() (WatermarkBehavior, error) {
	_ = "STUB: not implemented"
	return *new(WatermarkBehavior), nil
}
