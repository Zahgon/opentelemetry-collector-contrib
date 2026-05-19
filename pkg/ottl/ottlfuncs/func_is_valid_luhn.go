// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"
import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IsValidLuhnArguments[K any] struct {
	Target ottl.StringLikeGetter[K]
}

func NewIsValidLuhnFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIsValidLuhnFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isValidLuhnFunc[K any](target ottl.StringLikeGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// first trim all spaces

// return false if the value is an empty string

// extract the check digit (the right most digit)

// double the digit

// subtract 9 if the number is greater than 9

// calculate the check sum
