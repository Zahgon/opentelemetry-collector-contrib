// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package parseutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/parseutils"

// SplitString will split the input on the delimiter and return the resulting slice while respecting quotes. Outer quotes are stripped.
// Use in place of `strings.Split` when quotes need to be respected.
// Requires `delimiter` not be an empty string
func SplitString(input, delimiter string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 0 means we are not in quotes

// delimiter
// leading || trailing delimiter; ignore

// consider quote termination so long as previous character wasn't backslash
// start of quote

// end of quote

// Only if we weren't escaped could the next character result in escaped state
// potentially escaping next character

// check for closed quotes

// avoid adding empty value bc of a trailing delimiter

// ParseKeyValuePairs will split each string in `pairs` on the `delimiter` into a key and value string that get added to a map and returned.
func ParseKeyValuePairs(pairs []string, delimiter string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
