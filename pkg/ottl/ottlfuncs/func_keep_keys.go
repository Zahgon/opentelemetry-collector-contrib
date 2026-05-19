// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type KeepKeysArguments[K any] struct {
	Target ottl.PMapGetSetter[K]
	Keys   []ottl.StringGetter[K]
}

func NewKeepKeysFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createKeepKeysFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func keepKeys[K any](target ottl.PMapGetSetter[K], keys []ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	// Check if all keys are literals and pre-build the key set if so
	return nil
}

// Use pre-built key set for literal keys

// Build key set at runtime for dynamic keys
