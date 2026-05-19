// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remove // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/remove"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer is an operator that deletes a field
type Transformer struct {
	helper.TransformerOperator
	Field rootableField
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will process an entry with a remove transformation.
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Transform will apply the remove operation to an entry
func (t *Transformer) Transform(entry *entry.Entry) error { _ = "STUB: not implemented"; return nil }
