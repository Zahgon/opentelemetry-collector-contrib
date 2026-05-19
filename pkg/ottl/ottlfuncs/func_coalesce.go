// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type CoalesceArguments[K any] struct {
	Values []ottl.Getter[K]
}

func NewCoalesceFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createCoalesceFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func coalesce[K any](values []ottl.Getter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
