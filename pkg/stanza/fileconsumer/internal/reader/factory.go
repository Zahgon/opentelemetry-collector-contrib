// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package reader // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"

import (
	"bufio"
	"os"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/attrs"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/emit"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fingerprint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/header"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/trim"
)

const (
	DefaultMaxLogSize   = 1024 * 1024
	DefaultFlushPeriod  = 500 * time.Millisecond
	DefaultMaxBatchSize = 100
)

type Factory struct {
	component.TelemetrySettings
	HeaderConfig            *header.Config
	FromBeginning           bool
	FingerprintSize         int
	BufPool                 sync.Pool
	InitialBufferSize       int
	MaxLogSize              int
	TruncateOnMaxLogSize    bool
	Encoding                encoding.Encoding
	SplitFunc               bufio.SplitFunc
	TrimFunc                trim.Func
	FlushTimeout            time.Duration
	EmitFunc                emit.Callback
	Attributes              attrs.Resolver
	DeleteAtEOF             bool
	IncludeFileRecordNumber bool
	IncludeFileRecordOffset bool
	Compression             string
	AcquireFSLock           bool
}

func (f *Factory) NewFingerprint(file *os.File) (*fingerprint.Fingerprint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Factory) NewReader(file *os.File, fp *fingerprint.Fingerprint) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Factory) NewReaderFromMetadata(file *os.File, m *Metadata) (r *Reader, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Re-detect file type when compression is enabled.
// This handles the case where a file was compressed (e.g. test.log → test.log.gz):
// fingerprint matching succeeds because the decompressed content of the .gz matches the original
// plaintext fingerprint, but the file format has changed. Reusing the old FileType and old
// plaintext Offset with a gzip-compressed file causes ReadToEnd to seek to the wrong position
// and read raw compressed bytes as plaintext, producing corrupted log entries.

// Plaintext → gzip compression: the old offset represents the number of
// decompressed bytes already consumed. Decompress the .gz
// from byte 0 and skip that many decompressed bytes so we only emit
// new lines.

// Zero the persisted offset so that if ReadToEnd is skipped (e.g. due to
// context cancellation) and Close() is called immediately, the saved
// metadata carries Offset=0 rather than the stale plaintext value.

// User has reconfigured fingerprint_size

// Merge header attributes with file attributes by creating a new map
// to avoid data race when the original map is accessed concurrently
