// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"time"
)

func (p *Parser[K]) evaluateMathExpression(expr *mathExpression) (Getter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) evaluateAddSubTerm(term *addSubTerm) (Getter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) evaluateMathValue(val *mathValue) (Getter[K], error) {
	_ = "STUB: not implemented"
	// we have to handle unary plus/minus here because the lexer cannot
	// differentiate between binary and unary operators
	return nil, nil
}

// If the literal is numeric, fold the sign into the literal

// Non-numeric literals fall back to dynamic negation

// Unary plus: no-op to be explicit about it

// Unary plus: no-op to be explicit about it

func negateGetter[K any](baseGetter Getter[K]) Getter[K] { _ = "STUB: not implemented"; return nil }

func attemptMathOperation[K any](lhs Getter[K], op mathOp, rhs Getter[K]) Getter[K] {
	_ = "STUB: not implemented"
	return nil
}

func performOpTime(x time.Time, y any, op mathOp) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func performOpDuration(x time.Duration, y any, op mathOp) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func performOp[N int64 | float64](x, y N, op mathOp) (N, error) {
	_ = "STUB: not implemented"
	return *new(N), nil
}
