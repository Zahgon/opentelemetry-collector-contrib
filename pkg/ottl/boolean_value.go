// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"context"
)

// boolExpr represents a condition in OTTL
type boolExpr[K any] interface {
	Eval(ctx context.Context, tCtx K) (bool, error)

	unexported()
}

type literalBoolExpr[K any] = literalExpr[K, bool]

func newAlwaysTrue[K any]() boolExpr[K] { _ = "STUB: not implemented"; return nil }

func newAlwaysFalse[K any]() boolExpr[K] { _ = "STUB: not implemented"; return nil }

func newNot[K any](expr boolExpr[K]) boolExpr[K] { _ = "STUB: not implemented"; return nil }

type notBoolExpr[K any] struct {
	expr boolExpr[K]
}

func (*notBoolExpr[K]) unexported() {
	_ = "STUB: not implemented"

	// Eval evaluates an OTTL condition
	return
}

func (e *notBoolExpr[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// newAndExprs a boolExpr that returns a short-circuited result of ANDing
// boolExpressionEvaluator funcs
func newAndExprs[K any](exprs []boolExpr[K]) boolExpr[K] { _ = "STUB: not implemented"; return nil }

// If any literal evaluates to false, we can simply return false.
// If a literal evaluates to true, it won't affect the expression's outcome, so we can skip evaluating it again.

// All expressions evaluated to true, just return literal true.

// One expression left, no need to wrap in "andExprs".

type andExprs[K any] struct {
	exprs []boolExpr[K]
}

func (*andExprs[K]) unexported() {
	_ = "STUB: not implemented"

	// Eval evaluates an OTTL condition
	return
}

func (e *andExprs[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// newOrExprs a boolExpr that returns a short-circuited result of ORing
// boolExpressionEvaluator funcs
func newOrExprs[K any](exprs []boolExpr[K]) boolExpr[K] { _ = "STUB: not implemented"; return nil }

// If any literal evaluates to true, we can simply return true.
// If a literal evaluates to false, it won't affect the expression's outcome, so we can skip evaluating it again.

// All expressions evaluated to false, just return literal false.

// One expression left, no need to wrap in "orExprs".

type orExprs[K any] struct {
	exprs []boolExpr[K]
}

func (*orExprs[K]) unexported() {
	_ = "STUB: not implemented"

	// Eval evaluates an OTTL condition
	return
}

func (e *orExprs[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Parser[K]) newComparisonExpr(comparison *comparison) (boolExpr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The parser ensures that we'll never get an invalid comparison.Op, so we don't have to check that case.

type comparisonExpr[K any] struct {
	left, right Getter[K]
	comparator  ValueComparator
	op          compareOp
}

func (*comparisonExpr[K]) unexported() { _ = "STUB: not implemented"; return }

func (e *comparisonExpr[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Parser[K]) newBoolExpr(expr *booleanExpression) (boolExpr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) newBooleanTermEvaluator(term *term) (boolExpr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) newBooleanValueEvaluator(value *booleanValue) (boolExpr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) newConverterEvaluator(c converter) (boolExpr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newConverterExpr[K any](getter Getter[K]) (boolExpr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type converterExpr[K any] struct {
	getter Getter[K]
}

func (*converterExpr[K]) unexported() { _ = "STUB: not implemented"; return }

func (e *converterExpr[K]) Eval(ctx context.Context, tCtx K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
