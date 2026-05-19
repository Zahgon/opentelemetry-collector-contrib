// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// NewOutputConfig creates a new output config
func NewOutputConfig(operatorID, operatorType string) OutputConfig {
	_ = "STUB: not implemented"
	return *new(OutputConfig)
}

// OutputConfig provides a basic implementation of an output operator config.
type OutputConfig struct {
	BasicConfig `mapstructure:",squash"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// Build will build an output operator.
func (c OutputConfig) Build(set component.TelemetrySettings) (OutputOperator, error) {
	_ = "STUB: not implemented"
	return *new(OutputOperator), nil
}

// OutputOperator provides a basic implementation of an output operator.
type OutputOperator struct {
	BasicOperator
	// prevent unkeyed literal initialization
	_ struct{}
}

// CanProcess will always return true for an output operator.
func (*OutputOperator) CanProcess() bool {
	_ = "STUB: not implemented"

	// CanOutput will always return false for an output operator.
	return false
}

func (*OutputOperator) CanOutput() bool {
	_ = "STUB: not implemented"

	// Outputs will always return an empty array for an output operator.
	return false
}

func (*OutputOperator) Outputs() []operator.Operator { _ = "STUB: not implemented"; return nil }

// GetOutputIDs will always return an empty array for an output ID.
func (*OutputOperator) GetOutputIDs() []string {
	_ = "STUB: not implemented"

	// SetOutputs will return an error if called.
	return nil
}

func (*OutputOperator) SetOutputs(_ []operator.Operator) error {
	_ = "STUB: not implemented"
	return nil
}

// SetOutputIDs will return nothing and does nothing.
func (*OutputOperator) SetOutputIDs([]string) { _ = "STUB: not implemented"; return }
