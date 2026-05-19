// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package assignkeys // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/assignkeys"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer transforms a list in the entry field into a map. Each value is assigned a key from configuration keys
type Transformer struct {
	helper.TransformerOperator
	Field entry.Field
	Keys  []string
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will process an entry with AssignKeys transformation.
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Transform will apply AssignKeys to an entry
func (t *Transformer) Transform(entry *entry.Entry) error { _ = "STUB: not implemented"; return nil }

// The field doesn't exist, so ignore it

func (*Transformer) AssignKeys(keys []string, values []any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
