// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package jsonarray // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/jsonarray"

import (
	"context"

	"github.com/valyala/fastjson"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Parser is an operator that parses json array in an entry.
type Parser struct {
	helper.ParserOperator
	parse parseFunc
}

type parseFunc func(any) (any, error)

func (p *Parser) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will parse an entry for json array.
func (p *Parser) Process(ctx context.Context, e *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func generateParseToArrayFunc(pool *fastjson.ParserPool) parseFunc {
	_ = "STUB: not implemented"
	return *new(parseFunc)
}

// a is a []*Value slice

// Nested objects handled as a string since this parser doesn't support nested headers

func generateParseToMapFunc(pool *fastjson.ParserPool, header []string) parseFunc {
	_ = "STUB: not implemented"
	return *new(parseFunc)
}

// a is a []*Value slice

// Nested objects handled as a string since this parser doesn't support nested headers

// valueAsString interprets the given value as a string.
func valueAsString(value any) (string, error) { _ = "STUB: not implemented"; return "", nil }
