// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package router // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/router"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "router"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig config creates a new router operator config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID config creates a new router operator config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a router operator
type Config struct {
	helper.BasicConfig `mapstructure:",squash"`
	Routes             []*RouteConfig `mapstructure:"routes"`
	Default            []string       `mapstructure:"default"`
}

// RouteConfig is the configuration of a route on a router operator
type RouteConfig struct {
	helper.AttributerConfig `mapstructure:",squash"`
	Expression              string   `mapstructure:"expr"`
	OutputIDs               []string `mapstructure:"output"`
}

// Build will build a router operator from the supplied configuration
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
