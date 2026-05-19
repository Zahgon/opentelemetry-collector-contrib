// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// NewTraceParser creates a new trace parser with default values
func NewTraceParser() TraceParser { _ = "STUB: not implemented"; return *new(TraceParser) }

// TraceParser is a helper that parses trace spans (and flags) onto an entry.
type TraceParser struct {
	TraceID    *TraceIDConfig    `mapstructure:"trace_id,omitempty"`
	SpanID     *SpanIDConfig     `mapstructure:"span_id,omitempty"`
	TraceFlags *TraceFlagsConfig `mapstructure:"trace_flags,omitempty"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type TraceIDConfig struct {
	ParseFrom *entry.Field `mapstructure:"parse_from,omitempty"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type SpanIDConfig struct {
	ParseFrom *entry.Field `mapstructure:"parse_from,omitempty"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type TraceFlagsConfig struct {
	ParseFrom *entry.Field `mapstructure:"parse_from,omitempty"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// Validate validates a TraceParser, and reconfigures it if necessary
func (t *TraceParser) Validate() error { _ = "STUB: not implemented"; return nil }

// Best effort hex parsing for trace, spans and flags
func parseHexField(entry *entry.Entry, field *entry.Field) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse will parse a trace (trace_id, span_id and flags) from a field and attach it to the entry
func (t *TraceParser) Parse(entry *entry.Entry) error { _ = "STUB: not implemented"; return nil }
