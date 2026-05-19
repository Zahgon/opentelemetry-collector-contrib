// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ReplaceMatchArguments[K any] struct {
	Target            ottl.GetSetter[K]
	Pattern           ottl.StringGetter[K]
	Replacement       ottl.StringGetter[K]
	Function          ottl.Optional[ottl.FunctionGetter[K]]
	ReplacementFormat ottl.Optional[ottl.StringGetter[K]]
}

type replaceMatchFuncArgs[K any] struct {
	Input ottl.StringGetter[K]
}

func NewReplaceMatchFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createReplaceMatchFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replaceMatch[K any](target ottl.GetSetter[K], pattern, replacement ottl.StringGetter[K], fn ottl.Optional[ottl.FunctionGetter[K]], replacementFormat ottl.Optional[ottl.StringGetter[K]]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
