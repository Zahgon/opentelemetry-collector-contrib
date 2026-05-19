// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"

import (
	"bufio"
	"io"
)

// bufferedWriteCloser is intended to use more memory
// in order to optimize writing to disk to help improve performance.
type bufferedWriteCloser struct {
	wrapped  io.Closer
	buffered *bufio.Writer
}

var _ io.WriteCloser = (*bufferedWriteCloser)(nil)

func newBufferedWriteCloser(f io.WriteCloser) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func (bwc *bufferedWriteCloser) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bwc *bufferedWriteCloser) Close() error { _ = "STUB: not implemented"; return nil }

func (bwc *bufferedWriteCloser) flush() error { _ = "STUB: not implemented"; return nil }
