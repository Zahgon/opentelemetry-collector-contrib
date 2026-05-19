// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"
import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type KeepMatchingKeysArguments[K any] struct {
	Target  ottl.PMapGetSetter[K]
	Pattern ottl.StringGetter[K]
}

func NewKeepMatchingKeysFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createKeepMatchingKeysFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func keepMatchingKeys[K any](target ottl.PMapGetSetter[K], pattern ottl.StringGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
