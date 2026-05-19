// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package move // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/move"

import (
	"errors"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "move"

var (
	errMissingFrom = errors.New("copy: missing from field")
	errMissingTo   = errors.New("copy: missing to field")
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new move operator config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new move operator config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a move operator
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	From                     entry.Field `mapstructure:"from"`
	To                       entry.Field `mapstructure:"to"`
}

// Build will build a Move operator from the supplied configuration
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
