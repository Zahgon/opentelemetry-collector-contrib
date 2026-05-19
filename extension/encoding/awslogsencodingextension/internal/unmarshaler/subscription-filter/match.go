// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package subscriptionfilter // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/subscription-filter"

// catchAllPattern is the bare wildcard that matches any single-segment path.
// Use this as a fallback/catch-all entry.
const catchAllPattern = "*"

// matchPrefixWithWildcard reports whether targetParts has patternParts as a
// prefix. Each pattern segment may be:
//   - an exact literal — compared by string equality
//   - "*" — matches any single segment
//   - an affix wildcard within a single segment:
//     "foo*" / "*foo" / "*foo*"
//
// Mid-segment globs such as "foo*bar" are not supported and never match.
// Callers split paths on "/" before invoking. For non-path strings such as a
// CloudWatch log_stream name, callers pass a single-element slice.
// The target may contain extra trailing segments beyond the pattern's length.
func matchPrefixWithWildcard(targetParts, patternParts []string) bool {
	_ = "STUB: not implemented"
	// Target must have at least as many segments as the pattern.
	return false
}

// wildcard matches any single segment

// matchAffix matches a single segment against an affix-wildcard pattern.
// Supported pattern shapes:
//   - "foo*"   — target must start with "foo"
//   - "*foo"   — target must end with "foo"
//   - "*foo*"  — target must contain "foo"
//
// Any other use of "*" inside the pattern (mid-segment, more than two
// occurrences) returns false.
func matchAffix(target, pattern string) bool { _ = "STUB: not implemented"; return false }

// comparePatternSpecificity orders two pre-split patterns by specificity.
// It returns a negative value if a is more specific than b, positive if b
// is more specific, and zero if they are equally specific.
//
// Rules:
//   - At each position, an exact segment is more specific than an affix
//     wildcard, which is more specific than the bare "*" wildcard.
//   - When all shared positions are equally specific, the longer pattern wins.
func comparePatternSpecificity(partsA, partsB []string) int { _ = "STUB: not implemented"; return 0 }

// Higher score = more specific = should sort earlier.

// Longer pattern is more specific.

// segmentSpecificity returns a score for a single pattern segment.
// Higher score = more specific. The score is:
//   - 2 for an exact literal (no '*')
//   - 1 for an affix wildcard (contains '*' but isn't bare "*")
//   - 0 for the bare "*"
func segmentSpecificity(segment string) int { _ = "STUB: not implemented"; return 0 }
