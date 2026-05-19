// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package compress // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/compress"

type Compressor func(in []byte) ([]byte, error)

func NewCompressor(format string) (Compressor, error) {
	_ = "STUB: not implemented"
	return *new(Compressor), nil
}

func flateCompressor(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func gzipCompressor(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func zlibCompressor(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func noopCompressor(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
