// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	sortAsc  = "asc"
	sortDesc = "desc"
)

type SortArguments[K any] struct {
	Target ottl.Getter[K]
	Order  ottl.Optional[string]
}

func NewSortFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createSortFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sort[K any](target ottl.Getter[K], order string) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// handle Sort([1,2,3])

// sortSlice sorts a pcommon.Slice based on the specified order.
// It gets the common type for all elements in the slice and converts all elements to this common type, creating a new copy
// Parameters:
//   - slice: The pcommon.Slice to be sorted
//   - order: The sort order. "asc" for ascending, "desc" for descending
//
// Returns:
//   - A sorted slice as []any or the original pcommon.Slice
//   - An error if an unsupported type is encountered
func sortSlice(slice pcommon.Slice, order string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type targetType interface {
	~int64 | ~float64 | ~string
}

// findCommonValueType determines the most appropriate common type for all elements in a pcommon.Slice.
// It returns two values:
//   - A pcommon.ValueType representing the desired common type for all elements.
//     Mixed Numeric types return ValueTypeDouble. Integer type returns ValueTypeInt. Double type returns ValueTypeDouble.
//     String, Bool, Empty and mixed of the mentioned types return ValueTypeStr, as they require string conversion for comparison.
//   - A boolean indicating whether a common type could be determined (true) or not (false).
//     returns false for ValueTypeMap, ValueTypeSlice and ValueTypeBytes. They are unsupported types for sort.
func findCommonValueType(slice pcommon.Slice) (pcommon.ValueType, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.ValueType), false
}

func makeCopy[T targetType](src []T) []T { _ = "STUB: not implemented"; return nil }

func sortTypedSlice[T targetType](arr []T, order string) []T { _ = "STUB: not implemented"; return nil }

type convertedValue[T targetType] struct {
	value         T
	originalValue any
}

func makeConvertedCopy[T targetType](slice pcommon.Slice, converter func(idx int) T) []convertedValue[T] {
	_ = "STUB: not implemented"
	return nil
}

func sortConvertedSlice[T targetType](cvs []convertedValue[T], order string) []any {
	_ = "STUB: not implemented"
	return nil
}
