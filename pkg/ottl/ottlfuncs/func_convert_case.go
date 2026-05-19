// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ConvertCaseArguments[K any] struct {
	Target ottl.StringGetter[K]
	ToCase string
}

func NewConvertCaseFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createConvertCaseFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertCase[K any](target ottl.StringGetter[K], toCase string) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert string to lowercase (SOME_NAME -> some_name)

// Convert string to uppercase (some_name -> SOME_NAME)

// Convert string to snake case (someName -> some_name)

// Convert string to camel case (some_name -> SomeName)
