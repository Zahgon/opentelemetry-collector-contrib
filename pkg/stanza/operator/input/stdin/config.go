// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdin // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/stdin"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "stdin"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig("") })
}

// NewConfig creates a new stdin input config with default values
func NewConfig(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a stdin input operator.
type Config struct {
	helper.InputConfig `mapstructure:",squash"`
}

// Build will build a stdin input operator.
func (c *Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
