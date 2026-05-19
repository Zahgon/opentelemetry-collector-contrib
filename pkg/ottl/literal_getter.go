// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"context"
)

// literalGetter is an optional interface that allows Getter implementations to indicate
// if they support literal values retrieval.
type literalGetter interface {
	isLiteral()
}

type literal[K any, T any] struct {
	value T
}

func newLiteral[K, T any](value T) *literal[K, T] { _ = "STUB: not implemented"; return nil }

func (l *literal[K, T]) Get(context.Context, K) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (*literal[K, T]) isLiteral() { _ = "STUB: not implemented"; return }

func isLiteralGetter[K, V any](getter typedGetter[K, V]) bool {
	_ = "STUB: not implemented"
	return false
}

// GetLiteralValue retrieves the literal value from the given getter.
// If the getter is not a literal getter, or if the value it's currently holding is not a
// literal value, it returns the zero value of V and false.
func GetLiteralValue[K, V any](getter typedGetter[K, V]) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}
