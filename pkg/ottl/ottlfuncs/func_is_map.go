// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IsMapArguments[K any] struct {
	Target ottl.PMapGetter[K]
}

func NewIsMapFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIsMapFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errorlint
func isMap[K any](target ottl.PMapGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// Use type assertion because we don't want to check wrapped errors
