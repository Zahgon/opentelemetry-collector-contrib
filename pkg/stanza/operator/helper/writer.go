// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"context"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// NewWriterConfig creates a new writer config
func NewWriterConfig(operatorID, operatorType string) WriterConfig {
	_ = "STUB: not implemented"
	return *new(WriterConfig)
}

// WriterConfig is the configuration of a writer operator.
type WriterConfig struct {
	BasicConfig `mapstructure:",squash"`
	OutputIDs   []string `mapstructure:"output"`
}

// Build will build a writer operator from the config.
func (c WriterConfig) Build(set component.TelemetrySettings) (WriterOperator, error) {
	_ = "STUB: not implemented"
	return *new(WriterOperator), nil
}

// WriterOperator is an operator that can write to other operators.
type WriterOperator struct {
	BasicOperator
	OutputIDs       []string
	OutputOperators []operator.Operator
}

// WriteBatch writes a batch of entries to the outputs of the operator.
// A batch is a collection of entries that are sent in one go.
func (w *WriterOperator) WriteBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Write will write an entry to the outputs of the operator.
func (w *WriterOperator) Write(ctx context.Context, e *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// CanOutput always returns true for a writer operator.
func (*WriterOperator) CanOutput() bool {
	_ = "STUB: not implemented"

	// Outputs returns the outputs of the writer operator.
	return false
}

func (w *WriterOperator) Outputs() []operator.Operator { _ = "STUB: not implemented"; return nil }

// GetOutputIDs returns the output IDs of the writer operator.
func (w *WriterOperator) GetOutputIDs() []string {
	_ = "STUB: not implemented"

	// SetOutputs will set the outputs of the operator.
	return nil
}

func (w *WriterOperator) SetOutputs(operators []operator.Operator) error {
	_ = "STUB: not implemented"
	return nil
}

// SetOutputIDs will set the outputs of the operator.
func (w *WriterOperator) SetOutputIDs(opIDs []string) { _ = "STUB: not implemented"; return }

// FindOperator will find an operator matching the supplied id.
func (*WriterOperator) findOperator(operators []operator.Operator, operatorID string) (operator.Operator, bool) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), false
}

type WriteFunction = func(context.Context, *entry.Entry) error
