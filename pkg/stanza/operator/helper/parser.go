// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"context"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// NewParserConfig creates a new parser config with default values
func NewParserConfig(operatorID, operatorType string) ParserConfig {
	_ = "STUB: not implemented"
	return *new(ParserConfig)
}

// ParserConfig provides the basic implementation of a parser config.
type ParserConfig struct {
	TransformerConfig `mapstructure:",squash"`
	ParseFrom         entry.Field         `mapstructure:"parse_from"`
	ParseTo           entry.RootableField `mapstructure:"parse_to"`
	BodyField         *entry.Field        `mapstructure:"body"`
	TimeParser        *TimeParser         `mapstructure:"timestamp,omitempty"`
	SeverityConfig    *SeverityConfig     `mapstructure:"severity,omitempty"`
	TraceParser       *TraceParser        `mapstructure:"trace,omitempty"`
	ScopeNameParser   *ScopeNameParser    `mapstructure:"scope_name,omitempty"`
}

// Build will build a parser operator.
func (c ParserConfig) Build(set component.TelemetrySettings) (ParserOperator, error) {
	_ = "STUB: not implemented"
	return *new(ParserOperator), nil
}

// ParserOperator provides a basic implementation of a parser operator.
type ParserOperator struct {
	TransformerOperator
	ParseFrom       entry.Field
	ParseTo         entry.Field
	BodyField       *entry.Field
	TimeParser      *TimeParser
	SeverityParser  *SeverityParser
	TraceParser     *TraceParser
	ScopeNameParser *ScopeNameParser
}

func (p *ParserOperator) ProcessBatchWith(ctx context.Context, entries []*entry.Entry, parse ParseFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *ParserOperator) ProcessBatchWithCallback(ctx context.Context, entries []*entry.Entry, parse ParseFunction, cb func(*entry.Entry) error) error {
	_ = "STUB: not implemented"
	return nil
}

// ProcessWith will run ParseWith on the entry, then forward the entry on to the next operators.
func (p *ParserOperator) ProcessWith(ctx context.Context, entry *entry.Entry, parse ParseFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *ParserOperator) ProcessWithCallback(ctx context.Context, entry *entry.Entry, parse ParseFunction, cb func(*entry.Entry) error) error {
	_ = "STUB: not implemented"
	// Short circuit if the "if" condition does not match
	return nil
}

// ParseWith will process an entry's field with a parser function.
func (p *ParserOperator) ParseWith(ctx context.Context, entry *entry.Entry, parse ParseFunction, write WriteFunction) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle parsing errors after attempting to parse all

// ParseFunction is function that parses a raw value.
type ParseFunction = func(any) (any, error)
