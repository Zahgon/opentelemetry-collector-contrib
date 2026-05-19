// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package regexreplace // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/regexreplace"

import (
	"context"
	"regexp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Transformer is an operator that performs a regex-replace on a string field.
type Transformer struct {
	helper.TransformerOperator
	field       entry.Field
	regexp      *regexp.Regexp
	replaceWith string
}

func (t *Transformer) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transformer) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transformer) replace(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }
