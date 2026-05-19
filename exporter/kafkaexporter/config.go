// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/kafkaclient"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/kafka/configkafka"
)

var _ component.Config = (*Config)(nil)

var (
	errRecordPartitionerMultipleSet = errors.New("at most one record_partitioner strategy may be configured")
	errRecordPartitionerMissing     = errors.New("no partitioner type configured")
)

var errLogsPartitionExclusive = errors.New(
	"partition_logs_by_resource_attributes and partition_logs_by_trace_id cannot both be enabled",
)

var (
	errTracesMessageKeyExclusive        = errors.New("traces::message_key_from_metadata_key cannot be combined with partition_traces_by_id")
	errMetricsMessageKeyExclusive       = errors.New("metrics::message_key_from_metadata_key cannot be combined with partition_metrics_by_resource_attributes")
	errLogsMessageKeyExclusive          = errors.New("logs::message_key_from_metadata_key cannot be combined with partition_logs_by_resource_attributes or partition_logs_by_trace_id")
	errMessageKeyMetadataKeyNotIncluded = errors.New("message_key_from_metadata_key must be present in sending_queue::batch::partition::metadata_keys if batching is enabled")
)

var (
	errTopicMetadataKeyNotIncluded        = errors.New("topic_from_metadata_key must be present in sending_queue::batch::partition::metadata_keys if batching is enabled")
	errBatchPartitionMetadataKeysRequired = errors.New("sending_queue::batch::partition::metadata_keys must be configured when include_metadata_keys is set and batching is enabled")
	errIncludeMetadataKeysNotPartitioned  = errors.New("sending_queue::batch::partition::metadata_keys must include all include_metadata_keys values")
)

const (
	HasherSaramaCompat = "sarama_compat"
	HasherMurmur2      = "murmur2"
)

