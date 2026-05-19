// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package docker // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/docker"

import (
	"regexp"

	"github.com/gobwas/glob"
)

type stringMatcher struct {
	standardItems       map[string]bool
	anyNegatedStandards bool
	regexItems          []regexItem
	globItems           []globbedItem
}

// This utility defines a regex as
// any string between two '/' characters
// with the option of a leading '!' to
// signify negation.
type regexItem struct {
	re        *regexp.Regexp
	isNegated bool
}

func isRegex(s string) bool { _ = "STUB: not implemented"; return false }

type globbedItem struct {
	glob      glob.Glob
	isNegated bool
}

func isGlobbed(s string) bool { _ = "STUB: not implemented"; return false }

func newStringMatcher(items []string) (*stringMatcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// by definition this must lead and end with '/' chars

// isNegatedItem strips a leading '!' and returns
// the remaining substring and true.  If no leading
// '!' is found, it returns the input string and false.
func isNegatedItem(value string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (f *stringMatcher) matches(s string) bool { _ = "STUB: not implemented"; return false }

// If negated standard item "!something" is provided
// and "anything else" is evaluated we will always match.
