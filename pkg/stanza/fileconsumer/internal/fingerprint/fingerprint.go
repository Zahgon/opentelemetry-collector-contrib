// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fingerprint // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fingerprint"

import (
	"os"

	"go.uber.org/zap"
)

const DefaultSize = 1000 // bytes

const MinSize = 16 // bytes

// Fingerprint is used to identify a file
// A file's fingerprint is the first N bytes of the file
type Fingerprint struct {
	firstBytes []byte
}

func New(first []byte) *Fingerprint { _ = "STUB: not implemented"; return nil }

// NewFromFile computes fingerprint of the given file using first 'N' bytes
// Set decompressData to true to compute fingerprint of compressed files by decompressing its data first
func NewFromFile(file *os.File, size int, decompressData bool, logger *zap.Logger) (*Fingerprint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the file is of compressed type, uncompress the data before creating its fingerprint

// Copy creates a new copy of the fingerprint
func (f Fingerprint) Copy() *Fingerprint { _ = "STUB: not implemented"; return nil }

func (f *Fingerprint) Len() int { _ = "STUB: not implemented"; return 0 }

// Equal returns true if the fingerprints have the same FirstBytes,
// false otherwise. This does not compare other aspects of the fingerprints
// because the primary purpose of a fingerprint is to convey a unique
// identity, and only the FirstBytes field contributes to this goal.
func (f Fingerprint) Equal(other *Fingerprint) bool { _ = "STUB: not implemented"; return false }

// StartsWith returns true if the fingerprints are the same
// or if the new fingerprint starts with the old one
// This is important functionality for tracking new files,
// since their initial size is typically less than that of
// a fingerprint. As the file grows, its fingerprint is updated
// until it reaches a maximum size, as configured on the operator
func (f Fingerprint) StartsWith(old *Fingerprint) bool { _ = "STUB: not implemented"; return false }

func (f *Fingerprint) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Fingerprint) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type marshal struct {
	FirstBytes []byte `json:"first_bytes"`
}

// Bytes returns a copy of the raw fingerprint bytes.
func (f *Fingerprint) Bytes() []byte { _ = "STUB: not implemented"; return nil }
