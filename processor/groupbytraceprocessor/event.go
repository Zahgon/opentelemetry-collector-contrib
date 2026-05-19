// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package groupbytraceprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"

import (
	"errors"
	"hash/maphash"
	"sync"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor/internal/metadata"
)

const (
	// traces received from the previous processors
	traceReceived eventType = iota

	// traceID to be released
	traceExpired

	// released traces
	traceReleased

	// traceID to be removed
	traceRemoved
)

var (
	errNoTraceID = errors.New("trace doesn't have traceID")

	seed = maphash.MakeSeed()

	hashPool = sync.Pool{
		New: func() any {
			var hash maphash.Hash
			hash.SetSeed(seed)
			return &hash
		},
	}
)

type (
	eventType int
	event     struct {
		typ     eventType
		payload any
	}
)

type tracesWithID struct {
	id pcommon.TraceID
	td ptrace.Traces
}

// eventMachine is a machine that accepts events in a typically non-blocking manner,
// processing the events serially per worker scope, to ensure that data at the consumer is consistent.
// Just like the machine itself is non-blocking, consumers are expected to also not block
// on the callbacks, otherwise, events might pile up. When enough events are piled up, firing an
// event will block until enough capacity is available to accept the events.
type eventMachine struct {
	workers                   []*eventMachineWorker
	close                     chan struct{}
	metricsCollectionInterval time.Duration
	shutdownTimeout           time.Duration

	logger          *zap.Logger
	telemetry       *metadata.TelemetryBuilder
	onTraceReceived func(td tracesWithID, worker *eventMachineWorker) error
	onTraceExpired  func(traceID pcommon.TraceID, worker *eventMachineWorker) error
	onTraceReleased func(rss []ptrace.ResourceSpans) error
	onTraceRemoved  func(traceID pcommon.TraceID) error

	onError func(event)

	// shutdown sync
	shutdownLock *sync.RWMutex
	closed       bool
}

func newEventMachine(logger *zap.Logger, bufferSize, numWorkers, numTraces int, telemetry *metadata.TelemetryBuilder) *eventMachine {
	_ = "STUB: not implemented"
	return nil
}

func (em *eventMachine) startInBackground() { _ = "STUB: not implemented"; return }

func (em *eventMachine) numEvents() int { _ = "STUB: not implemented"; return 0 }

func (em *eventMachine) periodicMetrics() { _ = "STUB: not implemented"; return }

func (em *eventMachine) startWorkers() { _ = "STUB: not implemented"; return }

func (em *eventMachine) handleEvent(e event, w *eventMachineWorker) {
	_ = "STUB: not implemented"
	return
}

// the payload had an unexpected type!

// the payload had an unexpected type!

// the payload had an unexpected type!

// the payload had an unexpected type!

// consume takes a single trace and routes it to one of the workers.
func (em *eventMachine) consume(td ptrace.Traces) error { _ = "STUB: not implemented"; return nil }

func workerIndexForTraceID(traceID pcommon.TraceID, numWorkers int) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (em *eventMachine) shutdown() { _ = "STUB: not implemented"; return }

// we never return an error here

// Check immediately first

func (em *eventMachine) callOnError(e event) { _ = "STUB: not implemented"; return }

// handleEventWithObservability uses the given function to process and event,
// recording the event's latency and timing out if it doesn't finish within a reasonable duration
func (em *eventMachine) handleEventWithObservability(event string, do func() error) {
	_ = "STUB: not implemented"
	return
}

type eventMachineWorker struct {
	machine *eventMachine

	// the ring buffer holds the IDs for all the in-flight traces
	buffer *ringBuffer

	events chan event
}

func (w *eventMachineWorker) start() {
	_ = "STUB: not implemented"

	// Prioritize shutdown: check if we should stop before processing next event
	return
}

// Process events or handle shutdown

func (w *eventMachineWorker) fire(events ...event) { _ = "STUB: not implemented"; return }

// we are not accepting new events

// doWithTimeout wraps a function in a timeout, returning whether it succeeded before timing out.
// If the function returns an error within the timeout, it's considered as succeeded and the error will be returned back to the caller.
func doWithTimeout(timeout time.Duration, do func() error) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getTraceID(td ptrace.Traces) (pcommon.TraceID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID), nil
}
