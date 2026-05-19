// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type AppendArguments[K any] struct {
	Target ottl.GetSetter[K]
	Value  ottl.Optional[ottl.Getter[K]]
	Values ottl.Optional[[]ottl.Getter[K]]
}

func NewAppendFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createAppendFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendTo[K any](target ottl.GetSetter[K], value ottl.Optional[ottl.Getter[K]], values ottl.Optional[[]ottl.Getter[K]]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// init res with target values

// retype []any to Slice, having []any sometimes misbehaves and nils pcommon.Value

func appendMultiple[K any](target []any, values []K) []any { _ = "STUB: not implemented"; return nil }
