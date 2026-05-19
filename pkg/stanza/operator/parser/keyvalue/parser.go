// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package keyvalue // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/keyvalue"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Parser is an operator that parses key value pairs.
type Parser struct {
	helper.ParserOperator
	delimiter     string
	pairDelimiter string
}

func (p *Parser) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will parse an entry for key value pairs.
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// parse will parse a value as key values.
func (p *Parser) parse(value any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (p *Parser) parser(input, delimiter, pairDelimiter string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
