// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// NewIdentifierConfig creates a new identifier config with default values
func NewIdentifierConfig() IdentifierConfig {
	_ = "STUB: not implemented"
	return *new(IdentifierConfig)
}

// IdentifierConfig is the configuration of a resource identifier
type IdentifierConfig struct {
	Resource map[string]ExprStringConfig `mapstructure:"resource"`
}

// Build will build an identifier from the supplied configuration
func (c IdentifierConfig) Build() (Identifier, error) {
	_ = "STUB: not implemented"
	return *new(Identifier), nil
}

// Identifier is a helper that adds values to the resource of an entry
type Identifier struct {
	resource map[string]*ExprString
}

// Identify will add values to the resource of an entry
func (i *Identifier) Identify(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }
