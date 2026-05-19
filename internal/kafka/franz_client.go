// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafka // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/kafka"

import (
	"context"
	"time"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/kafka/configkafka"
)

const (
	SCRAMSHA512          = "SCRAM-SHA-512"
	SCRAMSHA256          = "SCRAM-SHA-256"
	PLAIN                = "PLAIN"
	AWSMSKIAMOAUTHBEARER = "AWS_MSK_IAM_OAUTHBEARER" //nolint:gosec // These aren't credentials.
	OAUTHBEARER          = "OAUTHBEARER"
)

type contextTokenSource interface {
	Token(context.Context) (*oauth2.Token, error)
}

// NewFranzSyncProducer creates a new Kafka client using the franz-go library.
func NewFranzSyncProducer(
	ctx context.Context,
	host component.Host,
	clientCfg configkafka.ClientConfig,
	cfg configkafka.ProducerConfig,
	timeout time.Duration,
	logger *zap.Logger,
	opts ...kgo.Opt,
) (*kgo.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prepend a default sarama-compatible partitioner so that callers can override
// it by appending their own kgo.RecordPartitioner option in opts.

// Configure required acks

// NOTE(marclop) only disable if acks != all.

// WaitForLocal
// NOTE(marclop) only disable if acks != all.

// Configure auto topic creation

// NewFranzConsumerGroup creates a new Kafka consumer client using the franz-go library.
func NewFranzConsumerGroup(
	ctx context.Context,
	host component.Host,
	clientCfg configkafka.ClientConfig,
	consumerCfg configkafka.ConsumerConfig,
	topics []string,
	excludeTopics []string,
	logger *zap.Logger,
	opts ...kgo.Opt,
) (*kgo.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if any topic uses regex pattern

// Similar to librdkafka, if the topic starts with `^`, it is a regex topic:
// https://github.com/confluentinc/librdkafka/blob/b871fdabab84b2ea1be3866a2ded4def7e31b006/src/rdkafka.h#L3899-L3938

// Add exclude topics only when regex consumption is enabled

// Set auto-commit interval to a very high value to "disable" it, but
// still allow using marks.

// Configure auto-commit to use marks, this simplifies the committing
// logic and makes it more consistent with the Sarama client.

// Configure the offset to reset to if an exception is found (or no current
// partition offset is found.

// Configure group instance ID if provided

// Configure rebalance strategy

// NewFranzClient creates a franz-go client using the same commonOpts used for producer/consumer.
func NewFranzClient(
	ctx context.Context,
	host component.Host,
	clientCfg configkafka.ClientConfig,
	logger *zap.Logger,
	opts ...kgo.Opt,
) (*kgo.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewFranzClusterAdminClient creates a kadm admin client from a freshly created franz client.
func NewFranzClusterAdminClient(
	ctx context.Context,
	host component.Host,
	clientCfg configkafka.ClientConfig,
	logger *zap.Logger,
	opts ...kgo.Opt,
) (*kadm.Client, *kgo.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// balancerOptFromStrategy returns a kgo.Opt that sets the group balancer, or
// nil if strategy is empty (franz-go default applies). Built-in strategy names
// map to the corresponding kgo balancer. Any other value is treated as a
// component ID referencing an extension that implements kgo.GroupBalancer.
func balancerOptFromStrategy(strategy configkafka.GroupRebalanceStrategy, host component.Host) (kgo.Opt, error) {
	_ = "STUB: not implemented"
	return *new(kgo.Opt), nil
}

func commonOpts(
	ctx context.Context,
	host component.Host,
	clientCfg configkafka.ClientConfig,
	logger *zap.Logger,
	opts ...kgo.Opt,
) ([]kgo.Opt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disable client metrics, since some brokers may falsely indicate
// that they support them when they don't, causing errors to be
// logged. We may want to make this configurable in the future.

// Configure TLS if needed

// Configure authentication

// Configure client ID

// Configure client rack if provided

// Reuse existing metadata refresh interval for franz-go metadataMaxAge

// Configure connection idle timeout

// Configure the min/max protocol version if provided

func configureKgoSASL(cfg *configkafka.SASLConfig, host component.Host) (kgo.Opt, error) {
	_ = "STUB: not implemented"
	return *new(kgo.Opt), nil
}

func configureKgoKerberos(cfg *configkafka.KerberosConfig) (kgo.Opt, error) {
	_ = "STUB: not implemented"
	return *new(kgo.Opt), nil
}

func compressionCodec(compression string) kgo.CompressionCodec {
	_ = "STUB: not implemented"
	return *new(kgo.CompressionCodec)
}

func newSaramaCompatPartitioner() kgo.Partitioner {
	_ = "STUB: not implemented"
	return *new(kgo.Partitioner)
}

// NewSaramaCompatHasher returns a PartitionerHasher that replicates the default
// Sarama partitioning behavior: FNV-1a hashing with Sarama's int32 sign convention.
func NewSaramaCompatHasher() kgo.PartitionerHasher {
	_ = "STUB: not implemented"
	return *new(kgo.PartitionerHasher)
}

func saramaHashFn(b []byte) uint32 { _ = "STUB: not implemented"; return 0 }
