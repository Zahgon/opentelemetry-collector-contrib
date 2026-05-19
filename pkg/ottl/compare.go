// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"golang.org/x/exp/constraints"
)

// ValueComparator defines methods for comparing values using the OTTL comparison rules
// (https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/pkg/ottl/LANGUAGE.md#comparison-rules)
type ValueComparator interface {
	// Equal compares two values for equality, returning true if they are equals
	// according to the OTTL comparison rules.
	Equal(a, b any) bool
	// NotEqual compares two values for equality, returning true if they are different
	// according to the OTTL comparison rules.
	NotEqual(a, b any) bool
	// Less compares two values, returning true if the first value is less than the second
	// value, using the OTTL comparison rules.
	Less(a, b any) bool
	// LessEqual compares two values, returning true if the first value is less or equal
	// to the second value, using the OTTL comparison rules.
	LessEqual(a, b any) bool
	// Greater compares two values, returning true if the first value is greater than the
	// second value, using the OTTL comparison rules.
	Greater(a, b any) bool
	// GreaterEqual compares two values, returning true if the first value is greater or
	// equal to the second value, using the OTTL comparison rules.
	GreaterEqual(a, b any) bool
	// compare is a private method that compares two values using the grammar compareOp,
	// it also restricts custom implementations outside of this package.
	compare(a, b any, op compareOp) bool
}

// ottlValueComparator is the default implementation of the ValueComparator
type ottlValueComparator struct{}

// The functions in this file implement a general-purpose comparison of two
// values of type any, which for the purposes of OTTL mean values that are one of
// int, float, string, bool, or pointers to those, or []byte, or nil.

// invalidComparison returns false for everything except ne (where it returns true to indicate that the
// objects were definitely not equivalent).
func (*ottlValueComparator) invalidComparison(op compareOp) bool {
	_ = "STUB: not implemented"

	// comparePrimitives implements a generic comparison helper for all Ordered types (derived from Float, Int, or string).
	// According to benchmarks, it's faster than explicit comparison functions for these types.
	return false
}

func comparePrimitives[T constraints.Ordered](a, b T, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (*ottlValueComparator) compareBools(a, b bool, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (*ottlValueComparator) compareBytes(a, b []byte, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareBool(a bool, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareString(a string, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareByte(a []byte, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareInt64(a int64, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareFloat64(a float64, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareDuration(a time.Duration, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareTime(a time.Time, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareMap(a map[string]any, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) comparePMap(a pcommon.Map, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) compareSlice(a []any, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ottlValueComparator) comparePSlice(a pcommon.Slice, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	return false
}

// a and b are the return values from a Getter; we try to compare them
// according to the given operator.
func (p *ottlValueComparator) compare(a, b any, op compareOp) bool {
	_ = "STUB: not implemented"
	// nils are equal to each other and never equal to anything else,
	// so if they're both nil, report equality.
	return false
}

// Anything else, we switch on the left side first.

// If a was nil, it means b wasn't and inequalities don't apply,
// so let's swap and give it the chance to get evaluated.

// If we don't know what type it is, we can't do inequalities yet. So we can fall back to the old behavior where we just
// use Go's standard equality.

func (p *ottlValueComparator) Equal(a, b any) bool { _ = "STUB: not implemented"; return false }

func (p *ottlValueComparator) NotEqual(a, b any) bool { _ = "STUB: not implemented"; return false }

func (p *ottlValueComparator) Less(a, b any) bool { _ = "STUB: not implemented"; return false }

func (p *ottlValueComparator) LessEqual(a, b any) bool { _ = "STUB: not implemented"; return false }

func (p *ottlValueComparator) Greater(a, b any) bool { _ = "STUB: not implemented"; return false }

func (p *ottlValueComparator) GreaterEqual(a, b any) bool { _ = "STUB: not implemented"; return false }

// NewValueComparator creates a new ValueComparator instance that can be used to compare
// values using the OTTL comparison rules.
func NewValueComparator() ValueComparator { _ = "STUB: not implemented"; return *new(ValueComparator) }
