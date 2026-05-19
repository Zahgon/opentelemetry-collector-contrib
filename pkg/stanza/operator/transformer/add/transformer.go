// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package add // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/add"

import (
	"context"

	"github.com/expr-lang/expr/vm"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer is an operator that adds a string value or an expression value
type Transformer struct {
	helper.TransformerOperator

	Field   entry.Field
	Value   any
	program *vm.Program
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will process an entry with a add transformation.
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Transform will apply the add operations to an entry
func (t *Transformer) Transform(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }

func isExpr(str string) bool { _ = "STUB: not implemented"; return false }
