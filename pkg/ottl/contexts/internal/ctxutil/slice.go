// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxutil"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"golang.org/x/exp/constraints"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	typeNotIndexableError = "type %T does not support indexing"
)

var (
	errMissingSetKey = errors.New("cannot set slice value without key")
	errMissingGetKey = errors.New("cannot get slice value without key")
)

func GetSliceIndexFromKeys[K any](ctx context.Context, tCtx K, sliceLen int, keys []ottl.Key[K]) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetSliceValue[K any](ctx context.Context, tCtx K, s pcommon.Slice, keys []ottl.Key[K]) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func SetSliceValue[K any](ctx context.Context, tCtx K, s pcommon.Slice, keys []ottl.Key[K], val any) error {
	_ = "STUB: not implemented"
	return nil
}

// CommonTypedSlice is an interface for typed pdata slices, such as pcommon.StringSlice,
// pcommon.Int64Slice, pcommon.Int32Slice, etc.
type CommonTypedSlice[T any] interface {
	At(int) T
	Len() int
	FromRaw(val []T)
	SetAt(int, T)
	AsRaw() []T
}

// GetCommonTypedSliceValue is like GetSliceValue, but for retrieving a value from a pdata
// typed slice. [V] is the type of the slice elements. If no keys are provided, it returns
// an error. This function does not support slice elements indexing.
func GetCommonTypedSliceValue[K, V any](ctx context.Context, tCtx K, s CommonTypedSlice[V], keys []ottl.Key[K]) (V, error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

// SetCommonTypedSliceValue sets the value of a pdata typed slice element. [V] is the type
// of the slice elements. The any-wrapped value type must be [V], otherwise an error is
// returned. This function does not support slice elements indexing.
func SetCommonTypedSliceValue[K, V any](ctx context.Context, tCtx K, s CommonTypedSlice[V], keys []ottl.Key[K], val any) error {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonTypedSliceValues sets the value of a pdata typed slice. It does handle all
// different input types OTTL generate, such as []V, []any, or a pcommon.Slice.
// If the value is a slice of [any] or pcommon.Slice, and it has an element that the type
// is not [V], an error is returned.
func SetCommonTypedSliceValues[V any](s CommonTypedSlice[V], val any) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCommonIntSliceValues converts a pdata typed slice of [constraints.Integer] into
// []int64, which is the standard OTTL type for integer slices.
func GetCommonIntSliceValues[V constraints.Integer](val CommonTypedSlice[V]) []int64 {
	_ = "STUB: not implemented"
	return nil
}

// GetCommonIntSliceValue is like GetCommonTypedSliceValue, but for integer pdata typed
// slices.
func GetCommonIntSliceValue[K any, V constraints.Integer](ctx context.Context, tCtx K, s CommonTypedSlice[V], keys []ottl.Key[K]) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SetCommonIntSliceValue is like SetCommonTypedSliceValue, but for integer pdata typed
// slice element values. [V] is the type of the slice elements.
// The any-wrapped value type must be and integer convertible to [V], otherwise an error
// is returned. This function does not support slice elements indexing.
func SetCommonIntSliceValue[K any, V constraints.Integer](ctx context.Context, tCtx K, s CommonTypedSlice[V], keys []ottl.Key[K], val any) error {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonIntSliceValues is like SetCommonTypedSliceValues, but for integer pdata typed
// slices. [V] is the type of the slice elements. The value must be []any, []int64, []T,
// or a pcommon.Slice which elements are type inferable to int64, otherwise an error is
// returned.
func SetCommonIntSliceValues[V constraints.Integer](s CommonTypedSlice[V], val any) error {
	_ = "STUB: not implemented"
	return nil
}
