// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type NanosecondsArguments[K any] struct {
	Duration ottl.DurationGetter[K]
}

func NewNanosecondsFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createNanosecondsFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Nanoseconds[K any](duration ottl.DurationGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
