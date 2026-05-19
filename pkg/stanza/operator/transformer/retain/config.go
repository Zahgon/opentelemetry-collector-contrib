// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package retain // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/retain"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "retain"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new retain operator config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new retain operator config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a retain operator
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	Fields                   []entry.Field `mapstructure:"fields"`
}

// Build will build a retain operator from the supplied configuration
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
