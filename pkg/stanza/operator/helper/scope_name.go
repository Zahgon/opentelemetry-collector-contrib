// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// ScopeNameParser is a helper that parses severity onto an entry.
type ScopeNameParser struct {
	ParseFrom entry.Field `mapstructure:"parse_from,omitempty"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// NewScopeNameParser creates a new scope parser with default values
func NewScopeNameParser() ScopeNameParser {
	_ = "STUB: not implemented"
	return *

	// Parse will parse severity from a field and attach it to the entry
	new(ScopeNameParser)
}

func (p *ScopeNameParser) Parse(ent *entry.Entry) error { _ = "STUB: not implemented"; return nil }
