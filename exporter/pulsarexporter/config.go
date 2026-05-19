// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"

import (
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for Pulsar exporter.
type Config struct {
	TimeoutSettings           exporterhelper.TimeoutConfig                             `mapstructure:",squash"`
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`

	// Endpoint of pulsar broker (default "pulsar://localhost:6650")
	Endpoint string `mapstructure:"endpoint"`
	// The name of the pulsar topic to export to (default otlp_spans for traces, otlp_metrics for metrics)
	Topic string `mapstructure:"topic"`
	// Encoding of messages (default "otlp_proto")
	Encoding string `mapstructure:"encoding"`
	// Producer configuration of the Pulsar producer
	Producer Producer `mapstructure:"producer"`
	// Set the path to the trusted TLS certificate file
	TLSTrustCertsFilePath string `mapstructure:"tls_trust_certs_file_path"`
	// Configure whether the Pulsar client accept untrusted TLS certificate from broker (default: false)
	TLSAllowInsecureConnection bool           `mapstructure:"tls_allow_insecure_connection"`
	Authentication             Authentication `mapstructure:"auth"`
	OperationTimeout           time.Duration  `mapstructure:"operation_timeout"`
	ConnectionTimeout          time.Duration  `mapstructure:"connection_timeout"`
	MaxConnectionsPerBroker    int            `mapstructure:"max_connections_per_broker"`
}

type Authentication struct {
	TLS    configoptional.Optional[TLS]    `mapstructure:"tls"`
	Token  configoptional.Optional[Token]  `mapstructure:"token"`
	Athenz configoptional.Optional[Athenz] `mapstructure:"athenz"`
	OAuth2 configoptional.Optional[OAuth2] `mapstructure:"oauth2"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type TLS struct {
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type Token struct {
	Token configopaque.String `mapstructure:"token"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type Athenz struct {
	ProviderDomain  string              `mapstructure:"provider_domain"`
	TenantDomain    string              `mapstructure:"tenant_domain"`
	TenantService   string              `mapstructure:"tenant_service"`
	PrivateKey      configopaque.String `mapstructure:"private_key"`
	KeyID           string              `mapstructure:"key_id"`
	PrincipalHeader string              `mapstructure:"principal_header"`
	ZtsURL          string              `mapstructure:"zts_url"`
}

type OAuth2 struct {
	IssuerURL  string `mapstructure:"issuer_url"`
	ClientID   string `mapstructure:"client_id"`
	Audience   string `mapstructure:"audience"`
	PrivateKey string `mapstructure:"private_key"`
	Scope      string `mapstructure:"scope"`
}

// Producer defines configuration for producer
type Producer struct {
	MaxReconnectToBroker            *uint            `mapstructure:"max_reconnect_broker"`
	HashingScheme                   HashingScheme    `mapstructure:"hashing_scheme"`
	CompressionLevel                CompressionLevel `mapstructure:"compression_level"`
	CompressionType                 CompressionType  `mapstructure:"compression_type"`
	MaxPendingMessages              int              `mapstructure:"max_pending_messages"`
	BatcherBuilderType              BatchBuilderType `mapstructure:"batch_builder_type"`
	PartitionsAutoDiscoveryInterval time.Duration    `mapstructure:"partitions_auto_discovery_interval"`
	BatchingMaxPublishDelay         time.Duration    `mapstructure:"batching_max_publish_delay"`
	BatchingMaxMessages             uint             `mapstructure:"batching_max_messages"`
	BatchingMaxSize                 uint             `mapstructure:"batching_max_size"`
	DisableBlockIfQueueFull         bool             `mapstructure:"disable_block_if_queue_full"`
	DisableBatching                 bool             `mapstructure:"disable_batching"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (*Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) auth() pulsar.Authentication {
	_ = "STUB: not implemented"
	return *new(pulsar.Authentication)
}

func (cfg *Config) clientOptions() pulsar.ClientOptions {
	_ = "STUB: not implemented"
	return *new(pulsar.ClientOptions)
}

func (cfg *Config) getProducerOptions() pulsar.ProducerOptions {
	_ = "STUB: not implemented"
	return *new(pulsar.ProducerOptions)
}

type BatchBuilderType string

const (
	DefaultBatchBuilder  BatchBuilderType = "default"
	KeyBasedBatchBuilder BatchBuilderType = "key_based"
)

func (c *BatchBuilderType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c *BatchBuilderType) ToPulsar() pulsar.BatcherBuilderType {
	_ = "STUB: not implemented"
	return *new(pulsar.BatcherBuilderType)
}

type CompressionType string

const (
	None CompressionType = "none"
	LZ4  CompressionType = "lz4"
	ZLib CompressionType = "zlib"
	ZStd CompressionType = "zstd"
)

func (c *CompressionType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c *CompressionType) ToPulsar() pulsar.CompressionType {
	_ = "STUB: not implemented"
	return *new(pulsar.CompressionType)
}

type CompressionLevel string

const (
	Default CompressionLevel = "default"
	Faster  CompressionLevel = "faster"
	Better  CompressionLevel = "better"
)

func (c *CompressionLevel) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c *CompressionLevel) ToPulsar() pulsar.CompressionLevel {
	_ = "STUB: not implemented"
	return *new(pulsar.CompressionLevel)
}

type HashingScheme string

const (
	JavaStringHash HashingScheme = "java_string_hash"
	Murmur3_32Hash HashingScheme = "murmur3_32hash"
)

func (c *HashingScheme) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c *HashingScheme) ToPulsar() pulsar.HashingScheme {
	_ = "STUB: not implemented"
	return *new(pulsar.HashingScheme)
}
