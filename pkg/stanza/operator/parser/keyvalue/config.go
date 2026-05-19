// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package keyvalue // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/keyvalue"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "key_value_parser"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new key value parser config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new key value parser config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a key value parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`

	Delimiter     string `mapstructure:"delimiter"`
	PairDelimiter string `mapstructure:"pair_delimiter"`
}

// Build will build a key value parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
