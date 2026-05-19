// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build darwin

package macosunifiedloggingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/macosunifiedloggingreceiver"

// Valid field names for macOS unified logging predicates
var validPredicateFields = map[string]bool{
	"activityIdentifier":             true,
	"bootUUID":                       true,
	"category":                       true,
	"composedMessage":                true,
	"continuousNanosecondsSinceBoot": true,
	"creatorActivityIdentifier":      true,
	"creatorProcessUniqueIdentifier": true,
	"date":                           true,
	"formatString":                   true,
	"logType":                        true,
	"machContinuousTimestamp":        true,
	"parentActivityIdentifier":       true,
	"process":                        true,
	"processIdentifier":              true,
	"processImagePath":               true,
	"processImageUUID":               true,
	"sender":                         true,
	"senderImageOffset":              true,
	"senderImagePath":                true,
	"senderImageUUID":                true,
	"signpostIdentifier":             true,
	"signpostScope":                  true,
	"signpostType":                   true,
	"size":                           true,
	"subsystem":                      true,
	"threadIdentifier":               true,
	"timeToLive":                     true,
	"traceIdentifier":                true,
	"transitionActivityIdentifier":   true,
	"type":                           true,
}

// Valid event types
var validEventTypes = map[string]bool{
	"activityCreateEvent":     true,
	"activityTransitionEvent": true,
	"userActionEvent":         true,
	"traceEvent":              true,
	"logEvent":                true,
	"timesyncEvent":           true,
	"signpostEvent":           true,
	"lossEvent":               true,
	"stateEvent":              true,
}

// Valid log types
var validLogTypes = map[string]bool{
	"default": true,
	"release": true,
	"info":    true,
	"debug":   true,
	"error":   true,
	"fault":   true,
}

// Valid signpost scopes
var validSignpostScopes = map[string]bool{
	"thread":  true,
	"process": true,
	"system":  true,
}

// Valid signpost types
var validSignpostTypes = map[string]bool{
	"event": true,
	"begin": true,
	"end":   true,
}

// Valid comparison operators
var validOperators = []string{
	"AND", "&&", "&",
	"OR", "||",
	"NOT", "!",
	"!=", "<>", "==", "=",
	"<", ">", "<=", "=<", ">=", "=>",
	"BEGINSWITH",
	"CONTAINS",
	"ENDSWITH",
	"LIKE",
	"MATCHES",
}

// Validate checks the Config is valid
func (cfg *Config) Validate() error {
	_ = "STUB: not implemented"
	// Set default format if not specified
	return nil
}

// Validate format

// Validate predicate to prevent invalid characters

// Validate archive path if specified

// Store the resolved paths internally

// Validate time format if specified

// validatePredicate performs basic validation on the predicate expression
func validatePredicate(predicate *string) error {
	_ = "STUB: not implemented"

	// Check for balanced quotes
	return nil
}

// Check for balanced parentheses

// Validate that at least one valid field name appears in the predicate

// Normalize && to AND to prevent command chaining

// Normalize || to OR to prevent command chaining

// Command separator (not valid in predicates)
// Pipe (not valid in predicates - use AND/OR instead)
// Variable expansion (not valid in predicates)
// Backtick command substitution (not valid in predicates)
// Newline (not valid in predicates)
// Carriage return (not valid in predicates)
// Append redirect (not valid in predicates)
// Here document (not valid in predicates)

func hasValidEventType(predicate string) bool { _ = "STUB: not implemented"; return false }

func hasValidLogType(predicate string) bool { _ = "STUB: not implemented"; return false }

func hasValidSignpostScope(predicate string) bool { _ = "STUB: not implemented"; return false }

func hasValidSignpostType(predicate string) bool { _ = "STUB: not implemented"; return false }

// hasBalancedQuotes checks if the string has balanced double quotes
func hasBalancedQuotes(s string) bool { _ = "STUB: not implemented"; return false }

// hasBalancedParentheses checks if the string has balanced parentheses
func hasBalancedParentheses(s string) bool { _ = "STUB: not implemented"; return false }

// resolveArchivePath takes a path (potentially with glob pattern)
// and returns a list of resolved archive directory paths
func resolveArchivePath(pattern string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil,

		// Clean the pattern
		nil
}

// Check if pattern contains glob characters

// Use doublestar for glob matching

// Validate each matched path

// Skip invalid matches and continue

// Direct path without glob

// validateArchivePath checks if a path exists and is a valid archive directory
func validateArchivePath(path string) error { _ = "STUB: not implemented"; return nil }
