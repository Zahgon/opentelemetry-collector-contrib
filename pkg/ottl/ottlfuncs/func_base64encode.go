// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type Base64EncodeArguments[K any] struct {
	Target  ottl.StringGetter[K]
	Variant ottl.Optional[ottl.StringGetter[K]]
}

func NewBase64EncodeFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createBase64EncodeFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func base64Encode[K any](target ottl.StringGetter[K], variant ottl.Optional[ottl.StringGetter[K]]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
