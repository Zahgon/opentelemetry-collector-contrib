// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	parseCSVModeStrict       = "strict"
	parseCSVModeLazyQuotes   = "lazyQuotes"
	parseCSVModeIgnoreQuotes = "ignoreQuotes"
)

const (
	parseCSVDefaultDelimiter = ','
	parseCSVDefaultMode      = parseCSVModeStrict
)

type ParseCSVArguments[K any] struct {
	Target          ottl.StringGetter[K]
	Header          ottl.StringGetter[K]
	Delimiter       ottl.Optional[string]
	HeaderDelimiter ottl.Optional[string]
	Mode            ottl.Optional[string]
}

func (p ParseCSVArguments[K]) validate() error { _ = "STUB: not implemented"; return nil }

func NewParseCSVFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createParseCSVFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// headerDelimiter defaults to the chosen delimter,
// since in most cases headerDelimiter == delmiter.

type parseCSVRowFunc func(row string, delimiter rune) ([]string, error)

func parseCSV[K any](target, header ottl.StringGetter[K], delimiter rune, headerDelimiter string, parseRow parseCSVRowFunc) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

func parseCSVRow(lazyQuotes bool) parseCSVRowFunc {
	_ = "STUB: not implemented"
	return *new(parseCSVRowFunc)
}

func parseCSVRowIgnoreQuotes() parseCSVRowFunc {
	_ = "STUB: not implemented"
	return *new(parseCSVRowFunc)
}
