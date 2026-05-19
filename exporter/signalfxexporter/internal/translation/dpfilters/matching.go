// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dpfilters // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation/dpfilters"

import (
	"regexp"

	"github.com/gobwas/glob"
)

// Contains all of the logic for glob and regex based filtering.

func isGlobbed(s string) bool { _ = "STUB: not implemented"; return false }

func isRegex(s string) bool { _ = "STUB: not implemented"; return false }

// remove the bracketing slashes for a regex.
func stripSlashes(s string) string { _ = "STUB: not implemented"; return "" }

// stripNegation checks if a string is prefixed with "!"
// and will returned the stripped string and true if so
// else, return original value and false.
func stripNegation(value string) (string, bool) { _ = "STUB: not implemented"; return "", false }

type matcher interface {
	// Returns whether the string matched and whether it was a negated match.
	Matches(s string) (bool, bool)
}

type regexMatcher struct {
	re      *regexp.Regexp
	negated bool
}

var _ matcher = (*regexMatcher)(nil)

func (m *regexMatcher) Matches(s string) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

type globMatcher struct {
	glob    glob.Glob
	negated bool
}

var _ matcher = &globMatcher{}

func (m *globMatcher) Matches(s string) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}
