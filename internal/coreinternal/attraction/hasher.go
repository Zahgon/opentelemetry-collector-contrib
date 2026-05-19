// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attraction // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	int64ByteSize   = 8
	float64ByteSize = 8
)

var (
	byteTrue  = [1]byte{1}
	byteFalse = [1]byte{0}
)

// sha2Hasher hashes an AttributeValue using SHA2-256 and returns a
// hashed version of the attribute. In practice, this would mostly be used
// for string attributes, but we support all types for completeness/correctness
// and eliminate any surprises.
func sha2Hasher(attr pcommon.Value) { _ = "STUB: not implemented"; return }

// No-op for empty, maps, slices, bytes.

// Hex encoding is always 64 bytes for SHA-256.
