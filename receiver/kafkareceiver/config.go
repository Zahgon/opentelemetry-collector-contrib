// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/kafka/configkafka"
)

var _ component.Config = (*Config)(nil)

// Config defines configuration for Kafka receiver.
type Config struct {
	configkafka.ClientConfig   `mapstructure:",squash"`
	configkafka.ConsumerConfig `mapstructure:",squash"`

	// Logs holds configuration about how logs should be consumed.
	Logs TopicEncodingConfig `mapstructure:"logs"`

	// Metrics holds configuration about how metrics should be consumed.
	Metrics TopicEncodingConfig `mapstructure:"metrics"`

	// Traces holds configuration about how traces should be consumed.
	Traces TopicEncodingConfig `mapstructure:"traces"`

	// Profiles holds configuration about how profiles should be consumed.
	Profiles TopicEncodingConfig `mapstructure:"profiles"`

	// MessageMarking controls the way the messages are marked as consumed.
	MessageMarking MessageMarking `mapstructure:"message_marking"`

	// HeaderExtraction controls extraction of headers from Kafka records.
	HeaderExtraction HeaderExtraction `mapstructure:"header_extraction"`

	// ErrorBackoff controls backoff/retry behavior when the next consumer
	// returns an error.
	ErrorBackOff configretry.BackOffConfig `mapstructure:"error_backoff"`

	// Telemetry controls optional telemetry configuration.
	Telemetry TelemetryConfig `mapstructure:"telemetry"`
}

func (c *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// Check if deprecated fields have been explicitly set
// give them  precedence

// Set OnPermanentError default value to inherit from OnError for backward compatibility
// Only if OnPermanentError was not explicitly set in the config

// If message_marking section doesn't exist, set defaults

// Validate checks the receiver configuration is valid.
func (c *Config) Validate() error {
	_ = "STUB: not implemented"
	// Validate that exclude_topic is only used with regex topic patterns
	return nil
}

// validateExcludeTopic checks that exclude_topic is only configured when topics uses regex pattern
func validateExcludeTopic(signalType string, topics, excludeTopics []string) error {
	_ = "STUB: not implemented"
	return nil
}

// No exclude_topic configured, nothing to validate

// if none of the configured topic uses regex return error

// Validate that exclude_topic is not empty

// Validate that exclude_topic is a valid regex pattern

// TopicEncodingConfig holds signal-specific topic and encoding configuration.
type TopicEncodingConfig struct {
	// Topics holds the name of the Kafka topics from which messages of the
	// signal type should be consumed.
	//
	// The default depends on the signal type:
	//  - "otlp_spans" for traces
	//  - "otlp_metrics" for metrics
	//  - "otlp_logs" for logs
	//  - "otlp_profiles" for profiles
	Topics []string `mapstructure:"topics"`

	// Encoding holds the expected encoding of messages for the signal type
	//
	// Defaults to "otlp_proto".
	Encoding string `mapstructure:"encoding"`

	// Optional exclude topics option, used only in regex mode.
	ExcludeTopics []string `mapstructure:"exclude_topics"`
}

type MessageMarking struct {
	// If true, the messages are marked after the pipeline execution
	After bool `mapstructure:"after"`

	// If false, only the successfully processed messages are marked. This only applies
	// to non-permanent errors. It has no impact if After is set to false.
	// Note: this can block the entire partition in case a message processing returns
	// a non-permanent error.
	OnError bool `mapstructure:"on_error"`

	// If false, only the successfully processed messages are marked. This only applies
	// to permanent errors. It has no impact if After is set to false.
	// Default value inherits from OnError for backward compatibility.
	// Note: this can block the entire partition in case a message processing returns
	// a permanent error.
	OnPermanentError bool `mapstructure:"on_permanent_error"`
}

type HeaderExtraction struct {
	ExtractHeaders bool     `mapstructure:"extract_headers"`
	Headers        []string `mapstructure:"headers"`
}

type TelemetryConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
	_       struct{}      // avoids unkeyed_literal_initialization
}

// MetricsConfig provides config for optional receiver metrics.
type MetricsConfig struct {
	// KafkaReceiverRecordsDelay controls whether the metric kafka_receiver_records_delay
	// that measures the time in seconds between producing and receiving a batch of records
	// will be reported or not. This metric is not reported by default because
	// it may slow down high-volume consuming.
	KafkaReceiverRecordsDelay MetricConfig `mapstructure:"kafka_receiver_records_delay"`
	_                         struct{}     // avoids unkeyed_literal_initialization
}

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool     `mapstructure:"enabled"`
	_       struct{} // avoids unkeyed_literal_initialization
}
