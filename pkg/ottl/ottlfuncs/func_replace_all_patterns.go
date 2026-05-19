// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	modeKey   = "key"
	modeValue = "value"
)

type ReplaceAllPatternsArguments[K any] struct {
	Target            ottl.PMapGetSetter[K]
	Mode              string
	RegexPattern      ottl.StringGetter[K]
	Replacement       ottl.StringGetter[K]
	Function          ottl.Optional[ottl.FunctionGetter[K]]
	ReplacementFormat ottl.Optional[ottl.StringGetter[K]]
}

func NewReplaceAllPatternsFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createReplaceAllPatternsFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replaceAllPatterns[K any](target ottl.PMapGetSetter[K], mode string, regexPattern, replacement ottl.StringGetter[K], fn ottl.Optional[ottl.FunctionGetter[K]], replacementFormat ottl.Optional[ottl.StringGetter[K]]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Because we are changing the keys we cannot do in-place update, but we can move values to the
// updated map and then move back the updated map to the initial map to avoid a copy in the target.Set,
// because the pcommon.Map.CopyTo will not do a copy if it is the same object in this case val.
