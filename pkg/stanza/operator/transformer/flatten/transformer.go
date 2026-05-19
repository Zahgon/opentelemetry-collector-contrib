// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package flatten // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/flatten"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer flattens an object in the entry field
type Transformer[T interface {
	entry.BodyField | entry.ResourceField | entry.AttributeField
	entry.FieldInterface
	Parent() T
	Child(string) T
}] struct {
	helper.TransformerOperator
	Field T
}

func (t *Transformer[T]) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will process an entry with a flatten transformation.
func (t *Transformer[T]) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Transform will apply the flatten operation to an entry
func (t *Transformer[T]) Transform(entry *entry.Entry) error { _ = "STUB: not implemented"; return nil }

// The field doesn't exist, so ignore it

// The field we were asked to flatten was not a map, so put it back
