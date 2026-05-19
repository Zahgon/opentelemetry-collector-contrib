// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type HasPrefixArguments[K any] struct {
	Target ottl.StringGetter[K]
	Prefix ottl.StringGetter[K]
}

func NewHasPrefixFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createHasPrefixFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HasPrefix[K any](target, prefix ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
