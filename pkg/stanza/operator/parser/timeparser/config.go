// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package timeparser // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/timeparser"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "time_parser"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new time parser config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new time parser config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a time parser operator.
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	helper.TimeParser        `mapstructure:",omitempty,squash"`
}

func (c *Config) Unmarshal(component *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// Build will build a time parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
