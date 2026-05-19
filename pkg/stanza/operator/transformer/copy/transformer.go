// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package copy // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/copy"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer copies a value from one field and creates a new field with that value
type Transformer struct {
	helper.TransformerOperator
	From entry.Field
	To   entry.Field
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will process an entry with a copy transformation.
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Transform will apply the copy operation to an entry
func (t *Transformer) Transform(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }
