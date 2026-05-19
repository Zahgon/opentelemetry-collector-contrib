// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ContainsValueArguments[K any] struct {
	Target ottl.PSliceGetter[K]
	Item   ottl.Getter[K]
}

func NewContainsValueFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createContainsValueFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func containsValue[K any](target ottl.PSliceGetter[K], itemGetter ottl.Getter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
