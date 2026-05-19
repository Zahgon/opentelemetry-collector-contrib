// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver/internal/metadata"
)

type topicPartition struct {
	topic     string
	partition int32
}

// brokerReadKey is the cache key for OnBrokerRead/OnBrokerWrite metric options.
type brokerReadKey struct {
	nodeID  int32
	outcome string // "success" or "failure"
}

// franzConsumer implements a Kafka consumer using the franz-go client library.
type franzConsumer struct {
	config           *Config
	topics           []string
	excludeTopics    []string
	settings         receiver.Settings
	telemetryBuilder *metadata.TelemetryBuilder
	newConsumeFn     newConsumeMessageFunc
	consumeMessage   consumeMessageFunc

	mu             sync.RWMutex
	started        chan struct{}
	consumerClosed chan struct{}
	closing        chan struct{}

	client      *kgo.Client
	obsrecv     *receiverhelper.ObsReport
	assignments map[topicPartition]*pc

	// brokerReadOpts caches MeasurementOptions for OnBrokerRead, which fires on
	// every fetch request. Entries are evicted in OnBrokerDisconnect; growth is
	// bounded by 2 × number-of-brokers (success + failure).
	brokerReadMu   sync.RWMutex
	brokerReadOpts map[brokerReadKey]metric.MeasurementOption

	// ---- status reporting ----
	host         component.Host
	stoppingOnce sync.Once
	stoppedOnce  sync.Once
}

// pc represents the partition consumer shared information.
type pc struct {
	logger *zap.Logger
	attrs  attribute.Set

	ctx    context.Context
	cancel context.CancelCauseFunc
	// Not safe for concurrent use, this field is never accessed concurrently.
	backOff *backoff.ExponentialBackOff

	mu sync.RWMutex // protects the fields below
	// wg tracks the number of in-flight message processing goroutines for this
	// partition. The wg must not be used directly; instead, the helper methods
	// add() and done() should be called to safely mutate it. These methods ensure
	// that no new goroutines are added once the partition consumer is stopping
	// (i.e. after the partition is lost / revoked).
	wg sync.WaitGroup
}

// add increments the wait group counter if the partition consumer is not
// stopping. It returns true if the counter was incremented, false otherwise.
func (p *pc) add(delta int) bool { _ = "STUB: not implemented"; return false }

// cancelContext cancels the partition consumer context while holding the write
// lock.
func (p *pc) cancelContext(err error) { _ = "STUB: not implemented"; return }

// done decrements the wait group counter.
func (p *pc) done() {
	_ = "STUB: not implemented"

	// wait waits for all in-flight goroutines to finish.
	return
}

func (p *pc) wait() {
	_ = "STUB: not implemented"

	// newFranzKafkaConsumer creates a new franz-go based Kafka consumer
	return
}

