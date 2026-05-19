// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package router // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/router"

import (
	"context"

	"github.com/expr-lang/expr/vm"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer is an operator that routes entries based on matching expressions
type Transformer struct {
	helper.BasicOperator
	routes []*Route
}

// Route is a route on a router operator
type Route struct {
	helper.Attributer
	Expression      *vm.Program
	OutputIDs       []string
	OutputOperators []operator.Operator
}

// CanProcess will always return true for a router operator
func (*Transformer) CanProcess() bool { _ = "STUB: not implemented"; return false }

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	// Group entries by the route they match
	return nil
}

// Track which route matched

// we compile the expression with "AsBool", so this should be safe

// Group entries by their matching route

// Process batches for each route

// Process will route incoming entries based on matching expressions
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// we compile the expression with "AsBool", so this should be safe

// CanOutput will always return true for a router operator
func (*Transformer) CanOutput() bool {
	_ = "STUB: not implemented"

	// Outputs will return all connected operators.
	return false
}

func (t *Transformer) Outputs() []operator.Operator { _ = "STUB: not implemented"; return nil }

// GetOutputIDs will return all connected operators.
func (t *Transformer) GetOutputIDs() []string { _ = "STUB: not implemented"; return nil }

// SetOutputs will set the outputs of the router operator.
func (t *Transformer) SetOutputs(operators []operator.Operator) error {
	_ = "STUB: not implemented"
	return nil
}

// SetOutputIDs will do nothing.
func (*Transformer) SetOutputIDs(_ []string) {
	_ = "STUB: not implemented"

	// findOperators will find a subset of operators from a collection.
	return
}

func (t *Transformer) findOperators(operators []operator.Operator, operatorIDs []string) ([]operator.Operator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findOperator will find an operator from a collection.
func (*Transformer) findOperator(operators []operator.Operator, operatorID string) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

func zapAttributes(entry *entry.Entry, err error) []zap.Field {
	_ = "STUB: not implemented"
	return nil
}
