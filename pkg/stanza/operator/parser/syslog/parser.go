// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslog // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog"

import (
	"context"
	"regexp"
	"time"

	sl "github.com/leodido/go-syslog/v4"
	"github.com/leodido/go-syslog/v4/nontransparent"
	"github.com/leodido/go-syslog/v4/rfc3164"
	"github.com/leodido/go-syslog/v4/rfc5424"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

var priRegex = regexp.MustCompile(`<\d{1,3}>`)

// parseFunc a parseFunc determines how the raw input is to be parsed into a syslog message
type parseFunc func(input []byte) (sl.Message, error)

// Parser is an operator that parses syslog.
type Parser struct {
	helper.ParserOperator
	protocol                     string
	location                     *time.Location
	enableOctetCounting          bool
	allowSkipPriHeader           bool
	nonTransparentFramingTrailer *string
	maxOctets                    int
}

func (p *Parser) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Determine which callback to use based on entry data

// Process will parse an entry field as syslog.
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	// if pri header is missing and this is an expected behavior then facility and severity values should be skipped.
	return nil
}

// parse will parse a value as syslog.
func (p *Parser) parse(value any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (p *Parser) buildParseFunc() (parseFunc, error) {
	_ = "STUB: not implemented"
	return *new(parseFunc), nil
}

// Octet Counting Parsing RFC6587

// Non-Transparent-Framing Parsing RFC6587

// Raw RFC5424 parsing

func (p *Parser) shouldSkipPriorityValues(value []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// check if entry starts with '<'.
// if not it means that the pre header was missing from the body and hence we should skip it.

// parseRFC3164 will parse an RFC3164 syslog message.
func (p *Parser) parseRFC3164(syslogMessage *rfc3164.SyslogMessage, skipPriHeaderValues bool) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseRFC5424 will parse an RFC5424 syslog message.
func (p *Parser) parseRFC5424(syslogMessage *rfc5424.SyslogMessage, skipPriHeaderValues bool) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// toSafeMap will dereference any pointers on the supplied map.
func (*Parser) toSafeMap(message map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertMap converts map[string]map[string]string to map[string]any
// which is expected by stanza converter
func convertMap(data map[string]map[string]string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func toBytes(value any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var severityMapping = [...]entry.Severity{
	0: entry.Fatal,
	1: entry.Error3,
	2: entry.Error2,
	3: entry.Error,
	4: entry.Warn,
	5: entry.Info2,
	6: entry.Info,
	7: entry.Debug,
}

var severityText = [...]string{
	0: "emerg",
	1: "alert",
	2: "crit",
	3: "err",
	4: "warning",
	5: "notice",
	6: "info",
	7: "debug",
}

var severityField = entry.NewAttributeField("severity")

func cleanupTimestamp(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }

func postprocessWithoutPriHeader(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }

func postprocess(e *entry.Entry) error { _ = "STUB: not implemented"; return nil }

func newOctetCountingParseFunc(maxOctets int) parseFunc {
	_ = "STUB: not implemented"
	return *new(parseFunc)
}

func newNonTransparentFramingParseFunc(trailerType nontransparent.TrailerType) parseFunc {
	_ = "STUB: not implemented"
	return *new(parseFunc)
}

// isQuietMode returns true if the operator is configured to use quiet mode
func (p *Parser) isQuietMode() bool { _ = "STUB: not implemented"; return false }
