// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/kafkaclient"

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
)

var (
	_ kgo.HookBrokerConnect    = (*StatusReporter)(nil)
	_ kgo.HookBrokerDisconnect = (*StatusReporter)(nil)
)

type StatusReporter struct {
	host        component.Host
	connections int
	mu          sync.Mutex
}

func (s *StatusReporter) OnBrokerConnect(_ kgo.BrokerMetadata, _ time.Duration, _ net.Conn, err error) {
	_ = "STUB: not implemented"
	return
}

// only report recoverable errors if none of the brokers are connected

func (s *StatusReporter) OnBrokerDisconnect(_ kgo.BrokerMetadata, _ net.Conn) {
	_ = "STUB: not implemented"
	return
}

func NewStatusReporter(host component.Host) *StatusReporter { _ = "STUB: not implemented"; return nil }

// MessageTooLargeError wraps a MessageTooLarge Kafka error with the actual
// record size that caused the rejection. The size is computed the same way as
// franz-go's Record.userSize: len(Key) + len(Value) + Σ(len(header.Key) + len(header.Value)).
type MessageTooLargeError struct {
	// RecordBytes is the user-visible size of the record (key + value + headers).
	RecordBytes int
	// MaxMessageBytes is the configured producer max message size.
	MaxMessageBytes int
	Err             error
}

func (e *MessageTooLargeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *MessageTooLargeError) Unwrap() error {
	_ = "STUB: not implemented"

	// recordUserSize returns the user-visible size of a kgo.Record, matching
	// franz-go's internal userSize calculation.
	return nil
}

func recordUserSize(r *kgo.Record) int { _ = "STUB: not implemented"; return 0 }

// RecordHeader includes key-value pairs to be added as headers to Kafka records.
type RecordHeader struct {
	Name  string              `mapstructure:"name"`
	Value configopaque.String `mapstructure:"value"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// FranzSyncProducer is a wrapper around the franz-go client that implements
// the Producer interface. Allowing us to use the franz-go client while
// maintaining compatibility with the existing Kafka exporter code.
type FranzSyncProducer struct {
	client          *kgo.Client
	clientCancel    context.CancelFunc
	metadataKeys    []string
	recordHeaders   []kgo.RecordHeader
	maxMessageBytes int
}

// NewFranzSyncProducer Franz-go producer from a kgo.Client and a Messenger.
// clientCancel must cancel the context passed to kgo.WithContext when the client was created;
// it is called by Close to unblock any in-flight ProduceSync calls.
func NewFranzSyncProducer(client *kgo.Client,
	metadataKeys []string,
	recordHeaders []RecordHeader,
	maxMessageBytes int,
	clientCancel context.CancelFunc,
) *FranzSyncProducer {
	_ = "STUB: not implemented"
	return nil
}

// ExportData sends a batch of records to Kafka. It attaches configured
// record headers and per-call metadata-derived headers to each record before
// producing.
func (p *FranzSyncProducer) ExportData(ctx context.Context, records []*kgo.Record) error {
	_ = "STUB: not implemented"
	return nil
}

// check if its defined as a non-retriable error by franzgo

// Close shuts down the producer, unblocking any in-flight ExportData call.
func (p *FranzSyncProducer) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
