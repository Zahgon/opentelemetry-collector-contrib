// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"

import "github.com/klauspost/compress/zstd"

// compressFunc defines how to compress encoded telemetry data.
type compressFunc func(src []byte) []byte

var encoder, _ = zstd.NewWriter(nil)

var encoders = map[string]compressFunc{
	compressionZSTD: zstdCompress,
}

func buildCompressor(compression string) compressFunc {
	_ = "STUB: not implemented"
	return *new(compressFunc)
}

// zstdCompress compress a buffer with zstd
func zstdCompress(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// noneCompress return src
func noneCompress(src []byte) []byte { _ = "STUB: not implemented"; return nil }
