// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type LogArguments[K any] struct {
	Target ottl.FloatLikeGetter[K]
}

func NewLogFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createLogFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func logFunc[K any](target ottl.FloatLikeGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
