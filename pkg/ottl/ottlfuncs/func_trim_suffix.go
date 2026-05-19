// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type TrimSuffixArguments[K any] struct {
	Target ottl.StringGetter[K]
	Suffix ottl.StringGetter[K]
}

func NewTrimSuffixFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createTrimSuffixFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func trimSuffix[K any](target, prefix ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
