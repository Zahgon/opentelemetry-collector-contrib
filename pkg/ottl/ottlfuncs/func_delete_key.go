// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type DeleteKeyArguments[K any] struct {
	Target ottl.PMapGetSetter[K]
	Key    ottl.StringGetter[K]
}

func NewDeleteKeyFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createDeleteKeyFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteKey[K any](target ottl.PMapGetSetter[K], key ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
