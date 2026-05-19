// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package uri // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/uri"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Parser is an operator that parses a uri.
type Parser struct {
	helper.ParserOperator
}

func (p *Parser) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will parse an entry.
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// parse will parse a uri from a field and attach it to an entry.
func (*Parser) parse(value any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