func newFranzKafkaConsumer(
	config *Config,
	set receiver.Settings,
	topics []string,
	excludeTopics []string,
	newConsumeFn newConsumeMessageFunc,
) (*franzConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reportStatus emits a component status event if we have a host.
func (c *franzConsumer) reportStatus(s componentstatus.Status) { _ = "STUB: not implemented"; return }

// reportRecoverable reports a recoverable error status event.
func (c *franzConsumer) reportRecoverable(err error) { _ = "STUB: not implemented"; return }

func (c *franzConsumer) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Report "Starting" as soon as Start() is called.

// Create franz-go consumer client

// Create franz-go consumer client

func (c *franzConsumer) consumeLoop(ctx context.Context) {
	_ = "STUB: not implemented"
	// When the loop exits, report Stopped.
	return
}

// Consume messages until the ctx is cancelled (the client is closed).
// NOTE(marclop) we should make the fetch size configurable. It returns
// all the internally buffered records. This isn't something that's
// configurable in Sarama, and theoretically the max records to iterate
// on is a factor of default / max (byte) fetch size.

// consume consumes a batch of messages from the Kafka topic. This is meant to
// be called in a loop until consume returns false.
func (c *franzConsumer) consume(ctx context.Context, size int) bool {
	_ = "STUB: not implemented"
	return false
}

// Shut down the consumer loop.

// There's a variety of errors that are returned by fetch.Errors(). We
// handle the errors that require a client restart above. The rest can
// simply be logged and keep fetching.

// Report recoverable error while consuming.

// Return right away after errors or empty fetch.

// Acquire the read lock on each consume to ensure the client is not closed
// and the assignments map is not modified while consuming. Copy the map
// to avoid locking for the duration of the consume loop.

// Process messages on a per partition basis, wait for them to finish and
// commit the processed records (if autocommit is disabled).

// Skip partitions without any records.

// NOTE(marclop): This could happen if the partition is lost between
// the time the assignments map is copied and the partition is accessed.

// Try to add a new in-flight message processing goroutine to the
// partition consumer. Return immediately if the partition has been
// lost or reassigned.

// Log at DEBUG level for shutdown/rebalance interruptions
// (context cancellation), ERROR for real processing failures.

// Pause consumption for partitions that have fatal errors.
// handleMessage only returns an error when After=true and
// the message should not be marked, so checking !shouldMark
// here is consistent with that contract.

// Stop processing messages.

// Store so we can commit later.

// Handle fatal processing errors. For non-permanent errors
// with backoff enabled, rewind the fetch cursor via SetOffsets
// so the failed record is retried on the next PollRecords call,
// consistent with how a rebalance restarts from the last
// committed offset. No pause/resume is needed because
// PollRecords is blocked on wg.Wait() until this goroutine
// finishes.
// Permanent errors and partitions without backoff configured
// are paused until a rebalance triggers assigned(), which
// calls ResumeFetchPartitions.

// Skip rewind if the consumer is shutting down or the
// partition was lost. In these cases the error is from
// context cancellation, not a real processing failure,
// and calling SetOffsets could interfere with the final
// offset commit.

// No metrics nor marks to update.

// Otherwise, publish consumer lag.

// Wait for all records to be processed and commit if autocommit=false.

// Surface as recoverable error.

func (c *franzConsumer) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Report Stopping at shutdown start.
	return nil
}

// Idempotent: never fail if not started.
// We still want to ensure Stopped is eventually emitted (consumeLoop defer handles it).
// However, if the loop was never started, emit Stopped here too.

// If it returns false, the caller should return immediately.
func (c *franzConsumer) triggerShutdown() bool { _ = "STUB: not implemented"; return false }

// Return immediately if the receiver hasn't started.

// Close the client without holding the write mutex, otherwise, the
// Shutdown will deadlock when `franzConsumer` inevitably calls the
// lost/assigned callback.

// assigned must be set as kgo.OnPartitionsAssigned callback. Ensuring all
// assigned partitions to this consumer process received records.
func (c *franzConsumer) assigned(ctx context.Context, cl *kgo.Client, assigned map[string][]int32) {
	_ = "STUB: not implemented"
	// Report OK on each successful assignment so we can recover status after transient errors.
	return
}

// Resume any partitions that were previously paused due to processing errors.
// PauseFetchPartitions persists across rebalances in franz-go, so we must
// explicitly resume partitions when they are (re)assigned.
// ResumeFetchPartitions is a no-op for partitions that are not paused.

// lost must be set both on kgo.OnPartitionsLost and kgo.OnPartitionsReassigned
// callbacks. Ensures that partitions that are lost (see kgo.OnPartitionsLost
// for more details) or reassigned (see kgo.OnPartitionsReassigned for more
// details) have their partition consumer stopped.
// This callback must finish within the re-balance timeout.
func (c *franzConsumer) lost(ctx context.Context, _ *kgo.Client,
	lost map[string][]int32, fatal bool,
) {
	_ = "STUB: not implemented"
	return
}

// In some cases, it is possible for the `lost` to be called with
// no assignments. So, check if assignments exists first.
//
// - OnPartitionLost can be called without the group ever joining
// and getting assigned.
// - OnPartitionRevoked can be called multiple times for cooperative
// balancer on topic lost/deleted.

