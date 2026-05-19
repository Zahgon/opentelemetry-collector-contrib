// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package scanner // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/scanner"

import (
	"bufio"
	"io"
)

const DefaultBufferSize = 16 * 1024

// Scanner is a scanner that maintains position
type Scanner struct {
	pos int64
	*bufio.Scanner
}

// New creates a new positional scanner
func New(r io.Reader, maxLogSize int, buf []byte, startOffset int64, splitFunc bufio.SplitFunc, isGzip bool) *Scanner {
	_ = "STUB: not implemented"
	return nil
}

// flush data if there are no more tokens but there is still data left
// this is because gzip reader reads the entire compressed stream in one go

// Pos returns the current position of the scanner
func (s *Scanner) Pos() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) Error() error { _ = "STUB: not implemented"; return nil }
