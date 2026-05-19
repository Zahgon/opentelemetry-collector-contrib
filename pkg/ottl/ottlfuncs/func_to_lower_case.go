// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ToLowerCaseArguments[K any] struct {
	Target ottl.StringGetter[K]
}

func NewToLowerCaseFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createToLowerCaseFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toLowerCase[K any](target ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
