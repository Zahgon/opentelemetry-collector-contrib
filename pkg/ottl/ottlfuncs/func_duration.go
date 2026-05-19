// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type DurationArguments[K any] struct {
	Duration ottl.StringGetter[K]
}

func NewDurationFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createDurationFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Duration[K any](duration ottl.StringGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
