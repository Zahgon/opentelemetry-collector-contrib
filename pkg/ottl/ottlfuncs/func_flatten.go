// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type FlattenArguments[K any] struct {
	Target           ottl.PMapGetSetter[K]
	Prefix           ottl.Optional[string]
	Depth            ottl.Optional[int64]
	ResolveConflicts ottl.Optional[bool]
}

type flattenData struct {
	result          pcommon.Map
	existingKeys    map[string]int
	resolveConflict bool
	maxDepth        int64
}

func NewFlattenFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createFlattenFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func flatten[K any](target ottl.PMapGetSetter[K], p ottl.Optional[string], d ottl.Optional[int64], c ottl.Optional[bool]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initFlattenData(resolveConflict bool, maxDepth int64) *flattenData {
	_ = "STUB: not implemented"
	return nil
}

func (f *flattenData) flattenMap(m pcommon.Map, prefix string, currentDepth int64) {
	_ = "STUB: not implemented"
	return
}

func (f *flattenData) flattenSlice(s pcommon.Slice, prefix string, currentDepth int64) {
	_ = "STUB: not implemented"
	return
}

func (f *flattenData) flattenValue(k string, v pcommon.Value, currentDepth int64, prefix string) {
	_ = "STUB: not implemented"
	return
}

func (f *flattenData) handleConflict(key string, v pcommon.Value) {
	_ = "STUB: not implemented"
	return
}
