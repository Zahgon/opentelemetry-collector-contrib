// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// map[string or int input]sev-level
func getBuiltinMapping(name string) severityMap {
	_ = "STUB: not implemented"
	return *new(severityMap)
}

// Add some additional values that are automatically recognized

func (m severityMap) add(severity entry.Severity, parseableValues ...string) {
	_ = "STUB: not implemented"
	return
}

const (
	// HTTP2xx is a special key that is represents a range from 200 to 299. Literal value is "2xx"
	HTTP2xx = "2xx"

	// HTTP3xx is a special key that is represents a range from 300 to 399. Literal value is "3xx"
	HTTP3xx = "3xx"

	// HTTP4xx is a special key that is represents a range from 400 to 499. Literal value is "4xx"
	HTTP4xx = "4xx"

	// HTTP5xx is a special key that is represents a range from 500 to 599. Literal value is "5xx"
	HTTP5xx = "5xx"
)

// NewSeverityConfig creates a new severity parser config
func NewSeverityConfig() SeverityConfig {
	_ = "STUB: not implemented"
	return *

	// SeverityConfig allows users to specify how to parse a severity from a field.
	new(SeverityConfig)
}

type SeverityConfig struct {
	ParseFrom     *entry.Field   `mapstructure:"parse_from,omitempty"`
	Preset        string         `mapstructure:"preset,omitempty"`
	Mapping       map[string]any `mapstructure:"mapping,omitempty"`
	OverwriteText bool           `mapstructure:"overwrite_text,omitempty"`
}

// Build builds a SeverityParser from a SeverityConfig
func (c *SeverityConfig) Build(_ component.TelemetrySettings) (SeverityParser, error) {
	_ = "STUB: not implemented"
	return *new(SeverityParser), nil
}

// check before any

func validateSeverity(severity any) (entry.Severity, error) {
	_ = "STUB: not implemented"
	return *new(entry.Severity), nil
}

func isRange(value any) (int, int, bool) { _ = "STUB: not implemented"; return 0, 0, false }

func expandRange(minR, maxR int) []string { _ = "STUB: not implemented"; return nil }

func parseableValues(value any) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// store as string because we will compare as string
