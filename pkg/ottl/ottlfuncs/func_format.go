// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type FormatArguments[K any] struct {
	Format string
	Vals   []ottl.Getter[K]
}

func NewFormatFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createFormatFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func format[K any](formatString string, vals []ottl.Getter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
