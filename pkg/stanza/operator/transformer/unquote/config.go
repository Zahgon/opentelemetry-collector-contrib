// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unquote // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/unquote"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "unquote"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new unquote config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new unquote config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of an unquote operator.
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	Field                    entry.Field `mapstructure:"field"`
}

// Build will build an unquote operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
