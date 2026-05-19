// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expr // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"

import (
	"context"
)

// BoolExpr is an interface that allows matching a context K against a configuration of a match.
type BoolExpr[K any] interface {
	Eval(ctx context.Context, tCtx K) (bool, error)
}

type notMatcher[K any] struct {
	matcher BoolExpr[K]
}

func (nm notMatcher[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func Not[K any](matcher BoolExpr[K]) BoolExpr[K] { _ = "STUB: not implemented"; return nil }

type alwaysTrueMatcher[K any] struct{}

func (alwaysTrueMatcher[K]) Eval(context.Context, K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func AlwaysTrue[K any]() BoolExpr[K] { _ = "STUB: not implemented"; return nil }

type orMatcher[K any] struct {
	matchers []BoolExpr[K]
}

func (om orMatcher[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func Or[K any](matchers ...BoolExpr[K]) BoolExpr[K] { _ = "STUB: not implemented"; return nil }
