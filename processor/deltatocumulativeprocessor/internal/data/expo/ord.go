// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expo // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/expo"

import "cmp"

// HiLo returns the greater of a and b by comparing the result of applying fn to
// each. If equal, returns operands as passed
func HiLo[T any, N cmp.Ordered](a, b T, fn func(T) N) (hi, lo T) {
	_ = "STUB: not implemented"
	return *new(T), *new(T)
}
