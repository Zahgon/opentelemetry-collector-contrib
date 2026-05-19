// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package subscriptionfilter // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/subscription-filter"

import (
	"go.opentelemetry.io/collector/component"
)

// PayloadMode controls how each CloudWatch subscription-filter envelope is
// handed to the inner encoding. See README for semantics.
type PayloadMode string

const (
	PayloadMessage  PayloadMode = "message"
	PayloadEnvelope PayloadMode = "envelope"
)

// DefaultPayloadMode is applied when a stream is matched and Payload is
// unset and the stream's name has no default in defaultCWPatterns.
const DefaultPayloadMode = PayloadMessage

// CloudWatchStream maps a logGroup and/or logStream pattern (or a known
// service name) to an inner encoding extension.
type CloudWatchStream struct {
	Name             string       `mapstructure:"name"`
	Encoding         component.ID `mapstructure:"encoding"`
	LogGroupPattern  string       `mapstructure:"log_group_pattern"`
	LogStreamPattern string       `mapstructure:"log_stream_pattern"`
	Payload          PayloadMode  `mapstructure:"payload"`
}

var defaultCWPatterns = map[string]CloudWatchStream{
	"vpcflow":    {Name: "vpcflow", LogStreamPattern: "eni-*", Payload: PayloadEnvelope},
	"cloudtrail": {Name: "cloudtrail", LogStreamPattern: "*_CloudTrail_*", Payload: PayloadEnvelope},
	"lambda":     {Name: "lambda", LogGroupPattern: "/aws/lambda/*", Payload: PayloadMessage},
	"waf":        {Name: "waf", LogGroupPattern: "aws-waf-logs-*", Payload: PayloadMessage},
	"rds":        {Name: "rds", LogGroupPattern: "/aws/rds/instance/*/*", Payload: PayloadMessage},
	"eks":        {Name: "eks", LogGroupPattern: "/aws/eks/*", Payload: PayloadMessage},
	"apigateway": {Name: "apigateway", LogGroupPattern: "API-Gateway-Execution-Logs_*", Payload: PayloadMessage},
}

// withDefaults fills in unset fields from defaultCWPatterns. User-supplied
// values are never overwritten. Payload always ends up non-empty:
// per-name default if available, otherwise DefaultPayloadMode.
func (r CloudWatchStream) withDefaults() CloudWatchStream {
	_ = "STUB: not implemented"
	return *new(CloudWatchStream)
}

// ValidateStreams reports configuration errors in a list of CloudWatch streams:
// missing name, missing encoding, missing pattern when the name has no
// defaults, invalid payload mode, and duplicate names.
func ValidateStreams(streams []CloudWatchStream) error { _ = "STUB: not implemented"; return nil }

// Name is required: it identifies the stream in logs / errors and
// is the deduplication key.

// Encoding is required.

// When the user supplied no patterns, the name must be a known
// default; otherwise the stream has no way to match anything.

// Payload, if set, must be a known mode.

// Duplicate name.

// sortStreams returns a copy of routes ordered by routing precedence:
//
//  1. Catch-all entries (a pattern equal to "*") are placed last.
//  2. Entries with a log_group_pattern come before entries that use only
//     log_stream_pattern.
//  3. Within each group, more-specific patterns come first
//     (see comparePatternSpecificity).
//
// Defaults are applied to each stream via withDefaults before sorting, so the
// returned slice has all patterns populated for known names.
//
// Stable order is preserved among entries that compare equal.
func sortStreams(streams []CloudWatchStream) []CloudWatchStream {
	_ = "STUB: not implemented"
	return nil
}

// Pre-split non-empty patterns once and key by pattern string so the cache
// remains correct as the sort swaps elements.

// compareStreams implements the three-level routing precedence rule.
func compareStreams(a, b CloudWatchStream, splitCache map[string][]string) int {
	_ = "STUB: not implemented"
	// Level 1: catch-all "*" goes last.
	return 0
}

// Level 2: log_group_pattern before log_stream-only entries.

// Level 3: more-specific pattern first.

// resolveStreams applies defaults, sorts by precedence, and resolves each
// stream's component.ID against host.GetExtensions().
func resolveStreams(
	streams []CloudWatchStream,
	host component.Host,
	selfID component.ID,
) ([]route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
