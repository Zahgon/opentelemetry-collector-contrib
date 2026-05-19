// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type DeleteIndexArguments[K any] struct {
	Target     ottl.PSliceGetSetter[K]
	StartIndex ottl.IntGetter[K]
	EndIndex   ottl.Optional[ottl.IntGetter[K]]
}

func NewDeleteIndexFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createDeleteIndexFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteIndexFrom[K any](target ottl.PSliceGetSetter[K], startIndexGetter ottl.IntGetter[K], endIndexGetter ottl.Optional[ottl.IntGetter[K]]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// If deleting all elements, return an empty slice without looping

// No elements to delete

func validateBounds(startIndex, endIndex, sliceLen int64) error {
	_ = "STUB: not implemented"
	return nil
}
