// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package jsonarray // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/jsonarray"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const (
	operatorType    = "json_array_parser"
	headerDelimiter = ","
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new json array parser config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new json array parser config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a json array parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`
	Header              string `mapstructure:"header"`
}

// Build will build a json array parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
