// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IndexArguments[K any] struct {
	Target ottl.Getter[K]
	Value  ottl.Getter[K]
}

func NewIndexFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIndexFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func index[K any](valueComparator ottl.ValueComparator, target, value ottl.Getter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
