// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IsMatchArguments[K any] struct {
	Target  ottl.StringLikeGetter[K]
	Pattern ottl.StringGetter[K]
}

func NewIsMatchFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIsMatchFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isMatch[K any](target ottl.StringLikeGetter[K], pattern ottl.StringGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
