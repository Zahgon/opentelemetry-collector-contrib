// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type TrimArguments[K any] struct {
	Target      ottl.StringGetter[K]
	Replacement ottl.Optional[string]
}

func NewTrimFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createTrimFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func trim[K any](target ottl.StringGetter[K], replacement ottl.Optional[string]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
