// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"context"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// NewInputConfig creates a new input config with default values.
func NewInputConfig(operatorID, operatorType string) InputConfig {
	_ = "STUB: not implemented"
	return *new(InputConfig)
}

// InputConfig provides a basic implementation of an input operator config.
type InputConfig struct {
	AttributerConfig `mapstructure:",squash"`
	IdentifierConfig `mapstructure:",squash"`
	WriterConfig     `mapstructure:",squash"`
}

// Build will build a base producer.
func (c InputConfig) Build(set component.TelemetrySettings) (InputOperator, error) {
	_ = "STUB: not implemented"
	return *new(InputOperator), nil
}

// InputOperator provides a basic implementation of an input operator.
type InputOperator struct {
	Attributer
	Identifier
	WriterOperator
}

// NewEntry will create a new entry using the `attributes`, and `resource` configuration.
func (i *InputOperator) NewEntry(value any) (*entry.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CanProcess will always return false for an input operator.
func (*InputOperator) CanProcess() bool {
	_ = "STUB: not implemented"

	// ProcessBatch will always return an error if called.
	return false
}

func (i *InputOperator) ProcessBatch(_ context.Context, _ []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will always return an error if called.
func (i *InputOperator) Process(_ context.Context, _ *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}
