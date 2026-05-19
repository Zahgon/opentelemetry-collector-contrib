// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

type LogEmitter interface {
	operator.Operator
	Start(operator.Persister) error
	Stop() error
	ProcessBatch(context.Context, []*entry.Entry) error
	Process(context.Context, *entry.Entry) error
}

// BatchingLogEmitter is a stanza operator that emits log entries to the consumer callback function `consumerFunc` with batching
type BatchingLogEmitter struct {
	OutputOperator
	cancel        context.CancelFunc
	stopOnce      sync.Once
	batchMux      sync.Mutex
	batch         []*entry.Entry
	wg            sync.WaitGroup
	maxBatchSize  uint
	flushInterval time.Duration
	consumerFunc  func(context.Context, []*entry.Entry)
}

var (
	defaultFlushInterval      = 100 * time.Millisecond
	defaultMaxBatchSize  uint = 100
)

type EmitterOption interface {
	apply(*BatchingLogEmitter)
}

func WithMaxBatchSize(maxBatchSize uint) EmitterOption {
	_ = "STUB: not implemented"
	return *new(EmitterOption)
}

type maxBatchSizeOption struct {
	maxBatchSize uint
}

func (o maxBatchSizeOption) apply(e *BatchingLogEmitter) { _ = "STUB: not implemented"; return }

func WithFlushInterval(flushInterval time.Duration) EmitterOption {
	_ = "STUB: not implemented"
	return *new(EmitterOption)
}

type flushIntervalOption struct {
	flushInterval time.Duration
}

func (o flushIntervalOption) apply(e *BatchingLogEmitter) { _ = "STUB: not implemented"; return }

// NewBatchingLogEmitter creates a new receiver output
func NewBatchingLogEmitter(set component.TelemetrySettings, consumerFunc func(context.Context, []*entry.Entry), opts ...EmitterOption) *BatchingLogEmitter {
	_ = "STUB: not implemented"
	return nil
}

// Start starts the goroutine(s) required for this operator
func (e *BatchingLogEmitter) Start(_ operator.Persister) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop will close the log channel and stop running goroutines
func (e *BatchingLogEmitter) Stop() error { _ = "STUB: not implemented"; return nil }

// the cancel func could be nil if the emitter is never started.

// ProcessBatch emits the entries to the consumerFunc
func (e *BatchingLogEmitter) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// appendEntries appends the entry to the current batch. If maxBatchSize is reached, a new batch will be made, and the old batch
// (which should be flushed) will be returned
func (e *BatchingLogEmitter) appendEntries(entries []*entry.Entry) []*entry.Entry {
	_ = "STUB: not implemented"
	return nil
}

// Process will emit an entry to the consumerFunc
func (e *BatchingLogEmitter) Process(ctx context.Context, ent *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// appendEntry appends the entry to the current batch. If maxBatchSize is reached, a new batch will be made, and the old batch
// (which should be flushed) will be returned
func (e *BatchingLogEmitter) appendEntry(ent *entry.Entry) []*entry.Entry {
	_ = "STUB: not implemented"
	return nil
}

// flusher flushes the current batch every flush interval. Intended to be run as a goroutine
func (e *BatchingLogEmitter) flusher(ctx context.Context) { _ = "STUB: not implemented"; return }

// Create a new context with timeout for the final flush

// flush currently batched entries

// makeNewBatch replaces the current batch on the log emitter with a new batch, returning the old one
func (e *BatchingLogEmitter) makeNewBatch() []*entry.Entry { _ = "STUB: not implemented"; return nil }

// SynchronousLogEmitter is a stanza operator that emits log entries to the consumer callback function `consumerFunc` synchronously
type SynchronousLogEmitter struct {
	OutputOperator
	consumerFunc func(context.Context, []*entry.Entry)
}

func NewSynchronousLogEmitter(set component.TelemetrySettings, consumerFunc func(context.Context, []*entry.Entry)) *SynchronousLogEmitter {
	_ = "STUB: not implemented"
	return nil
}

func (*SynchronousLogEmitter) Start(operator.Persister) error {
	_ = "STUB: not implemented"
	return nil
}

func (*SynchronousLogEmitter) Stop() error { _ = "STUB: not implemented"; return nil }

func (e *SynchronousLogEmitter) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SynchronousLogEmitter) Process(ctx context.Context, ent *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}
