// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (

	// #nosec

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type SHA1Arguments[K any] struct {
	Target ottl.StringGetter[K]
}

func NewSHA1Factory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createSHA1Function[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SHA1HashString[K any](target ottl.StringGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec
