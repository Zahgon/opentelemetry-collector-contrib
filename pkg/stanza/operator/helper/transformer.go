// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"context"

	"github.com/expr-lang/expr/vm"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// NewTransformerConfig creates a new transformer config with default values
func NewTransformerConfig(operatorID, operatorType string) TransformerConfig {
	_ = "STUB: not implemented"
	return *new(TransformerConfig)
}

// TransformerConfig provides a basic implementation of a transformer config.
type TransformerConfig struct {
	WriterConfig `mapstructure:",squash"`
	OnError      string `mapstructure:"on_error"`
	IfExpr       string `mapstructure:"if"`
}

// Build will build a transformer operator.
func (c TransformerConfig) Build(set component.TelemetrySettings) (TransformerOperator, error) {
	_ = "STUB: not implemented"
	return *new(TransformerOperator), nil
}

// TransformerOperator provides a basic implementation of a transformer operator.
type TransformerOperator struct {
	WriterOperator
	OnError string
	IfExpr  *vm.Program
}

// CanProcess will always return true for a transformer operator.
func (*TransformerOperator) CanProcess() bool { _ = "STUB: not implemented"; return false }

func (*TransformerOperator) ProcessBatchWith(ctx context.Context, entries []*entry.Entry, process ProcessFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TransformerOperator) ProcessBatchWithTransform(ctx context.Context, entries []*entry.Entry, transform TransformFunction) error {
	_ = "STUB: not implemented"
	return nil
}

// Write the entry without transforming

// Only append error if not in quiet mode

// Write the transformed entry

// ProcessWith will process an entry with a transform function.
func (t *TransformerOperator) ProcessWith(ctx context.Context, entry *entry.Entry, transform TransformFunction) error {
	_ = "STUB: not implemented"
	// Short circuit if the "if" condition does not match
	return nil
}

// Return nil for quiet modes to prevent error from bubbling up

// HandleEntryError will handle an entry error using the on_error strategy.
func (t *TransformerOperator) HandleEntryError(ctx context.Context, entry *entry.Entry, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TransformerOperator) HandleEntryErrorWithWrite(ctx context.Context, entry *entry.Entry, err error, write WriteFunction) error {
	_ = "STUB: not implemented"
	return nil
}

// No need to construct the zap attributes if logging not enabled at debug level.

// isQuietMode returns true if the operator is configured to use quiet mode
func (t *TransformerOperator) isQuietMode() bool { _ = "STUB: not implemented"; return false }

func (t *TransformerOperator) Skip(_ context.Context, entry *entry.Entry) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func zapAttributes(entry *entry.Entry, action string, err error) []zap.Field {
	_ = "STUB: not implemented"
	return nil
}

// ProcessFunction is a function that processes an entry.
type ProcessFunction = func(context.Context, *entry.Entry) error

// TransformFunction is function that transforms an entry.
type TransformFunction = func(*entry.Entry) error

// SendOnError specifies an on_error mode for sending entries after an error.
const SendOnError = "send"

// SendOnErrorQuiet specifies an on_error mode for sending entries after an error but without logging on error level
const SendOnErrorQuiet = "send_quiet"

// DropOnError specifies an on_error mode for dropping entries after an error.
const DropOnError = "drop"

// DropOnErrorQuiet specifies an on_error mode for dropping entries after an error but without logging on error level
const DropOnErrorQuiet = "drop_quiet"
