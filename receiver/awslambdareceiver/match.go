// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awslambdareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awslambdareceiver"

// catchAllPattern is a special pattern value that matches any path.
// Use this as a fallback/catch-all encoding entry.
const catchAllPattern = "*"

// matchPrefixWithWildcard checks whether targetParts has patternParts as a prefix.
// Each pattern segment must be either an exact literal or "*" (matches any single segment).
// The target may have more segments than the pattern (prefix semantics).
//
// The caller is responsible for splitting the object key and the pattern on "/" before calling.
//
// Examples:
//
//	matchPrefixWithWildcard(
//	    []string{"AWSLogs", "123456789012", "vpcflowlogs", "file.gz"},
//	    []string{"AWSLogs", "*", "vpcflowlogs"},
//	) => true
//	matchPrefixWithWildcard([]string{"any", "path"}, []string{"*"}) => true
func matchPrefixWithWildcard(targetParts, patternParts []string) bool {
	_ = "STUB: not implemented"
	// Target must have at least as many segments as the pattern.
	return false
}

// wildcard matches any single segment

// comparePatternSpecificity compares two pre-split patterns by specificity.
// Returns negative if a is more specific, positive if b is more specific, 0 if equal.
//
// Rules (evaluated left-to-right per segment):
//   - Exact segment beats "*" at the same position.
//   - Longer pattern (more segments) beats a shorter one.
func comparePatternSpecificity(partsA, partsB []string) int { _ = "STUB: not implemented"; return 0 }

// Higher score = more specific = should sort earlier.

// Longer pattern is more specific.

// segmentSpecificity returns a specificity score for a single pattern segment.
// Higher score = more specific.
//   - 1: exact literal
//   - 0: wildcard "*"
func segmentSpecificity(segment string) int { _ = "STUB: not implemented"; return 0 }

// isCatchAllPattern returns true if the pattern is the bare catch-all "*".
func isCatchAllPattern(pattern string) bool { _ = "STUB: not implemented"; return false }
