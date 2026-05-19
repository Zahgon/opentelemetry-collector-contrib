// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package crlf // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/testing/testar/crlf"

const (
	CR   = '\r'
	LF   = '\n'
	CRLF = "\r\n"
)

// Strip turns CRLF line endings (\r\n) into LF (\n)
func Strip(data []byte) []byte { _ = "STUB: not implemented"; return nil }

// Join concats all lines with the [CRLF] separator
func Join(lines ...string) []byte { _ = "STUB: not implemented"; return nil }
