// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IsStringArguments[K any] struct {
	Target ottl.StringGetter[K]
}

func NewIsStringFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIsStringFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errorlint
func isString[K any](target ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// Use type assertion because we don't want to check wrapped errors
