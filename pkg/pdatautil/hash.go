// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pdatautil // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil"

import (
	"sync"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

var (
	extraByte       = []byte{'\xf3'}
	keyPrefix       = []byte{'\xf4'}
	valEmpty        = []byte{'\xf5'}
	valBytesPrefix  = []byte{'\xf6'}
	valStrPrefix    = []byte{'\xf7'}
	valBoolTrue     = []byte{'\xf8'}
	valBoolFalse    = []byte{'\xf9'}
	valIntPrefix    = []byte{'\xfa'}
	valDoublePrefix = []byte{'\xfb'}
	valMapPrefix    = []byte{'\xfc'}
	valMapSuffix    = []byte{'\xfd'}
	valSlicePrefix  = []byte{'\xfe'}
	valSliceSuffix  = []byte{'\xff'}

	emptyHash = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

// HashOption is a function that sets an option on the hash calculation.
type HashOption func(*hashWriter)

// WithMap adds a map to the hash calculation.
func WithMap(m pcommon.Map) HashOption { _ = "STUB: not implemented"; return *new(HashOption) }

// WithValue adds a value to the hash calculation.
func WithValue(v pcommon.Value) HashOption { _ = "STUB: not implemented"; return *new(HashOption) }

// WithString adds a string to the hash calculation.
func WithString(s string) HashOption { _ = "STUB: not implemented"; return *new(HashOption) }

type hashWriter struct {
	byteBuf []byte
	keysBuf []string
}

func newHashWriter() *hashWriter { _ = "STUB: not implemented"; return nil }

var hashWriterPool = &sync.Pool{
	New: func() any { return newHashWriter() },
}

// Hash generates a hash for the provided options and returns the computed hash as a [16]byte.
func Hash(opts ...HashOption) [16]byte { _ = "STUB: not implemented"; return nil }

// Hash64 generates a hash for the provided options and returns the computed hash as a uint64.
func Hash64(opts ...HashOption) uint64 { _ = "STUB: not implemented"; return 0 }

// MapHash return a hash for the provided map.
// Maps with the same underlying key/value pairs in different order produce the same deterministic hash value.
func MapHash(m pcommon.Map) [16]byte { _ = "STUB: not implemented"; return nil }

// ValueHash return a hash for the provided pcommon.Value.
func ValueHash(v pcommon.Value) [16]byte { _ = "STUB: not implemented"; return nil }

func (hw *hashWriter) writeMapHash(m pcommon.Map) {
	_ = "STUB: not implemented"
	// For each recursive call into this function we want to preserve the previous buffer state
	// while also adding new keys to the buffer. nextIndex is the index of the first new key
	// added to the buffer for this call of the function.
	// This also works for the first non-recursive call of this function because the buffer is always empty
	// on the first call due to it being cleared of any added keys at then end of the function.
	return
}

// Get only the newly added keys from the buffer by slicing the buffer from nextIndex to the end

// Remove all keys that were added to the buffer during this call of the function

func (hw *hashWriter) writeValueHash(v pcommon.Value) { _ = "STUB: not implemented"; return }

func (hw *hashWriter) writeString(s string) { _ = "STUB: not implemented"; return }

// hashSum128 returns a [16]byte hash sum.
func (hw *hashWriter) hashSum128() [16]byte { _ = "STUB: not implemented"; return nil }

// Append an extra byte to generate another part of the hash sum
