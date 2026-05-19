// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IntArguments[K any] struct {
	Target ottl.IntLikeGetter[K]
}

func NewIntFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIntFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func intFunc[K any](target ottl.IntLikeGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
