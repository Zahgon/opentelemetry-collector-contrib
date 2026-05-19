// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package model // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/model"

import (
	"sync"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

// attrHashBuf is a pooled byte buffer used to compute a 128-bit hash of a
// filtered attribute set without allocating a pcommon.Map.
type attrHashBuf struct {
	buf []byte
}

var attrHashBufPool = sync.Pool{
	New: func() any {
		return &attrHashBuf{buf: make([]byte, 0, 256)}
	},
}

// sum128 returns a 128-bit hash of b.buf using the same two-pass xxhash
// algorithm as pdatautil.MapHash.
func (b *attrHashBuf) sum128() [16]byte { _ = "STUB: not implemented"; return nil }

// appendAttrValue appends a type-tagged encoding of v to buf.
func appendAttrValue(buf []byte, v pcommon.Value) []byte { _ = "STUB: not implemented"; return nil }

// bytes, map, slice: encode as string representation (rare for attributes)
