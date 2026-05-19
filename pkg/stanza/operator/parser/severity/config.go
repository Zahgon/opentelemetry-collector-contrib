// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package severity // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/severity"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "severity_parser"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new severity parser config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new severity parser config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a severity parser operator.
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	helper.SeverityConfig    `mapstructure:",omitempty,squash"`
}

// Build will build a severity parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
