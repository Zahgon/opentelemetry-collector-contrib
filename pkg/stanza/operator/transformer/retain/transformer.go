// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package retain // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/retain"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer keeps the given fields and deletes the rest.
type Transformer struct {
	helper.TransformerOperator
	Fields             []entry.Field
	AllBodyFields      bool
	AllAttributeFields bool
	AllResourceFields  bool
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will process an entry with a retain transformation.
func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Transform will apply the retain operation to an entry
func (t *Transformer) Transform(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }

// The entry's Resource, Attributes & Body are modified.
// All other fields are left untouched (Ex: Timestamp, TraceID, ..)
