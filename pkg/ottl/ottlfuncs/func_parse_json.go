// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ParseJSONArguments[K any] struct {
	Target ottl.StringGetter[K]
}

func NewParseJSONFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createParseJSONFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseJSON returns a `pcommon.Map` or `pcommon.Slice` struct that is a result of parsing the target string as JSON
// Each JSON type is converted into a `pdata.Value` using the following map:
//
//	JSON boolean -> bool
//	JSON number  -> float64
//	JSON string  -> string
//	JSON null    -> nil
//	JSON arrays  -> pdata.SliceValue
//	JSON objects -> map[string]any
func parseJSON[K any](target ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
