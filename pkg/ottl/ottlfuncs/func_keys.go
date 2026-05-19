// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type KeysArguments[K any] struct {
	Target ottl.PMapGetter[K]
}

func NewKeysFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createKeysFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func keys[K any](target ottl.PMapGetter[K]) ottl.ExprFunc[K] { _ = "STUB: not implemented"; return nil }