// Cancel also locks the partition consumer. This ensures that
// the partition consumer stops processing messages when the
// partition is lost or reassigned.

// Wait for all partition consumers to exit before committing marked offsets.

// NOTE(marclop) commit the marked offsets when the partition is rebalanced
// away from this consumer.

// Report recoverable error on commit errors.

// handleMessage is called on a per-partition basis.
func (c *franzConsumer) handleMessage(pc *pc, record *kgo.Record) error {
	_ = "STUB: not implemented"
	return nil
}

// Successfully processed.

// In the future, with Consumer Share Groups, messages not processed
// within a configurable timeout, are re-delivered to the consumer in
// the Share group, however, at the time of writing this feature isn't
// yet GA nor widely deployed.
// Share groups are generally more in line with observability and OTel
// collector use cases than traditional consumer groups.
// https://cwiki.apache.org/confluence/display/KAFKA/KIP-932%3A+Queues+for+Kafka.
// One possible exception is if the OTel collector is used for analytics
// pipelines, where it may make sense to make share groups opt-in.

// Only return an error if messages are marked after successful processing.

// The methods below implement the relevant franz-go hook interfaces
// record the metrics defined in the metadata telemetry.

// brokerReadOpt returns a cached metric.MeasurementOption for broker read/write hooks.
func (c *franzConsumer) brokerReadOpt(nodeID int32, outcome string) metric.MeasurementOption {
	_ = "STUB: not implemented"
	return *new(metric.MeasurementOption)
}

func (c *franzConsumer) OnBrokerConnect(meta kgo.BrokerMetadata, _ time.Duration, _ net.Conn, err error) {
	_ = "STUB: not implemented"
	return
}

func (c *franzConsumer) OnBrokerDisconnect(meta kgo.BrokerMetadata, _ net.Conn) {
	_ = "STUB: not implemented"
	return
}

// Evict cached read opts for this broker.

func (c *franzConsumer) OnBrokerThrottle(meta kgo.BrokerMetadata, throttleInterval time.Duration, _ bool) {
	_ = "STUB: not implemented"
	return
}

// KafkaBrokerThrottlingDuration is deprecated in favor of KafkaBrokerThrottlingLatency.

func (c *franzConsumer) OnBrokerRead(meta kgo.BrokerMetadata, _ int16, _ int, readWait, timeToRead time.Duration, err error) {
	_ = "STUB: not implemented"
	return
}

// KafkaReceiverLatency is deprecated in favor of KafkaReceiverReadLatency.

// OnFetchBatchRead is called once per batch read from Kafka.
// https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo#HookFetchBatchRead
func (c *franzConsumer) OnFetchBatchRead(meta kgo.BrokerMetadata, topic string, partition int32, m kgo.FetchBatchMetrics) {
	_ = "STUB: not implemented"
	return
}

// KafkaReceiverMessages is deprecated in favor of KafkaReceiverRecords.

// franzConsumerWithOptionalHooks wraps franzConsumer
// so the optional OnFetchRecordUnbuffered can be enabled.
type franzConsumerWithOptionalHooks struct {
	*franzConsumer
}

// OnFetchRecordUnbuffered is called when a fetched record is unbuffered and ready to be processed.
// Note that this hook may slow down high-volume consuming a bit.
// https://pkg.go.dev/github.com/twmb/franz-go/pkg/kgo#HookFetchRecordUnbuffered
func (c franzConsumerWithOptionalHooks) OnFetchRecordUnbuffered(r *kgo.Record, polled bool) {
	_ = "STUB: not implemented"
	return

	// Record metrics when polled by `client.PollRecords()`.
}

func compressionFromCodec(c uint8) string {
	_ = "STUB: not implemented"
	// CompressionType signifies which algorithm the batch was compressed
	// with.
	//
	// 0 is no compression, 1 is gzip, 2 is snappy, 3 is lz4, and 4 is
	// zstd.
	return ""
}

func makeClearLeaderEpochAdjuster() func(context.Context, map[string]map[int32]kgo.Offset) (map[string]map[int32]kgo.Offset, error) {
	_ = "STUB: not implemented"
	return nil
}
