// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type URLArguments[K any] struct {
	URI ottl.StringGetter[K]
}

func NewURLFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createURIFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//revive:disable-line:var-naming

func url[K any](uriSource ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented" //revive:disable-line:var-naming
	return nil
}
