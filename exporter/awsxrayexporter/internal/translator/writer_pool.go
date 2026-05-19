// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/translator"

import (
	"bytes"
	"encoding/json"
	"sync"
)

const (
	maxBufSize = 65536
)

type writer struct {
	buffer  *bytes.Buffer
	encoder *json.Encoder
}

type writerPool struct {
	pool *sync.Pool
}

func newWriterPool(size int) *writerPool { _ = "STUB: not implemented"; return nil }

func (w *writer) Reset() { _ = "STUB: not implemented"; return }

func (w *writer) Encode(v any) error { _ = "STUB: not implemented"; return nil }

func (w *writer) String() string { _ = "STUB: not implemented"; return "" }

func (writerPool *writerPool) borrow() *writer { _ = "STUB: not implemented"; return nil }

func (writerPool *writerPool) release(w *writer) { _ = "STUB: not implemented"; return }
