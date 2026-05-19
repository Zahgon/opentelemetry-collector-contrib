// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package groupbytraceprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor/internal/metadata"
)

// groupByTraceProcessor is a processor that keeps traces in memory for a given duration, with the expectation
// that the trace will be complete once this duration expires. After the duration, the trace is sent to the next consumer.
// This processor uses a buffered event machine, which converts operations into events for non-blocking processing, but
// keeping all operations serialized per worker scope. This ensures that we don't need locks but that the state is consistent across go routines.
// Initially, all incoming batches are split into different traces and distributed among workers by a hash of traceID in eventMachine.consume method.
// Afterwards, the trace is registered with a go routine, which will be called after the given duration and dispatched to the event
// machine for further processing.
// The typical data flow looks like this:
// ConsumeTraces -> eventMachine.consume(trace) -> event(traceReceived) -> onTraceReceived -> AfterFunc(duration, event(traceExpired)) -> onTraceExpired
// async markAsReleased -> event(traceReleased) -> onTraceReleased -> nextConsumer
// Each worker in the eventMachine also uses a ring buffer to hold the in-flight trace IDs, so that we don't hold more than the given maximum number
// of traces in memory/storage. Items that are evicted from the buffer are discarded without warning.
type groupByTraceProcessor struct {
	nextConsumer     consumer.Traces
	config           Config
	logger           *zap.Logger
	telemetryBuilder *metadata.TelemetryBuilder
	// the event machine handling all operations for this processor
	eventMachine *eventMachine

	// the trace storage
	st storage
}

var _ processor.Traces = (*groupByTraceProcessor)(nil)

const bufferSize = 10_000

// newGroupByTraceProcessor returns a new processor.
func newGroupByTraceProcessor(set processor.Settings, nextConsumer consumer.Traces, config Config) *groupByTraceProcessor {
	_ = "STUB: not implemented"
	return nil
}

// the event machine will buffer up to N concurrent events before blocking

// register the callbacks

func (sp *groupByTraceProcessor) ConsumeTraces(_ context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (*groupByTraceProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// Start is invoked during service startup.
func (sp *groupByTraceProcessor) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	// start these metrics, as it might take a while for them to receive their first event
	return nil
}

// Shutdown is invoked during service shutdown.
func (sp *groupByTraceProcessor) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (sp *groupByTraceProcessor) onTraceReceived(trace tracesWithID, worker *eventMachineWorker) error {
	_ = "STUB: not implemented"
	return nil
}

// it exists in memory already, just append the spans to the trace in the storage

// we are done with this trace, move on

// at this point, we determined that we haven't seen the trace yet, so, record the
// traceID in the map and the spans to the storage

// place the trace ID in the buffer, and check if an item had to be evicted

// delete from the storage

// we have the traceID in the memory, place the spans in the storage too

// if the event machine has stopped, it will just discard the event

func (sp *groupByTraceProcessor) onTraceExpired(traceID pcommon.TraceID, worker *eventMachineWorker) error {
	_ = "STUB: not implemented"
	return nil
}

// we likely received multiple batches with spans for the same trace
// and released this trace already

// delete from the map and erase its memory entry

// this might block, but we don't need to wait

func (sp *groupByTraceProcessor) markAsReleased(traceID pcommon.TraceID, fire func(...event)) error {
	_ = "STUB: not implemented"
	// #get is a potentially blocking operation
	return nil
}

// signal that the trace is ready to be released

// atomically fire the two events, so that a concurrent shutdown won't leave
// an orphaned trace in the storage

func (sp *groupByTraceProcessor) onTraceReleased(rss []ptrace.ResourceSpans) error {
	_ = "STUB: not implemented"
	return nil
}

// Do async consuming not to block event worker

func (sp *groupByTraceProcessor) onTraceRemoved(traceID pcommon.TraceID) error {
	_ = "STUB: not implemented"
	return nil
}

func (sp *groupByTraceProcessor) addSpans(traceID pcommon.TraceID, trace ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}
