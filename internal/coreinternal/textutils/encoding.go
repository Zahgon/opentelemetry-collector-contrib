// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package textutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/textutils"

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

var encodingOverrides = map[string]encoding.Encoding{
	"utf-16":    unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM),
	"utf16":     unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM),
	"utf-8":     unicode.UTF8,
	"utf8":      unicode.UTF8,
	"utf-8-raw": UTF8Raw,
	"utf8-raw":  UTF8Raw,
	"ascii":     unicode.UTF8,
	"us-ascii":  unicode.UTF8,
	"nop":       encoding.Nop,
	"":          unicode.UTF8,
}

func LookupEncoding(enc string) (encoding.Encoding, error) {
	_ = "STUB: not implemented"
	return *new(encoding.Encoding), nil
}

func IsNop(enc string) bool { _ = "STUB: not implemented"; return false }

// DecodeAsString converts the given encoded bytes using the given decoder. It returns the converted
// bytes or nil, err if any error occurred.
func DecodeAsString(decoder *encoding.Decoder, buf []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// UnsafeBytesAsString converts the byte array to string.
// This function must be called iff the input buffer is not going to be re-used after.
func UnsafeBytesAsString(buf []byte) string { _ = "STUB: not implemented"; return "" }
