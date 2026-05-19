// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type SetArguments[K any] struct {
	Target ottl.Setter[K]
	Value  ottl.Getter[K]
}

func NewSetFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createSetFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func set[K any](target ottl.Setter[K], value ottl.Getter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// No fields currently support `null` as a valid type.
