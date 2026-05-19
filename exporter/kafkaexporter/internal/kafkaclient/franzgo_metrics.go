// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/kafkaclient"

import (
	"net"
	"sync"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/otel/metric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/metadata"
)

// brokerKey is the cache key for per-broker attribute sets used in OnBrokerE2E.
// The set of distinct (nodeID, host, outcome) combinations is bounded by
// 2 × number-of-brokers (success + failure), so the cache never grows unboundedly.
type brokerKey struct {
	nodeID  int32
	host    string
	outcome string // "success" | "failure"
}

// FranzProducerMetrics implements the relevant franz-go hook interfaces to
// record the metrics defined in the metadata telemetry.
type FranzProducerMetrics struct {
	tb            *metadata.TelemetryBuilder
	brokerE2EMu   sync.RWMutex
	brokerE2EOpts map[brokerKey]metric.MeasurementOption
}

// NewFranzProducerMetrics creates an instance of FranzProducerMetrics from metadata TelemetryBuilder.
func NewFranzProducerMetrics(tb *metadata.TelemetryBuilder) *FranzProducerMetrics {
	_ = "STUB: not implemented"
	return nil
}

// brokerE2EOpt returns a cached MeasurementOption for the given key, building
// and storing it on first use.
func (fpm *FranzProducerMetrics) brokerE2EOpt(k brokerKey) metric.MeasurementOption {
	_ = "STUB: not implemented"
	return *new(metric.MeasurementOption)
}

var _ kgo.HookBrokerConnect = (*FranzProducerMetrics)(nil)

func (fpm *FranzProducerMetrics) OnBrokerConnect(meta kgo.BrokerMetadata, _ time.Duration, _ net.Conn, err error) {
	_ = "STUB: not implemented"
	return
}

var _ kgo.HookBrokerDisconnect = (*FranzProducerMetrics)(nil)

func (fpm *FranzProducerMetrics) OnBrokerDisconnect(meta kgo.BrokerMetadata, _ net.Conn) {
	_ = "STUB: not implemented"
	return
}

// Evict cached attribute sets for this broker so that nodes which come
// and go over time do not cause unbounded growth of brokerE2EOpts.

var _ kgo.HookBrokerThrottle = (*FranzProducerMetrics)(nil)

func (fpm *FranzProducerMetrics) OnBrokerThrottle(meta kgo.BrokerMetadata, throttleInterval time.Duration, _ bool) {
	_ = "STUB: not implemented"
	return
}

// KafkaBrokerThrottlingDuration is deprecated in favor of KafkaBrokerThrottlingLatency.

var _ kgo.HookBrokerE2E = (*FranzProducerMetrics)(nil)

func (fpm *FranzProducerMetrics) OnBrokerE2E(meta kgo.BrokerMetadata, key int16, e2e kgo.BrokerE2E) {
	_ = "STUB: not implemented"
	// Do not pollute producer metrics with non-produce requests
	return
}

// KafkaExporterLatency is deprecated in favor of KafkaExporterWriteLatency.

var _ kgo.HookProduceBatchWritten = (*FranzProducerMetrics)(nil)

// OnProduceBatchWritten is called when a batch has been produced.
// https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo#HookProduceBatchWritten
func (fpm *FranzProducerMetrics) OnProduceBatchWritten(meta kgo.BrokerMetadata, topic string, partition int32, m kgo.ProduceBatchMetrics) {
	_ = "STUB: not implemented"
	return
}

// KafkaExporterMessages is deprecated in favor of KafkaExporterRecords.

var _ kgo.HookProduceRecordUnbuffered = (*FranzProducerMetrics)(nil)

// OnProduceRecordUnbuffered records the number of produced messages that were
// not produced due to errors. The successfully produced records is recorded by
// `OnProduceBatchWritten`.
// https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo#HookProduceRecordUnbuffered
func (fpm *FranzProducerMetrics) OnProduceRecordUnbuffered(r *kgo.Record, err error) {
	_ = "STUB: not implemented"
	return

	// Covered by OnProduceBatchWritten.
}

// KafkaExporterMessages is deprecated in favor of KafkaExporterRecords.

func compressionFromCodec(c uint8) string {
	_ = "STUB: not implemented"
	// CompressionType signifies which algorithm the batch was compressed
	// with.
	//
	// 0 is no compression, 1 is gzip, 2 is snappy, 3 is lz4, and 4 is
	// zstd.
	return ""
}
