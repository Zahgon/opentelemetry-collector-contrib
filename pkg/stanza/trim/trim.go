// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package trim // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/trim"

import (
	"bufio"
)

type Func func([]byte) []byte

func WithFunc(splitFunc bufio.SplitFunc, trimFunc Func) bufio.SplitFunc {
	_ = "STUB: not implemented"
	return *new(bufio.SplitFunc)
}

type Config struct {
	PreserveLeading  bool `mapstructure:"preserve_leading_whitespaces,omitempty"`
	PreserveTrailing bool `mapstructure:"preserve_trailing_whitespaces,omitempty"`
}

func (c Config) Func() Func { _ = "STUB: not implemented"; return *new(Func) }

func Nop(token []byte) []byte { _ = "STUB: not implemented"; return nil }

func isSpace(c byte) bool { _ = "STUB: not implemented"; return false }

func Leading(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func Trailing(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func Whitespace(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func ToLength(splitFunc bufio.SplitFunc, maxLength int) bufio.SplitFunc {
	_ = "STUB: not implemented"
	return *new(bufio.SplitFunc)
}

// No token was found, but we have enough data to return a token of max length.

// A token was found but it is longer than the max length.

// ToLengthWithTruncate wraps a bufio.SplitFunc to truncate tokens that exceed maxLength.
// Unlike ToLength which splits oversized content into multiple tokens, this function
// returns only the truncated portion (up to maxLength) and advances past the entire
// original content, effectively dropping the remainder.
// The skipping parameter is a pointer to a bool that tracks whether we're currently
// skipping the remainder of an oversized entry. This allows the state to be persisted
// across multiple reader recreations (e.g., between poll cycles).
func ToLengthWithTruncate(splitFunc bufio.SplitFunc, maxLength int, skipping *bool) bufio.SplitFunc {
	_ = "STUB: not implemented"
	return *new(bufio.SplitFunc)
}

// Use local state for lastDataLen tracking (not persisted across reader recreations)

// If we're in skip mode, use splitFunc to find the next entry boundary.
// This is important for multiline patterns where entries can contain newlines.

// splitFunc found a boundary or returned a token.
// This token is the remainder of the truncated entry - discard it.

// splitFunc needs more data but didn't return anything.
// If buffer appears to be at capacity, we must advance to avoid "token too long" error.
// Buffer is at capacity if: dataLen == maxLength (buffer limited to maxLength),
// OR dataLen == lastDataLen (buffer couldn't grow between calls).

// Buffer at capacity, skip all available data and stay in skip mode

// At EOF, done skipping

// Request more data, stay in skip mode

// No token was found (splitFunc needs more data to find entry boundary),
// but we have enough data to exceed maxLength.

// Determine if we should truncate now or wait for more data.
// Truncate if:
// - dataLen == maxLength: buffer is exactly at max capacity (common case for limited buffers)
// - dataLen == lastDataLen: buffer couldn't grow between calls
// - atEOF: no more data coming

// Buffer is at max capacity or EOF, truncate and enter skip mode.

// Buffer is larger than maxLength but might still grow.
// Let scanner try to read more data.

// A token was found but it is longer than the max length.
// Return truncated token but advance past the entire original token.
