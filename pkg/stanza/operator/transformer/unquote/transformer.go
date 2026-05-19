// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unquote // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/unquote"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer is an operator that unquotes a string.
type Transformer struct {
	helper.TransformerOperator
	field entry.Field
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will unquote a string
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// unquote will unquote a string
func (t *Transformer) unquote(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }
