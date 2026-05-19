// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// NewAttributerConfig creates a new attributer config with default values
func NewAttributerConfig() AttributerConfig {
	_ = "STUB: not implemented"
	return *new(AttributerConfig)
}

// AttributerConfig is the configuration of a attributer
type AttributerConfig struct {
	Attributes map[string]ExprStringConfig `mapstructure:"attributes"`
}

// Build will build a attributer from the supplied configuration
func (c AttributerConfig) Build() (Attributer, error) {
	_ = "STUB: not implemented"
	return *new(Attributer), nil
}

// Attributer is a helper that adds attributes to an entry
type Attributer struct {
	attributes map[string]*ExprString
}

// Attribute will add attributes to an entry
func (l *Attributer) Attribute(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }
