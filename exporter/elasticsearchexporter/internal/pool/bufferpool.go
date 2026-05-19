// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pool // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/pool"

import (
	"bytes"
	"io"
	"sync"
)

type BufferPool struct {
	pool *sync.Pool
}

func NewBufferPool() *BufferPool { _ = "STUB: not implemented"; return nil }

func (w *BufferPool) NewPooledBuffer() PooledBuffer {
	_ = "STUB: not implemented"
	return *new(PooledBuffer)
}

type PooledBuffer struct {
	Buffer *bytes.Buffer
	pool   *sync.Pool
}

func (p PooledBuffer) Recycle() { _ = "STUB: not implemented"; return }

func (p PooledBuffer) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
