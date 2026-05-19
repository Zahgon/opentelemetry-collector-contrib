// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ConcatArguments[K any] struct {
	Vals      []ottl.StringLikeGetter[K]
	Delimiter ottl.StringGetter[K]
}

func NewConcatFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createConcatFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func concat[K any](vals []ottl.StringLikeGetter[K], delimiter ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
