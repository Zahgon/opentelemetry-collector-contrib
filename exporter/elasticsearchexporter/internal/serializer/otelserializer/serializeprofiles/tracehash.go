// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package serializeprofiles // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer/serializeprofiles"

import "go.opentelemetry.io/ebpf-profiler/libpf/basehash"

// traceHash represents the unique hash of a trace
type traceHash struct {
	basehash.Hash128
}

func newTraceHash(hi, lo uint64) traceHash { _ = "STUB: not implemented"; return *new(traceHash) }

// traceHashFromBytes parses a byte slice of a trace hash into the internal data representation.
func traceHashFromBytes(b []byte) (traceHash, error) {
	_ = "STUB: not implemented"
	return *new(traceHash), nil
}
