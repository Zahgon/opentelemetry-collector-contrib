// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/filter"

import (
	"context"
	"math/big"

	"github.com/expr-lang/expr/vm"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer is an operator that filters entries based on matching expressions
type Transformer struct {
	helper.TransformerOperator
	expression *vm.Program
	dropCutoff *big.Int // [0..1000)
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will drop incoming entries that match the filter expression
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}
