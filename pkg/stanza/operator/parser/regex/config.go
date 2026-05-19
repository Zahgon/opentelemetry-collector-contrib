// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package regex // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/regex"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "regex_parser"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new regex parser config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new regex parser config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a regex parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`

	Regex string `mapstructure:"regex"`

	Cache struct {
		Size uint16 `mapstructure:"size"`
	} `mapstructure:"cache"`
}

// Build will build a regex parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
