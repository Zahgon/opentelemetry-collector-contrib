// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IsListArguments[K any] struct {
	Target ottl.Getter[K]
}

func NewIsListFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIsListFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isList[K any](target ottl.Getter[K]) ottl.ExprFunc[K] { _ = "STUB: not implemented"; return nil }