// RecordPartitionerConfig configures the strategy used to assign Kafka records to partitions.
// At most one field should be set.
type RecordPartitionerConfig struct {
	// StickyKey uses StickyKeyPartitioner.
	// When a record key is set, the partition is derived from the key hash.
	StickyKey *StickyKeyPartitionerConfig `mapstructure:"sticky_key"`

	// RoundRobin distributes records evenly across all available partitions in round-robin order.
	RoundRobin *struct{} `mapstructure:"round_robin"`

	// LeastBackup routes each record to the partition with the fewest buffered records.
	LeastBackup *struct{} `mapstructure:"least_backup"`

	// Extension is the component ID of an extension implementing RecordPartitionerExtension.
	// Setting this field delegates partition assignment to that extension.
	Extension *component.ID `mapstructure:"extension"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// StickyKeyPartitionerConfig configures the StickyKeyPartitioner.
type StickyKeyPartitionerConfig struct {
	// Hasher is the hash algorithm used for key-based partition assignment.
	// Valid values: "sarama_compat" (default).
	//   - "sarama_compat": Sarama-compatible FNV-1a hashing (SaramaCompatHasher).
	//   - "murmur2": Murmur2 hashing.
	Hasher string `mapstructure:"hasher"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *StickyKeyPartitionerConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *RecordPartitionerConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *RecordPartitionerConfig) Unmarshal(conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// no partitioner configured, will use default.

// Config defines configuration for Kafka exporter.
type Config struct {
	TimeoutSettings           exporterhelper.TimeoutConfig                             `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	QueueBatchConfig          configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	configkafka.ClientConfig  `mapstructure:",squash"`
	Producer                  configkafka.ProducerConfig `mapstructure:"producer"`

	// Logs holds configuration about how logs should be sent to Kafka.
	Logs SignalConfig `mapstructure:"logs"`

	// Metrics holds configuration about how metrics should be sent to Kafka.
	Metrics SignalConfig `mapstructure:"metrics"`

	// Traces holds configuration about how traces should be sent to Kafka.
	Traces SignalConfig `mapstructure:"traces"`

	// Profiles holds configuration about how profiles should be sent to Kafka.
	Profiles SignalConfig `mapstructure:"profiles"`

	// IncludeMetadataKeys indicates the receiver's client metadata keys to propagate as Kafka message headers.
	IncludeMetadataKeys []string `mapstructure:"include_metadata_keys"`

	// RecordHeaders sets static headers on every outgoing Kafka record.
	RecordHeaders []kafkaclient.RecordHeader `mapstructure:"record_headers"`

	// TopicFromAttribute is the name of the attribute to use as the topic name.
	TopicFromAttribute string `mapstructure:"topic_from_attribute"`

	// PartitionTracesByID sets the message key of outgoing trace messages to the trace ID.
	//
	// NOTE: this does not have any effect for Jaeger encodings. Jaeger encodings always use
	// use the trace ID for the message key.
	PartitionTracesByID bool `mapstructure:"partition_traces_by_id"`

	// PartitionMetricsByResourceAttributes controls the partitioning of metrics messages by
	// resource. If this is true, then the message key will be set to a hash of the resource's
	// identifying attributes.
	PartitionMetricsByResourceAttributes bool `mapstructure:"partition_metrics_by_resource_attributes"`

	// PartitionLogsByResourceAttributes controls the partitioning of logs messages by resource.
	// If this is true, then the message key will be set to a hash of the resource's identifying
	// attributes.
	PartitionLogsByResourceAttributes bool `mapstructure:"partition_logs_by_resource_attributes"`

	// PartitionLogsByTraceID controls partitioning of log messages by trace ID only.
	// When enabled, the exporter splits incoming logs per TraceID (using SplitLogs)
	// and sets the Kafka message key to the 16-byte hex string of that TraceID.
	// If a LogRecord has an empty TraceID, the key may be empty and partition
	// selection falls back to the Kafka client’s default strategy. Resource
	// attributes are not used for the key when this option is enabled.
	PartitionLogsByTraceID bool `mapstructure:"partition_logs_by_trace_id"`

	// RecordPartitioner configures how Kafka records are assigned to partitions.
	// The default ("sarama_compatible") retains the legacy Sarama-compatible hashing
	// behavior. Set to "sticky", "round_robin", or "least_backup" to use one of the
	// built-in franz-go partitioners, or "extension" to delegate to a custom extension.
	RecordPartitioner RecordPartitionerConfig `mapstructure:"record_partitioner"`
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// SignalConfig holds signal-specific configuration for the Kafka exporter.
type SignalConfig struct {
	// Topic holds the name of the Kafka topic to which messages of the
	// signal type should be produced.
	//
	// The default depends on the signal type:
	//  - "otlp_spans" for traces
	//  - "otlp_metrics" for metrics
	//  - "otlp_logs" for logs
	//  - "otlp_profiles" for profiles
	Topic string `mapstructure:"topic"`

	// TopicFromMetadataKey holds the name of the metadata key to use as the
	// topic name for this signal type. If this is set, it takes precedence
	// over the topic name set in the topic field.
	TopicFromMetadataKey string `mapstructure:"topic_from_metadata_key"`

	// MessageKeyFromMetadataKey holds the name of the metadata key whose value
	// will be used as the Kafka record key for this signal type. If the metadata
	// key is absent or empty the record key is left nil.
	// Mutually exclusive with the partition_* flags for the same signal.
	MessageKeyFromMetadataKey string `mapstructure:"message_key_from_metadata_key"`

	// Encoding holds the encoding of messages for the signal type.
	//
	// Defaults to "otlp_proto".
	Encoding string `mapstructure:"encoding"`
}

// validateBatchPartitionerKeys validates the partition keys if sending_queue::batch is enabled.
// The exporter relies on a few client metadata keys to be present, if configured, in the final
// batch that needs to be exported, however, since batching removes all client metadata keys by
// default we need to ensure proper partitioning is configured to keep the required metadata.
func validateBatchPartitionerKeys(c *Config) error { _ = "STUB: not implemented"; return nil }

// Validate if include_metadata_keys are included in partition keys

// Validate if topic_from_metadata_key is included in partition_keys

// Validate if message_key_from_metadata_key is included in partition_keys

func isBatchingEnabled(queueBatchConfig configoptional.Optional[exporterhelper.QueueBatchConfig]) bool {
	_ = "STUB: not implemented"
	return false
}

func validateTopicFromMetadataKey(topicFromMetadataKey string, partitionKeysSet map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMessageKeyFromMetadataKey(messageKeyFromMetadataKey string, partitionKeysSet map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}
