// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type UnixArguments[K any] struct {
	Seconds     ottl.IntGetter[K]
	Nanoseconds ottl.Optional[ottl.IntGetter[K]]
}

func NewUnixFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createUnixFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Unix[K any](seconds ottl.IntGetter[K], nanoseconds ottl.Optional[ottl.IntGetter[K]]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
