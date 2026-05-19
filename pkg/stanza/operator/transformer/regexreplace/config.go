// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package regexreplace // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/regexreplace"

import (
	"regexp"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "regex_replace"

// derived from https://en.wikipedia.org/wiki/ANSI_escape_code#CSIsection
var ansiCsiEscapeRegex = regexp.MustCompile(`\x1B\[[\x30-\x3F]*[\x20-\x2F]*[\x40-\x7E]`)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new ansi_control_sequences config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new ansi_control_sequences config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of an ansi_control_sequences operator.
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	RegexName                string      `mapstructure:"regex_name"`
	Regex                    string      `mapstructure:"regex"`
	ReplaceWith              string      `mapstructure:"replace_with"`
	Field                    entry.Field `mapstructure:"field"`
}

func (c *Config) getRegexp() (*regexp.Regexp, error) { _ = "STUB: not implemented"; return nil, nil }

// Build will build an ansi_control_sequences operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
