// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dpfilters // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation/dpfilters"

// StringFilter will match if any one of the given strings is a match.
type StringFilter struct {
	staticSet        map[string]bool
	regexps          []regexMatcher
	globs            []globMatcher
	anyStaticNegated bool
}

// NewStringFilter returns a filter that can match against the provided items.
func NewStringFilter(items []string) (*StringFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Matches if s is positively matched by the filter items OR
// if it is positively matched by a non-glob/regex pattern exactly
// and is negated as well.  See the unit tests for examples.
func (f *StringFilter) Matches(s string) bool { _ = "STUB: not implemented"; return false }

// If a metric is negated and it matched it won't match anything else by
// definition.

func (f *StringFilter) UnmarshalText(in []byte) error { _ = "STUB: not implemented"; return nil }
