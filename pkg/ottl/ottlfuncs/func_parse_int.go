// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ParseIntArguments[K any] struct {
	Target ottl.StringGetter[K]
	Base   ottl.IntGetter[K]
}

func NewParseIntFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createParseIntFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseIntFunc[K any](target ottl.StringGetter[K], base ottl.IntGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
