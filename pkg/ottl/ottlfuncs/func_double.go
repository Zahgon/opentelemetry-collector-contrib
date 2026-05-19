// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type DoubleArguments[K any] struct {
	Target ottl.FloatLikeGetter[K]
}

func NewDoubleFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createDoubleFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func doubleFunc[K any](target ottl.FloatLikeGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
