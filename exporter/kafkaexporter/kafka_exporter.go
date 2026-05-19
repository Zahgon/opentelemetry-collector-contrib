// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"

import (
	"context"
	"iter"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/kafkaclient"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/marshaler"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/metadata"
)

type messenger[T any] interface {
	// partitionData returns an iterator that yields key-value pairs
	// where the key is the partition key, and the value is the pdata
	// type (plog.Logs, etc.)
	partitionData(T) iter.Seq2[[]byte, T]

	// marshalData marshals a pdata type into zero or more messages,
	// invoking yield once per message with its key and value.
	marshalData(data T, yield func(key, value []byte)) error

	// getTopic returns the topic name for the given context and data.
	getTopic(context.Context, T) string

	// getMessageKey returns the Kafka record key derived from client metadata,
	// or nil if message_key_from_metadata_key is not configured or the metadata
	// value is absent.
	getMessageKey(context.Context) []byte
}

// recordsBuffer is a pooled holder for a batch of kgo.Records. space owns
// the record values; pointers[i] points to space[i] and is what the producer
// API expects. Both slices are reused across exports.
type recordsBuffer struct {
	space    []kgo.Record
	pointers []*kgo.Record
}

type kafkaExporter[T any] struct {
	cfg          Config
	set          exporter.Settings
	tb           *metadata.TelemetryBuilder
	logger       *zap.Logger
	newMessenger func(host component.Host) (messenger[T], error)
	messenger    messenger[T]
	producer     *kafkaclient.FranzSyncProducer
	recordsPool  sync.Pool
}

func newKafkaExporter[T any](
	config Config,
	set exporter.Settings,
	newMessenger func(component.Host) (messenger[T], error),
) *kafkaExporter[T] {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaExporter[T]) Start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaExporter[T]) Close(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaExporter[T]) exportData(ctx context.Context, data T) error {
	_ = "STUB: not implemented"
	return nil
}

// Marshalers may set the key, but a non-nil partition key
// from partitionData takes precedence. The metadata-derived key
// is mutually exclusive with partition_* flags (validated at config
// time), so it applies when partitionData yields nil.

// Build the pointer slice from space. We do this once here rather
// than in lockstep with each append, since append may reallocate
// space's backing array and invalidate earlier pointers.

// TODO move this logging to a kgo hook, so we capture topic and partition details.

func newTracesExporter(config Config, set exporter.Settings) *kafkaExporter[ptrace.Traces] {
	_ = "STUB: not implemented"
	// Jaeger encodings do their own partitioning, so disable trace ID
	// partitioning when they are configured.
	return nil
}

type kafkaTracesMessenger struct {
	config    Config
	marshaler marshaler.TracesMarshaler
}

func (e *kafkaTracesMessenger) marshalData(td ptrace.Traces, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaTracesMessenger) getTopic(ctx context.Context, td ptrace.Traces) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *kafkaTracesMessenger) getMessageKey(ctx context.Context) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaTracesMessenger) partitionData(td ptrace.Traces) iter.Seq2[[]byte, ptrace.Traces] {
	_ = "STUB: not implemented"
	return nil
}

// Note that batchpersignal.SplitTraces guarantees that each trace
// has exactly one trace, and by implication, at least one span.

// NOTE: The same ptrace.Traces instance (newTraces) is reused and mutated on each iteration.
// Callers must treat the yielded pdata as ephemeral and must not retain it beyond
// the current callback/iteration, as its contents will be overwritten on the next yield.

func newLogsExporter(config Config, set exporter.Settings) *kafkaExporter[plog.Logs] {
	_ = "STUB: not implemented"
	return nil
}

type kafkaLogsMessenger struct {
	config    Config
	marshaler marshaler.LogsMarshaler
}

func (e *kafkaLogsMessenger) marshalData(ld plog.Logs, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaLogsMessenger) getTopic(ctx context.Context, ld plog.Logs) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *kafkaLogsMessenger) getMessageKey(ctx context.Context) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaLogsMessenger) partitionData(ld plog.Logs) iter.Seq2[[]byte, plog.Logs] {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: The same plog.Logs instance (newLogs) is reused and mutated on each iteration.
// Callers must treat the yielded pdata as ephemeral and must not retain it beyond
// the current callback/iteration, as its contents will be overwritten on the next yield.

func newMetricsExporter(config Config, set exporter.Settings) *kafkaExporter[pmetric.Metrics] {
	_ = "STUB: not implemented"
	return nil
}

type kafkaMetricsMessenger struct {
	config    Config
	marshaler marshaler.MetricsMarshaler
}

func (e *kafkaMetricsMessenger) marshalData(md pmetric.Metrics, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaMetricsMessenger) getTopic(ctx context.Context, md pmetric.Metrics) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *kafkaMetricsMessenger) getMessageKey(ctx context.Context) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaMetricsMessenger) partitionData(md pmetric.Metrics) iter.Seq2[[]byte, pmetric.Metrics] {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: The same pmetric.Metrics instance (newMetrics) is reused and mutated on each iteration.
// Callers must treat the yielded pdata as ephemeral and must not retain it beyond
// the current callback/iteration, as its contents will be overwritten on the next yield.

func newProfilesExporter(config Config, set exporter.Settings) *kafkaExporter[pprofile.Profiles] {
	_ = "STUB: not implemented"
	return nil
}

type kafkaProfilesMessenger struct {
	config    Config
	marshaler marshaler.ProfilesMarshaler
}

func (e *kafkaProfilesMessenger) marshalData(ld pprofile.Profiles, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaProfilesMessenger) getTopic(ctx context.Context, ld pprofile.Profiles) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *kafkaProfilesMessenger) getMessageKey(ctx context.Context) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (e *kafkaProfilesMessenger) partitionData(pd pprofile.Profiles) iter.Seq2[[]byte, pprofile.Profiles] {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: The same pprofile.Profiles instance (newProfiles) is reused and mutated on each iteration.
// Callers must treat the yielded pdata as ephemeral and must not retain it beyond
// the current callback/iteration, as its contents will be overwritten on the next yield.

type resourceSlice[T any] interface {
	Len() int
	At(int) T
}

type resource interface {
	Resource() pcommon.Resource
}

func getMessageKey(ctx context.Context, signalCfg SignalConfig) []byte {
	_ = "STUB: not implemented"
	return nil
}

func getTopic[T resource](ctx context.Context,
	signalCfg SignalConfig,
	topicFromAttribute string,
	resources resourceSlice[T],
) string {
	_ = "STUB: not implemented"
	return ""
}
