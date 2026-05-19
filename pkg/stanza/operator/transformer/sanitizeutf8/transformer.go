// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sanitizeutf8 // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/sanitizeutf8"

import (
	"context"
	"unicode/utf8"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const (
	operatorType           = "sanitize_utf8"
	invalidCharacterMarker = string(utf8.RuneError) // Unicode replacement character
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new sanitize_utf8 config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new sanitize_utf8 config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a sanitize_utf8 operator
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	Field                    entry.Field `mapstructure:"field"`
}

// Build creates a new Transformer from a config
func (c *Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

// Transformer is an operator that replaces invalid UTF-8 characters in a field
type Transformer struct {
	helper.TransformerOperator
	field entry.Field
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transformer) Process(ctx context.Context, e *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transformer) ProcessWith(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }
