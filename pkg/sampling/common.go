// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"

import (
	"errors"
	"io"
)

// KV represents a key-value parsed from a section of the TraceState.
type KV struct {
	Key   string
	Value string
}

// ErrTraceStateSize is returned when a TraceState is over its
// size limit, as specified by W3C.
var ErrTraceStateSize = errors.New("invalid tracestate size")

// keyValueScanner defines distinct scanner behaviors for lists of
// key-values.
type keyValueScanner struct {
	// maxItems is 32 or -1
	maxItems int
	// trim is set if OWS (optional whitespace) should be removed
	trim bool
	// separator is , or ;
	separator byte
	// equality is = or :
	equality byte
}

// commonTraceState is embedded in both W3C and OTel trace states.
type commonTraceState struct {
	kvs []KV
}

// ExtraValues returns additional values are carried in this
// tracestate object (W3C or OpenTelemetry).
func (cts commonTraceState) ExtraValues() []KV {
	_ = "STUB: not implemented"

	// trimOws removes optional whitespace on both ends of a string.
	// this uses the strict definition for optional whitespace tiven
	// in https://www.w3.org/TR/trace-context/#tracestate-header-field-values
	return nil
}

func trimOws(input string) string { _ = "STUB: not implemented"; return "" }

// scanKeyValues is common code to scan either W3C or OTel tracestate
// entries, as parameterized in the keyValueScanner struct.
func (s keyValueScanner) scanKeyValues(input string, f func(key, value string) error) error {
	_ = "STUB: not implemented"
	return nil
}

// W3C specifies max 32 entries, tested here
// instead of via the regexp.

// Trim only required for W3C; OTel does not
// specify whitespace for its value encoding.

// W3C allows empty list members.

// We expect to find the `s.equality`
// character in this string because we have
// already validated the whole input syntax
// before calling this parser.  I.e., this can
// never happen, and if it did, the result
// would be to skip malformed entries.

// serializer assists with checking and combining errors from
// (io.StringWriter).WriteString().
type serializer struct {
	writer io.StringWriter
	err    error
}

// write handles errors from io.StringWriter.
func (ser *serializer) write(str string) { _ = "STUB: not implemented"; return }

// check handles errors (e.g., from another serializer).
func (ser *serializer) check(err error) { _ = "STUB: not implemented"; return }

// =============================================================================
// Character validation functions shared by W3C and OTel tracestate parsers.
// These hand-written validators are significantly faster than regex-based
// validation (30-60x speedup).
// =============================================================================

// isLcAlpha returns true if c is a lowercase ASCII letter (a-z).
func isLcAlpha(c byte) bool { _ = "STUB: not implemented"; return false }

// isLcAlphaNum returns true if c is a lowercase ASCII letter or digit.
func isLcAlphaNum(c byte) bool { _ = "STUB: not implemented"; return false }

// isValidKeyChar returns true if c is valid in a W3C tracestate key
// (lowercase alphanumeric or one of: _ - * /).
func isValidKeyChar(c byte) bool { _ = "STUB: not implemented"; return false }

// isValidKeyChars returns true if all characters in s are valid W3C key characters.
func isValidKeyChars(s string) bool { _ = "STUB: not implemented"; return false }
