// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remove // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/remove"

import (
	"errors"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "remove"

var errMissingField = errors.New("remove: field is empty")

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new remove operator config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new remove operator config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a remove operator
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`

	Field rootableField `mapstructure:"field"`
}

// Build will build a Remove operator from the supplied configuration
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
