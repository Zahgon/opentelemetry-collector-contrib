// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package generate // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/generate"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "generate_input"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig("") })
}

// NewConfig creates a new generate input config with default values
func NewConfig(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a generate input operator.
type Config struct {
	helper.InputConfig `mapstructure:",squash"`
	Entry              entry.Entry `mapstructure:"entry"`
	Count              int         `mapstructure:"count"`
	Static             bool        `mapstructure:"static"`
}

// Build will build a generate input operator.
func (c *Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
