// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"context"
	"errors"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	errNonPermanent = errors.New("non permanent error")
	errPermanent    = errors.New("permanent error")
)

type decisionFunc func() error

// MockBackend is a backend that allows receiving the data locally.
type MockBackend struct {
	// Metric and trace consumers
	tc *MockTraceConsumer
	mc *MockMetricConsumer
	lc *MockLogConsumer

	receiver DataReceiver

	// Log file
	logFilePath string
	logFile     *os.File

	// Start/stop flags
	isStarted  bool
	stopOnce   sync.Once
	startedAt  time.Time
	startMutex sync.Mutex

	// Recording fields.
	isRecording     bool
	recordMutex     sync.Mutex
	ReceivedTraces  []ptrace.Traces
	ReceivedMetrics []pmetric.Metrics
	ReceivedLogs    []plog.Logs

	DroppedTraces  []ptrace.Traces
	DroppedMetrics []pmetric.Metrics
	DroppedLogs    []plog.Logs

	LogsToRetry []plog.Logs

	// decision to return permanent/non-permanent errors
	decision decisionFunc
}

// NewMockBackend creates a new mock backend that receives data using specified receiver.
func NewMockBackend(logFilePath string, receiver DataReceiver) *MockBackend {
	_ = "STUB: not implemented"
	return nil
}

func (mb *MockBackend) WithDecisionFunc(decision decisionFunc) { _ = "STUB: not implemented"; return }

// Start a backend.
func (mb *MockBackend) Start() error { _ = "STUB: not implemented"; return nil }

// Open log file

// Stop the backend
func (mb *MockBackend) Stop() { _ = "STUB: not implemented"; return }

// Print stats.

// EnableRecording enables recording of all data received by MockBackend.
func (mb *MockBackend) EnableRecording() { _ = "STUB: not implemented"; return }

func (mb *MockBackend) GetStats() string { _ = "STUB: not implemented"; return "" }

// DataItemsReceived returns total number of received spans and metrics.
func (mb *MockBackend) DataItemsReceived() uint64 { _ = "STUB: not implemented"; return 0 }

// ClearReceivedItems clears the list of received traces and metrics. Note: counters
// return by DataItemsReceived() are not cleared, they are cumulative.
func (mb *MockBackend) ClearReceivedItems() { _ = "STUB: not implemented"; return }

func (mb *MockBackend) GetReceivedLogs() []plog.Logs { _ = "STUB: not implemented"; return nil }

func (mb *MockBackend) GetReceivedTraces() []ptrace.Traces { _ = "STUB: not implemented"; return nil }

func (mb *MockBackend) ConsumeTrace(td ptrace.Traces) { _ = "STUB: not implemented"; return }

func (mb *MockBackend) ConsumeMetric(md pmetric.Metrics) { _ = "STUB: not implemented"; return }

var _ consumer.Traces = (*MockTraceConsumer)(nil)

func (mb *MockBackend) ConsumeLogs(ld plog.Logs) { _ = "STUB: not implemented"; return }

type MockTraceConsumer struct {
	numSpansReceived atomic.Uint64
	backend          *MockBackend
}

func (*MockTraceConsumer) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (tc *MockTraceConsumer) ConsumeTraces(_ context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore the seqnums for now. We will use them later.

var _ consumer.Metrics = (*MockMetricConsumer)(nil)

type MockMetricConsumer struct {
	numMetricsReceived atomic.Uint64
	backend            *MockBackend
}

func (*MockMetricConsumer) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (mc *MockMetricConsumer) ConsumeMetrics(_ context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (tc *MockTraceConsumer) MockConsumeTraceData(spanCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (mc *MockMetricConsumer) MockConsumeMetricData(metricsCount int) error {
	_ = "STUB: not implemented"
	return nil
}

type MockLogConsumer struct {
	numLogRecordsReceived atomic.Uint64
	backend               *MockBackend
}

func (*MockLogConsumer) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (lc *MockLogConsumer) ConsumeLogs(_ context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// randomNonPermanentError is a decision function that succeeds approximately
// half of the time and fails with a non-permanent error the rest of the time.
func RandomNonPermanentError() error { _ = "STUB: not implemented"; return nil }

func GenerateNonPernamentErrorUntil(ch chan bool) error { _ = "STUB: not implemented"; return nil }

// randomPermanentError is a decision function that succeeds approximately
// half of the time and fails with a permanent error the rest of the time.
func RandomPermanentError() error { _ = "STUB: not implemented"; return nil }
