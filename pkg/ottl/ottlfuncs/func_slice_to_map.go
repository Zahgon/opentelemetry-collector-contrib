// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"
import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type SliceToMapArguments[K any] struct {
	Target    ottl.PSliceGetter[K]
	KeyPath   ottl.Optional[[]string]
	ValuePath ottl.Optional[[]string]
}

func NewSliceToMapFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func sliceToMapFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSliceToMapFunc[K any](target ottl.PSliceGetter[K], keyPath, valuePath ottl.Optional[[]string]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

func sliceToMap(v pcommon.Slice, keyPath, valuePath ottl.Optional[[]string]) (pcommon.Map, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Map), nil
}

// If value_path is not set, value is the whole element

func extractValue(m pcommon.Map, path []string) (pcommon.Value, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), nil
}
