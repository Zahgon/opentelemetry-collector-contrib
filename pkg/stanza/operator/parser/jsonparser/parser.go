// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jsonparser // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/jsonparser"

import (
	"context"

	"github.com/goccy/go-json"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Parser is an operator that parses JSON.
type Parser struct {
	helper.ParserOperator

	parseInts bool
}

func (p *Parser) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will parse an entry for JSON.
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// parse will parse a value as JSON.
func (p *Parser) parse(value any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// when parseInts is disabled, `int` and `float` will be parsed as `float64`.
// when it is enabled, they will be parsed as `json.Number`, later the parser
// will convert them to `int` or `float64` according to the field type.

func convertNumbers(parsedValue map[string]any) { _ = "STUB: not implemented"; return }

func convertNumbersArray(arr []any) { _ = "STUB: not implemented"; return }

func convertNumber(value json.Number) any { _ = "STUB: not implemented"; return *new(any) }
