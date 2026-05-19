// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracker // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/tracker"

import (
	"context"
	"os"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/archive"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fileset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fingerprint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

const (
	FileTracker    = "fileTracker"
	NoStateTracker = "noStateTracker"
)

// Interface for tracking files that are being consumed.
type Tracker interface {
	Name() string
	Add(reader *reader.Reader)
	GetCurrentFile(fp *fingerprint.Fingerprint) *reader.Reader
	GetOpenFile(fp *fingerprint.Fingerprint) *reader.Reader
	GetClosedFile(fp *fingerprint.Fingerprint) *reader.Metadata
	GetMetadata() []*reader.Metadata
	LoadMetadata(metadata []*reader.Metadata)
	CurrentPollFiles() []*reader.Reader
	PreviousPollFiles() []*reader.Reader
	ClosePreviousFiles() int
	EndPoll(context.Context)
	EndConsume() int
	TotalReaders() int
	AddUnmatched(*os.File, *fingerprint.Fingerprint)
	LookupArchive(context.Context) ([]*os.File, []*fingerprint.Fingerprint, []*reader.Metadata)
}

// fileTracker tracks known offsets for files that are being consumed by the manager.
type fileTracker struct {
	set component.TelemetrySettings

	maxBatchFiles int

	currentPollFiles  *fileset.Fileset[*reader.Reader]
	previousPollFiles *fileset.Fileset[*reader.Reader]
	knownFiles        []*fileset.Fileset[*reader.Metadata]

	unmatchedFiles []*os.File
	unmatchedFps   []*fingerprint.Fingerprint

	archive archive.Archive
}

func NewFileTracker(ctx context.Context, set component.TelemetrySettings, maxBatchFiles, pollsToArchive int, persister operator.Persister) Tracker {
	_ = "STUB: not implemented"
	return *new(Tracker)
}

func (*fileTracker) Name() string { _ = "STUB: not implemented"; return "" }

func (t *fileTracker) Add(reader *reader.Reader) {
	_ = "STUB: not implemented"
	// add a new reader for tracking
	return
}

func (t *fileTracker) GetCurrentFile(fp *fingerprint.Fingerprint) *reader.Reader {
	_ = "STUB: not implemented"
	return nil
}

func (t *fileTracker) GetOpenFile(fp *fingerprint.Fingerprint) *reader.Reader {
	_ = "STUB: not implemented"
	return nil
}

func (t *fileTracker) GetClosedFile(fp *fingerprint.Fingerprint) *reader.Metadata {
	_ = "STUB: not implemented"
	return nil
}

func (t *fileTracker) AddUnmatched(file *os.File, fp *fingerprint.Fingerprint) {
	_ = "STUB: not implemented"
	// exclude duplicate fingerprints
	return
}

func (t *fileTracker) LookupArchive(ctx context.Context) ([]*os.File, []*fingerprint.Fingerprint, []*reader.Metadata) {
	_ = "STUB: not implemented"
	// LookupArchive performs fingerprint matching against the archive and returns matched metadata, files and fingerprints.
	return nil, nil, nil
}

func (t *fileTracker) GetMetadata() []*reader.Metadata {
	_ = "STUB: not implemented"
	// return all known metadata for checkpoining
	return nil
}

func (t *fileTracker) LoadMetadata(metadata []*reader.Metadata) { _ = "STUB: not implemented"; return }

func (t *fileTracker) CurrentPollFiles() []*reader.Reader { _ = "STUB: not implemented"; return nil }

func (t *fileTracker) PreviousPollFiles() []*reader.Reader { _ = "STUB: not implemented"; return nil }

func (t *fileTracker) ClosePreviousFiles() (filesClosed int) {
	_ = "STUB: not implemented"
	// t.previousPollFiles -> t.knownFiles[0]
	return 0
}

func (t *fileTracker) EndPoll(ctx context.Context) {
	_ = "STUB: not implemented"
	// shift the filesets at end of every poll() call
	// t.knownFiles[0] -> t.knownFiles[1] -> t.knownFiles[2]
	return
}

// Instead of throwing it away, archive it.

func (t *fileTracker) TotalReaders() int { _ = "STUB: not implemented"; return 0 }

// noStateTracker only tracks the current polled files. Once the poll is
// complete and telemetry is consumed, the tracked files are closed. The next
// poll will create fresh readers with no previously tracked offsets.
type noStateTracker struct {
	set              component.TelemetrySettings
	maxBatchFiles    int
	currentPollFiles *fileset.Fileset[*reader.Reader]
	unmatchedFiles   []*os.File
	unmatchedFps     []*fingerprint.Fingerprint
}

func NewNoStateTracker(set component.TelemetrySettings, maxBatchFiles int) Tracker {
	_ = "STUB: not implemented"
	return *new(Tracker)
}

func (*noStateTracker) Name() string { _ = "STUB: not implemented"; return "" }

func (t *noStateTracker) Add(reader *reader.Reader) {
	_ = "STUB: not implemented"
	// add a new reader for tracking
	return
}

func (t *noStateTracker) CurrentPollFiles() []*reader.Reader { _ = "STUB: not implemented"; return nil }

func (t *noStateTracker) GetCurrentFile(fp *fingerprint.Fingerprint) *reader.Reader {
	_ = "STUB: not implemented"
	return nil
}

func (t *noStateTracker) EndConsume() (filesClosed int) { _ = "STUB: not implemented"; return 0 }

func (*noStateTracker) GetOpenFile(*fingerprint.Fingerprint) *reader.Reader {
	_ = "STUB: not implemented"
	return nil
}

func (*noStateTracker) GetClosedFile(*fingerprint.Fingerprint) *reader.Metadata {
	_ = "STUB: not implemented"
	return nil
}

func (*noStateTracker) GetMetadata() []*reader.Metadata { _ = "STUB: not implemented"; return nil }

func (*noStateTracker) LoadMetadata([]*reader.Metadata) { _ = "STUB: not implemented"; return }

func (*noStateTracker) PreviousPollFiles() []*reader.Reader { _ = "STUB: not implemented"; return nil }

func (*noStateTracker) ClosePreviousFiles() int { _ = "STUB: not implemented"; return 0 }

func (*noStateTracker) EndPoll(context.Context) { _ = "STUB: not implemented"; return }

func (*noStateTracker) TotalReaders() int { _ = "STUB: not implemented"; return 0 }

func (t *noStateTracker) AddUnmatched(file *os.File, fp *fingerprint.Fingerprint) {
	_ = "STUB: not implemented"
	return
}

func (t *noStateTracker) LookupArchive(context.Context) ([]*os.File, []*fingerprint.Fingerprint, []*reader.Metadata) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
