// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"sync"
)

var errOverCapacity = errors.New("over capacity")

type buffer interface {
	io.Writer
	io.Reader
	io.Closer
	Reset()
	Len() int
	Empty() bool
	Bytes() []byte
}

type cancellableBytesWriter struct {
	innerWriter *bytes.Buffer
	maxCapacity uint
}

func (c *cancellableBytesWriter) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *cancellableBytesWriter) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *cancellableBytesWriter) Reset() { _ = "STUB: not implemented"; return }

func (*cancellableBytesWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (c *cancellableBytesWriter) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *cancellableBytesWriter) Empty() bool { _ = "STUB: not implemented"; return false }

func (c *cancellableBytesWriter) Bytes() []byte { _ = "STUB: not implemented"; return nil }

type cancellableGzipWriter struct {
	innerBuffer *bytes.Buffer
	innerWriter *gzip.Writer
	maxCapacity uint
	rawLen      int
}

func (c *cancellableGzipWriter) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// if we see that at a 50% compression rate, we'd be over max capacity, start flushing.

// we flush so the length of the underlying buffer is accurate.

// we find that the new content uncompressed, added to our buffer, would overflow our max capacity.

// so we create a copy of our content and add this new data, compressed, to check that it fits.

// we find that even compressed, the data overflows.

func (c *cancellableGzipWriter) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *cancellableGzipWriter) Reset() { _ = "STUB: not implemented"; return }

func (c *cancellableGzipWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (c *cancellableGzipWriter) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *cancellableGzipWriter) Empty() bool { _ = "STUB: not implemented"; return false }

func (c *cancellableGzipWriter) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// bufferPool is a pool of buffer objects.
type bufferPool struct {
	pool *sync.Pool
}

func (p bufferPool) get() buffer { _ = "STUB: not implemented"; return *new(buffer) }

func (p bufferPool) put(bf buffer) { _ = "STUB: not implemented"; return }

func newBufferPool(bufCap uint, compressionEnabled bool) bufferPool {
	_ = "STUB: not implemented"
	return *new(bufferPool)
}
