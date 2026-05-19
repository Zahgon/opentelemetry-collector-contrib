// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package regex // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/regex"

import (
	"context"
	"regexp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Parser is an operator that parses regex in an entry.
type Parser struct {
	helper.ParserOperator
	regexp *regexp.Regexp
	cache  cache
}

func (p *Parser) Stop() error { _ = "STUB: not implemented"; return nil }

func (p *Parser) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will parse an entry for regex.
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// parse will parse a value using the supplied regex.
func (p *Parser) parse(value any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (p *Parser) match(value string) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
