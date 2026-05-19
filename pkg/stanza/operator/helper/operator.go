// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// NewBasicConfig creates a new basic config
func NewBasicConfig(operatorID, operatorType string) BasicConfig {
	_ = "STUB: not implemented"
	return *new(BasicConfig)
}

// BasicConfig provides a basic implemention for an operator config.
type BasicConfig struct {
	OperatorID   string `mapstructure:"id"`
	OperatorType string `mapstructure:"type"`
}

// ID will return the operator id.
func (c BasicConfig) ID() string { _ = "STUB: not implemented"; return "" }

// SetID will Update the operator id.
func (c *BasicConfig) SetID(id string) {
	_ = "STUB: not implemented"

	// Type will return the operator type.
	return
}

func (c BasicConfig) Type() string { _ = "STUB: not implemented"; return "" }

// Build will build a basic operator.
func (c BasicConfig) Build(set component.TelemetrySettings) (BasicOperator, error) {
	_ = "STUB: not implemented"
	return *new(BasicOperator), nil
}

// BasicOperator provides a basic implementation of an operator.
type BasicOperator struct {
	OperatorID   string
	OperatorType string
	set          component.TelemetrySettings
	// prevent unkeyed literal initialization
	_ struct{}
}

// ID will return the operator id.
func (p *BasicOperator) ID() string { _ = "STUB: not implemented"; return "" }

// Type will return the operator type.
func (p *BasicOperator) Type() string { _ = "STUB: not implemented"; return "" }

// Logger returns the operator's scoped logger.
func (p *BasicOperator) Logger() *zap.Logger { _ = "STUB: not implemented"; return nil }

// Start will start the operator.
func (*BasicOperator) Start(operator.Persister) error {
	_ = "STUB: not implemented"

	// Stop will stop the operator.
	return nil
}

func (*BasicOperator) Stop() error { _ = "STUB: not implemented"; return nil }
