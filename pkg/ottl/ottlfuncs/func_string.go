// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type StringArguments[K any] struct {
	Target ottl.StringLikeGetter[K]
}

func NewStringFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createStringFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stringFunc[K any](target ottl.StringLikeGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
