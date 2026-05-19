// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IsDoubleArguments[K any] struct {
	Target ottl.FloatGetter[K]
}

func NewIsDoubleFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIsDoubleFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errorlint
func isDouble[K any](target ottl.FloatGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// Use type assertion because we don't want to check wrapped errors
