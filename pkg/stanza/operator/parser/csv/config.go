// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package csv // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/csv"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "csv_parser"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new csv parser config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new csv parser config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a csv parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`

	Header          string `mapstructure:"header"`
	HeaderDelimiter string `mapstructure:"header_delimiter"`
	HeaderAttribute string `mapstructure:"header_attribute"`
	FieldDelimiter  string `mapstructure:"delimiter"`
	LazyQuotes      bool   `mapstructure:"lazy_quotes"`
	IgnoreQuotes    bool   `mapstructure:"ignore_quotes"`
}

// Build will build a csv parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
