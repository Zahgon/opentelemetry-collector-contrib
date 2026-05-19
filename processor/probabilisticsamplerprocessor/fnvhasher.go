// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package probabilisticsamplerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"

// computeHash creates a hash using the FNV-1a algorithm
func computeHash(b []byte, seed uint32) uint32 {
	_ = "STUB: not implemented"

	// the implementation fnv.Write() does not return an error, see hash/fnv/fnv.go
	return 0
}

// i32tob converts a seed to a byte array to be used as part of fnv.Write()
func i32tob(val uint32) []byte { _ = "STUB: not implemented"; return nil }
