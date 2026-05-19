// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/prometheus/prometheus/prompb"
	"github.com/tidwall/wal"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter/internal/metadata"
)

type prwWalTelemetry interface {
	recordWALWriteLatency(ctx context.Context, durationMs int64)
	recordWALWrites(ctx context.Context)
	recordWALWritesFailures(ctx context.Context)
	recordWALReadLatency(ctx context.Context, durationMs int64)
	recordWALReads(ctx context.Context)
	recordWALReadsFailures(ctx context.Context)
	recordWALBytesWritten(ctx context.Context, bytes int)
	recordWALBytesRead(ctx context.Context, bytes int)
	recordWALLag(ctx context.Context, lag int64)
}

type prwWalTelemetryOTel struct {
	telemetryBuilder *metadata.TelemetryBuilder
	otelAttrs        []attribute.KeyValue
}

func (p *prwWalTelemetryOTel) recordWALWriteLatency(ctx context.Context, durationMs int64) {
	_ = "STUB: not implemented"
	return
}

func (p *prwWalTelemetryOTel) recordWALWrites(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (p *prwWalTelemetryOTel) recordWALWritesFailures(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (p *prwWalTelemetryOTel) recordWALReadLatency(ctx context.Context, durationMs int64) {
	_ = "STUB: not implemented"
	return
}

func (p *prwWalTelemetryOTel) recordWALReads(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (p *prwWalTelemetryOTel) recordWALReadsFailures(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (p *prwWalTelemetryOTel) recordWALBytesWritten(ctx context.Context, bytes int) {
	_ = "STUB: not implemented"
	return
}

func (p *prwWalTelemetryOTel) recordWALBytesRead(ctx context.Context, bytes int) {
	_ = "STUB: not implemented"
	return
}

func (p *prwWalTelemetryOTel) recordWALLag(ctx context.Context, lag int64) {
	_ = "STUB: not implemented"
	return
}

func newPRWWalTelemetry(set exporter.Settings) (prwWalTelemetry, error) {
	_ = "STUB: not implemented"
	return *new(prwWalTelemetry), nil
}

type prweWAL struct {
	wg        sync.WaitGroup // wg waits for the go routines to finish.
	mu        sync.Mutex     // mu protects the fields below.
	wal       *wal.Log
	walConfig *WALConfig
	walPath   string

	exportSink func(ctx context.Context, reqL []*prompb.WriteRequest) error

	stopOnce  sync.Once
	stopChan  chan struct{}
	rNotify   chan struct{}
	rWALIndex *atomic.Uint64
	wWALIndex *atomic.Uint64

	telemetry prwWalTelemetry
}

const (
	defaultWALBufferSize         = 300
	defaultWALTruncateFrequency  = 1 * time.Minute
	defaultWALLagRecordFrequency = 15 * time.Second
)

type WALConfig struct {
	Directory          string        `mapstructure:"directory"`
	BufferSize         int           `mapstructure:"buffer_size"`
	TruncateFrequency  time.Duration `mapstructure:"truncate_frequency"`
	LagRecordFrequency time.Duration `mapstructure:"lag_record_frequency"`
}

func (wc *WALConfig) bufferSize() int { _ = "STUB: not implemented"; return 0 }

func (wc *WALConfig) truncateFrequency() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (wc *WALConfig) lagRecordInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func newWAL(walConfig *WALConfig, set exporter.Settings, exportSink func(context.Context, []*prompb.WriteRequest) error) (*prweWAL, error) {
	_ = "STUB: not implemented"
	return nil,

		// There are cases for which the WAL can be disabled.
		// TODO: Perhaps log that the WAL wasn't enabled.
		nil
}

// Buffered to avoid lost wake-ups when the writer signals before the
// reader starts waiting on notifications.

func (wc *WALConfig) createWAL() (*wal.Log, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

var (
	errAlreadyClosed = errors.New("already closed")
	errNilWAL        = errors.New("wal is nil")
)

// retrieveWALIndices queries the WriteAheadLog for its current first and last indices.
func (prweWAL *prweWAL) retrieveWALIndices() (err error) { _ = "STUB: not implemented"; return nil }

func (prweWAL *prweWAL) stop() error { _ = "STUB: not implemented"; return nil }

// run begins reading from the WAL until prwe.stopChan is closed.
func (prweWAL *prweWAL) run(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

// Start the process of exporting but wait until the exporting has started.

// log err

// Restart WAL

func (prweWAL *prweWAL) recordLagLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

// In normal state, wIndex and rIndex will differ by one. To avoid having -1 as a final value, we set it to 0 as minimum.

// continuallyPopWALThenExport reads a prompb.WriteRequest proto encoded blob from the WAL, and moves
// the WAL's front index forward until either the read buffer period expires or the maximum
// buffer size is exceeded. When either of the two conditions are matched, it then exports
// the requests to the Remote-Write endpoint, and then truncates the head of the WAL to where
// it last read from.
func (prweWAL *prweWAL) continuallyPopWALThenExport(ctx context.Context, signalStart func()) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Keeping it within a closure to ensure that the later
// updated value of reqL is always flushed to disk.

// Added in a closure to ensure we capture the later
// updated value of timer when changed in the loop below.

// Otherwise, it is time to export, flush and then truncate the WAL, but also to kill the timer!

// Reset but reuse the write requests slice.

func (prweWAL *prweWAL) closeWAL() error { _ = "STUB: not implemented"; return nil }

func (prweWAL *prweWAL) syncAndTruncateFront() error { _ = "STUB: not implemented"; return nil }

// Save all the entries that aren't yet committed, to the tail of the WAL.

// Truncate the WAL from the front for the entries that we already
// read from the WAL and had already exported.

func (prweWAL *prweWAL) exportThenFrontTruncateWAL(ctx context.Context, reqL []*prompb.WriteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset by retrieving the respective read and write WAL indices.

// persistToWAL is the routine that'll be hooked into the exporter's receiving side and it'll
// write them to the Write-Ahead-Log so that shutdowns won't lose data, and that the routine that
// reads from the WAL can then process the previously serialized requests.
func (prweWAL *prweWAL) persistToWAL(ctx context.Context, requests []*prompb.WriteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Write all the requests to the WAL in a batch.

// Notify reader go routine that is possibly waiting for writes.

func (prweWAL *prweWAL) readPrompbFromWAL(ctx context.Context, index uint64) (wreq *prompb.WriteRequest, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Firstly check if we've been terminated, then exit if so.

// The read succeeded.

// Now increment the WAL's read index.

// If WAL was empty, let's wait for a notification from
// the writer go routine.

// record all failures apart ErrNotFound
