// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package archive // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/archive"

import (
	"context"

	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fileset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fingerprint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

const (
	archiveIndexKey = "knownFilesArchiveIndex"
)

type Archive interface {
	FindFiles(context.Context, []*fingerprint.Fingerprint) []*reader.Metadata
	WriteFiles(context.Context, *fileset.Fileset[*reader.Metadata])
}

func New(ctx context.Context, logger *zap.Logger, pollsToArchive int, persister operator.Persister) Archive {
	_ = "STUB: not implemented"
	return *new(Archive)
}

// restore last known archive index

// Try to craft log to explain in user facing terms?

// archiveIndex should point to index for the next write, hence increment it from last known value.

type archive struct {
	persister operator.Persister

	pollsToArchive int

	// archiveIndex points to the index for the next write.
	archiveIndex int
	logger       *zap.Logger
}

func (a *archive) FindFiles(ctx context.Context, fps []*fingerprint.Fingerprint) []*reader.Metadata {
	_ = "STUB: not implemented"
	// FindFiles goes through archive, one fileset at a time and tries to match all fingerprints against that loaded set.
	// To minimize disk access, we first access the index, then review unmatched files and update the metadata, if found.
	// We exit if all fingerprints are matched.
	return nil
}

// Track number of matched fingerprints so we can exit if all matched.

// Determine the index for reading archive, starting from the most recent and moving towards the oldest

// continue executing the loop until either all records are matched or all archive sets have been processed.

// Update the mostRecentIndex

// we load one fileset atmost once per poll

// we've already found a match for this index, continue

// update the matched metada for the index

// we save one fileset atmost once per poll

// Check if all metadata have been found

func (a *archive) WriteFiles(ctx context.Context, metadata *fileset.Fileset[*reader.Metadata]) {
	_ = "STUB: not implemented"
	// We make use of a ring buffer, where each set of files is stored under a specific index.
	// Instead of discarding knownFiles[2], write it to the next index and eventually roll over.
	// Separate storage keys knownFilesArchive0, knownFilesArchive1, ..., knownFilesArchiveN, roll over back to knownFilesArchive0
	return
}

// Archiving:  ┌─────────────────────on-disk archive─────────────────────────┐
//             |    ┌───┐     ┌───┐                     ┌──────────────────┐ |
// index       | ▶  │ 0 │  ▶  │ 1 │  ▶      ...       ▶ │ polls_to_archive │ |
//             | ▲  └───┘     └───┘                     └──────────────────┘ |
//             | ▲    ▲                                                ▼     |
//             | ▲    │ Roll over overriting older offsets, if any     ◀     |
//             └──────│──────────────────────────────────────────────────────┘
//                    │
//                    │
//                    │
//                   start
//                   index

// batch the updated index with metadata

func (a *archive) readArchive(ctx context.Context, index int) (*fileset.Fileset[*reader.Metadata], error) {
	_ = "STUB: not implemented"
	// readArchive loads data from the archive for a given index and returns a fileset.Filset.
	return nil, nil
}

func (a *archive) writeArchive(ctx context.Context, index int, rmds *fileset.Fileset[*reader.Metadata], ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	// writeArchive saves data to the archive for a given index and returns an error, if encountered.
	return nil
}

func getArchiveIndex(ctx context.Context, persister operator.Persister) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func archiveKey(i int) string { _ = "STUB: not implemented"; return "" }

type nopArchive struct{}

func (*nopArchive) FindFiles(_ context.Context, fps []*fingerprint.Fingerprint) []*reader.Metadata {
	_ = "STUB: not implemented"
	// we return an array of "nil"s, indicating 0 matches are found in archive
	return nil
}

func (*nopArchive) WriteFiles(context.Context, *fileset.Fileset[*reader.Metadata]) {
	_ = "STUB: not implemented"
	return
}
