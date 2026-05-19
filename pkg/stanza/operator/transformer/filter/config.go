// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/filter"

import (
	"crypto/rand"
	"math/big"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "filter"

var (
	upperBound = big.NewInt(1000)
	randInt    = rand.Int // allow override for testing
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a filter operator config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a filter operator config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a filter operator
type Config struct {
	helper.TransformerConfig `mapstructure:",squash"`
	Expression               string  `mapstructure:"expr"`
	DropRatio                float64 `mapstructure:"drop_ratio"`
}

// Build will build a filter operator from the supplied configuration
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
